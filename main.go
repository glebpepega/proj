package main

import "github.com/glebpepega/proj/internal/server"

func main() {
	s := server.New()
	go s.Start()
	s.GracefulShutdown()
}
