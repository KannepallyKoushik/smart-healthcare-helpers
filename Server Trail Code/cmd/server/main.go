package main

import (
	"log"

	"github.com/KannepallyKoushik/Go-gRPC-Server/internal/patient"
	"github.com/KannepallyKoushik/Go-gRPC-Server/internal/transport/grpc"
)

func Run() error {
	// responsible for initializing and starting our gRPC server
	pntService := patient.New(patientStore)

	pntHandler := grpc.New(pntService)

	if err := pntHandler.Server(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
