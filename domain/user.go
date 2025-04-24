package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username" binding:"required"`
	Password string             `bson:"password,omitempty" json:"password" binding:"required"`
	Email    string             `bson:"email" json:"email" binding:"required,email"`
}

type UserRepository interface {
	CreateUser(user *User) error
	FindUserByEmail(email string) (*User, error)
	FindUserByUsername(username string) (*User, error)
	FindUserByID(id string) (*User, error)
	UpdateUsername(id, newUsername string) error
	UpdatePassword(id, hashedPassword string) error
}
