package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	domain "github.com/shaloms4/Pass-Me-Core-Functionality/domain"
)

// userRepository is the implementation of the UserRepository interface
type userRepository struct {
	collection *mongo.Collection
}

// NewUserRepository initializes a new user repository
func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

// CreateUser stores a new user into the MongoDB database
func (r *userRepository) CreateUser(user *domain.User) error {
	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	// Update the user's ID with the MongoDB-generated ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	return nil
}

// FindUserByEmail retrieves a user by their email from MongoDB
func (r *userRepository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// FindUserByUsername retrieves a user by their username from MongoDB
func (r *userRepository) FindUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// repositories/user_repository.go

func (r *userRepository) FindUserByID(id string) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user domain.User
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUsername(id, newUsername string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"username": newUsername}},
	)
	return err
}

func (r *userRepository) UpdatePassword(id, hashedPassword string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"password": hashedPassword}},
	)
	return err
}
