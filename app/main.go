package main

import (
	"github.com/glebpepega/proj/server"
)

type Test struct {
	Name  string
	Count int
}

func main() {
	server.New().Start()
}
