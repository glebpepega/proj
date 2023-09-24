package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/glebpepega/proj/internal/apicaller"
	"github.com/glebpepega/proj/internal/person"
	"github.com/segmentio/kafka-go"
)

func (s *server) readFIO() {
	r := s.kafkareader
	for {
		var person person.Person
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			if err := r.Close(); err != nil {
				s.logger.Fatal("failed to close reader:", err)
			}
			s.logger.Fatal(err)
		}
		s.logger.Infof("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		if err := json.Unmarshal(m.Value, &person); err != nil {
			s.writeFIOFailed(err)
			continue
		}
		if person.Name == "" || person.Surname == "" {
			s.writeFIOFailed(fmt.Errorf("name and surname reqired"))
			continue
		} else {
			_, err := apicaller.CallAPI(&person, "https://api.agify.io/?name=")
			if err != nil {
				s.writeFIOFailed(err)
				continue
			}
			_, err = apicaller.CallAPI(&person, "https://api.genderize.io/?name=")
			if err != nil {
				s.writeFIOFailed(err)
				continue
			}
			pUpdated, err := apicaller.CallAPI(&person, "https://api.nationalize.io/?name=")
			if err != nil {
				s.writeFIOFailed(err)
				continue
			} else {
				if err := s.postgres.StoreInDB(pUpdated); err != nil {
					s.logger.Fatal(err)
				}
			}
		}
	}
}

func (s *server) writeFIOFailed(e error) {
	w := s.kafkawriter
	if err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(`{"error":"` + e.Error() + `"}`),
		},
	); err != nil {
		s.logger.Fatal("failed to write messages:", err)
	} else {
		s.logger.Info("message to FIO_FAILED queue: ", err)
	}

	if err := w.Close(); err != nil {
		s.logger.Fatal("failed to close writer:", err)
	}
}
