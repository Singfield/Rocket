package main

import (
	"log"

	"github.com/singfield/rocket/internal/db"
	"github.com/singfield/rocket/internal/rocket"
	"github.com/singfield/rocket/internal/transport/grpc"
)

func Run() error {
	// responsible for initializing and starting
	// our gRpc server
	log.Println("Starting up Rocket gRPC Service")
	
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migration")
		return err
	}

	rocketService :=rocket.New(rocketStore)

	rocketHandler := grpc.New(rocketService)

	if err :=rocketHandler.Serve(); err!=nil {
		return err
	}

	return nil
}

func main() {

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
