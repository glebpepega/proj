package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/glebpepega/proj/apicaller"
	"github.com/glebpepega/proj/person"
	"github.com/glebpepega/proj/producer"
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
)

func (s *server) Consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:     "FIO",
		Partition: 0,
	})
	r.SetOffset(8)

	validate := validator.New()

	for {
		var person person.Person
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
		if err := json.Unmarshal(m.Value, &person); err != nil {
			producer.FIOFailedInform(err)
			continue
		}
		if err := validate.Struct(person); err != nil {
			producer.FIOFailedInform(err)
			continue
		} else {
			_, err := apicaller.CallAPI(&person, "https://api.agify.io/?name=")
			if err != nil {
				producer.FIOFailedInform(err)
				continue
			}
			_, err = apicaller.CallAPI(&person, "https://api.genderize.io/?name=")
			if err != nil {
				producer.FIOFailedInform(err)
				continue
			}
			pUddated, err := apicaller.CallAPI(&person, "https://api.nationalize.io/?name=")
			if err != nil {
				producer.FIOFailedInform(err)
				continue
			} else {
				s.StoreInDB(pUddated)
			}
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
