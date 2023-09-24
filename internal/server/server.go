package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/glebpepega/proj/internal/cache"
	"github.com/glebpepega/proj/internal/postgres"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type server struct {
	server      *http.Server
	logger      *logrus.Logger
	postgres    *postgres.Postgres
	cache       *cache.Cache
	kafkareader *kafka.Reader
	kafkawriter *kafka.Writer
}

func New() *server {
	return &server{
		server:      &http.Server{},
		logger:      logrus.New(),
		postgres:    postgres.New(),
		cache:       cache.New(),
		kafkareader: &kafka.Reader{},
		kafkawriter: &kafka.Writer{},
	}
}

func (s *server) configure() {
	s.server = &http.Server{
		Addr: ":8080",
	}
	s.logger.Level = logrus.DebugLevel
	if err := s.postgres.Configure(); err != nil {
		s.logger.Fatal(err)
	}
	s.cache.Configure()
	s.kafkareader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     "FIO",
		Partition: 0,
	})
	s.kafkawriter = &kafka.Writer{
		Addr:  kafka.TCP("kafka:9092"),
		Topic: "FIO_FAILED",
	}
}

func (s *server) Start() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	s.configure()
	s.logger.Info("starting the server")
	go s.readFIO()
	sm := http.NewServeMux()
	s.server.Handler = sm
	sm.HandleFunc("/person", s.handlePerson)
	s.logger.Fatal(s.server.ListenAndServe())
}

func (s *server) GracefulShutdown() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	if err := s.server.Shutdown(context.Background()); err != nil {
		s.logger.Infof("HTTP server Shutdown: %v", err)
	}
}
