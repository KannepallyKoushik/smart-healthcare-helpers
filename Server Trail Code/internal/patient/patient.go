package patient

import "context"

// Patient - should contain the defination of our patient
type Patient struct {
	ID              string
	Name            string
	Age             int
	Sex             bool
	Smoker          bool
	CigsPerDay      int
	PrevalentStroke bool
	Diabetes        bool
}

// Store - define the interface we expect
//  Our DB implementation to follow
type Store interface {
	GetPatientByID(id string) (Patient, error)
	InsertPatient(pnt Patient) (Patient, error)
}

// Service - our Patient service responsible for updating the Patient Medical Data
// updating the Patient Medical Data
type Service struct {
	Store Store
}

// GetPatientByID- get a Patient based on ID from the DB
func (s Service) GetPatientByID(ctx context.Context, id string) (Patient, error) {
	pnt, err := s.Store.GetPatientByID(id)
	if err != nil {
		return Patient{}, err
	}
}

//InsertPatient - create a new instance of our Patient
func (s Service) InsertPatient(ctx context.Context, pnt Patient) (Patient, error) {
	pnt, err := s.Store.InsertPatient(pnt)
	if err != nil {
		return Patient{}, err
	}
	return pnt, nil
}

// New - return a new instance of our Patient Service
func New(store Store) (Service, error) {
	return Service{
		Store: store,
	}, nil
}
