package main

import (
	"log"

	"github.com/singfield/rocket/internal/db"
	"github.com/singfield/rocket/internal/rocket"
)

func Run() error {
	// responsible for initializing and starting
	// our gRpc server

	rocketStore, err := db.New()
	if err != nil {
		return err
	}
	
	rocket.New(rocketStore)

	return nil
}

func main() {

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}