package producer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func FIOFailedInform(e error) {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
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
