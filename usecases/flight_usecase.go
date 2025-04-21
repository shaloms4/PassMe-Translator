package usecases

import "github.com/shaloms4/Pass-Me-Core-Functionality/domain"

// FlightUseCase interface defines the business logic methods
type FlightUseCase interface {
    AddFlight(flight *domain.Flight) error
    FetchFlightByID(id string) (*domain.Flight, error)
    DeleteFlight(id string) error
    FetchFlightsByUserID(userID string) ([]domain.Flight, error)
}

// flightUseCase implements the FlightUseCase interface
type flightUseCase struct {
    flightRepo domain.FlightRepository
}

// NewFlightUseCase creates a new instance of flight use case
func NewFlightUseCase(repo domain.FlightRepository) FlightUseCase {
    return &flightUseCase{
        flightRepo: repo,
    }
}

// AddFlight creates a new flight
func (uc *flightUseCase) AddFlight(flight *domain.Flight) error {
    return uc.flightRepo.CreateFlight(flight)
}

// FetchFlightByID retrieves a flight by its ID
func (uc *flightUseCase) FetchFlightByID(id string) (*domain.Flight, error) {
    return uc.flightRepo.GetFlightByID(id)
}

// DeleteFlight removes a flight by its ID
func (uc *flightUseCase) DeleteFlight(id string) error {
    return uc.flightRepo.DeleteFlight(id)
}

// FetchFlightsByUserID retrieves all flights for a specific user
func (uc *flightUseCase) FetchFlightsByUserID(userID string) ([]domain.Flight, error) {
    return uc.flightRepo.GetFlightsByUserID(userID)
}
