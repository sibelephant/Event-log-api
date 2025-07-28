package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sibelephant/prisma/db"
	"github.com/sibelephant/internal/models"
)

// CreateUser creates a new user
func (app *application) CreateUser(c *gin.Context) {
	var user models.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := app.models.Users.CreateOne(
		db.Users.Email.Set(user.Email),
		db.Users.Name.Set(user.Name),
		db.Users.Password.Set(user.Password),
	).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to response model (without password)
	response := models.UserResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}

	c.JSON(http.StatusCreated, response)
}

// GetAllUsers retrieves all users
func (app *application) GetAllUsers(c *gin.Context) {
	users, err := app.models.Users.FindMany().Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to response models (without passwords)
	var response []models.UserResponse
	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"users": response})
}

// GetUser retrieves a single user by ID
func (app *application) GetUser(c *gin.Context) {
	id := c.Param("id")

	// Convert string ID to int
	var userID int
	if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := app.models.Users.FindUnique(
		db.Users.ID.Equals(userID),
	).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Convert to response model (without password)
	response := models.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	c.JSON(http.StatusOK, response)
}
