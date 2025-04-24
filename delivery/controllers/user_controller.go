package controllers

import (
	"net/http"

	domain "github.com/shaloms4/Pass-Me-Core-Functionality/domain"
	Infrastructure "github.com/shaloms4/Pass-Me-Core-Functionality/infrastructure"
	usecases "github.com/shaloms4/Pass-Me-Core-Functionality/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(uc usecases.UserUseCase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

// Register handles user registration
func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	if err := uc.userUseCase.RegisterUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// Login handles user login
func (uc *UserController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if loginData.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	if loginData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	user, err := uc.userUseCase.LoginUser(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token with both email and user ID
	token, err := Infrastructure.GenerateJWT(user.Email, user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (uc *UserController) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	user, err := uc.userUseCase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch profile"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"about":    "This app helps users schedule flights and translate queries.", // Example About
	})
}

func (uc *UserController) ChangeUsername(c *gin.Context) {
	userID := c.GetString("user_id")
	var req struct {
		NewUsername string `json:"new_username"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.userUseCase.UpdateUsername(userID, req.NewUsername)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Username updated successfully"})
}

func (uc *UserController) ChangePassword(c *gin.Context) {
	userID := c.GetString("user_id")
	var req struct {
		OldPassword     string `json:"old_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.NewPassword != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "new password and confirm password do not match"})
		return
	}
	err := uc.userUseCase.UpdatePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
