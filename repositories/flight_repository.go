package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	domain "github.com/shaloms4/Pass-Me-Core-Functionality/domain"
)

// flightRepository is the implementation of the FlightRepository interface
type flightRepository struct {
	collection *mongo.Collection
}

// NewFlightRepository initializes a new flight repository
func NewFlightRepository(db *mongo.Database) domain.FlightRepository {
	return &flightRepository{
		collection: db.Collection("flights"),
	}
}

// CreateFlight stores a new flight into the MongoDB database
func (r *flightRepository) CreateFlight(flight *domain.Flight) error {
	flight.Date = flight.Date.UTC() // Ensure consistency

	if flight.ID == "" {
		flight.ID = primitive.NewObjectID().Hex()
	}

	result, err := r.collection.InsertOne(context.Background(), flight)
	if err != nil {
		return err
	}

	if insertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		flight.ID = insertedID.Hex()
	}

	return nil
}

// GetFlightByID retrieves a flight by its ID from MongoDB
func (r *flightRepository) GetFlightByID(id string) (*domain.Flight, error) {
	var flight domain.Flight
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&flight)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("flight not found")
		}
		return nil, fmt.Errorf("error finding flight: %v", err)
	}
	return &flight, nil
}

// DeleteFlight removes a flight from MongoDB by its ID
func (r *flightRepository) DeleteFlight(id string) error {
	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("flight not found")
	}
	return nil
}

// GetFlightsByUserID retrieves all flights for a specific user from MongoDB
func (r *flightRepository) GetFlightsByUserID(userID string) ([]domain.Flight, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var flights []domain.Flight
	for cursor.Next(context.Background()) {
		var flight domain.Flight
		if err := cursor.Decode(&flight); err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}
	return flights, nil
}
