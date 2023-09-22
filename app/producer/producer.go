package producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func FIOFailedInform(e error) {
	w := &kafka.Writer{
		Addr:  kafka.TCP("kafka:9092"),
		Topic: "FIO_FAILED",
	}
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(`{"error":"` + e.Error() + `"}`),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
