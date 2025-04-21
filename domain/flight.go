package domain

import (
	"time"
)

// Define a new type to hold a question and its corresponding answer
type QA struct {
	Question string `bson:"question" json:"question"`
	Answer   string `bson:"answer" json:"answer"`
}

type Flight struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Title       string    `bson:"title" json:"title"`
	FromCountry string    `bson:"from_country" json:"from_country"`
	ToCountry   string    `bson:"to_country" json:"to_country"`
	Date        time.Time `bson:"date" json:"date"`
	UserID      string    `bson:"user_id" json:"user_id"`
	QA          []QA      `bson:"qa" json:"qa"`
}

type FlightRepository interface {
	CreateFlight(flight *Flight) error
	GetFlightByID(id string) (*Flight, error)
	DeleteFlight(id string) error
	GetFlightsByUserID(userID string) ([]Flight, error)
}

type FlightUseCase interface {
	AddFlight(flight *Flight) error
	FetchFlightByID(id string) (*Flight, error)
	DeleteFlight(id string) error
	FetchFlightsByUserID(userID string) ([]Flight, error)
}
