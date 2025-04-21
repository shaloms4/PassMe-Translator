package usecases

import (
	"errors"

	domain "github.com/shaloms4/Pass-Me-Core-Functionality/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase interface defines the business logic methods
type UserUseCase interface {
	RegisterUser(user *domain.User) error
	LoginUser(email, password string) (*domain.User, error)
}

// userUseCase implements the UserUseCase interface
type userUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase creates a new instance of user use case
func NewUserUseCase(repo domain.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

// RegisterUser creates a new user
func (uc *userUseCase) RegisterUser(user *domain.User) error {
	// Check if user with same email already exists
	existingUser, _ := uc.userRepo.FindUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Check if user with same username already exists
	existingUser, _ = uc.userRepo.FindUserByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already taken")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create the user
	return uc.userRepo.CreateUser(user)
}

// LoginUser authenticates a user
func (uc *userUseCase) LoginUser(email, password string) (*domain.User, error) {
	// Find user by email
	user, err := uc.userRepo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
} 