package controllers

import (
	"context"
	"net/http"

	"modulux/api/middleware"
	"modulux/database"
	"modulux/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ValidateJWT validates a JWT and returns the claims
func Login(c *gin.Context) {
	var req models.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check login credentials
	var person models.Person
	query := `SELECT person_id, email, password FROM person WHERE email = $1`
	err = database.DB.QueryRow(context.Background(), query, req.Email).Scan(&person.PersonID, &person.Email, &person.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !CheckPasswordHash(req.Password, person.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	token, err := middleware.GenerateJWT(person.PersonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// CheckPasswordHash compares a password with a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
