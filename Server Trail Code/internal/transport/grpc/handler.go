package grpc

import (
	"context"
	"log"
	"net"

	"github.com/KannepallyKoushik/Go-gRPC-Server/internal/patient"
	"google.golang.org/grpc"
)

type RockerService interface {
	GetPatientByID(ctx context.Context, id string) (patient.Patient, error)
	InsertPatient(ctx context.Context, pnt patient.Patient) (patient.Patient, error)
}

type Handler struct {
	PatientService PatientService
}

func New(ptnt PatientService) Handler {
	return Handler{PatientService: pnt}
}

func (h Handler) Server() error {
	lis, err := net.Listen("tcp", ":3003")
	if err != nil {
		log.Printf("Error could not listen at port 3003: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	pnt.RegisterPatientServiceServer(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Error could not serve: %v", err)
		return err
	}
	return nil
}

func (h Handler) GetPatientByID(ctx context.Context, req *pnt.GetPatientRequest) (*pnt.GetPatientResponse, error) {
	return &rkt.GetPatientResponse{}, nil
}

func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	return &rkt.AddRocketResponse{}, nil
}
