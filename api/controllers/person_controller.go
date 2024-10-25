package controllers

import (
	"context"
	"database/sql"
	"modulux/database"
	"modulux/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetPersons retrieves all persons from the database
func GetPersons(c *gin.Context) {

	ctx := context.Background()
	var persons []models.Person

	query := "SELECT person_id, titel, vorname, nachname, email, telefonnummer, raum, funktion FROM person"
	rows, err := database.DB.Query(ctx, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.PersonID, &person.Titel, &person.Vorname, &person.Nachname, &person.Email, &person.Telefonnummer, &person.Raum, &person.Funktion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		persons = append(persons, person)
	}

	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": rows.Err().Error()})
		return
	}

	c.JSON(http.StatusOK, persons)

}

// GetPerson retrieves a person by ID from the database
func GetPerson(c *gin.Context) {

	personID := c.Param("id")
	ctx := context.Background()
	var person models.Person

	if personID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	query := "SELECT person_id, titel, vorname, nachname, email, telefonnummer, raum, funktion FROM person WHERE person_id = $1"
	err := database.DB.QueryRow(ctx, query, personID).Scan(&person.PersonID, &person.Titel, &person.Vorname, &person.Nachname, &person.Email, &person.Telefonnummer, &person.Raum, &person.Funktion)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, person)
}

// UpdatePerson updates an existing person in the database
func UpdatePerson(c *gin.Context) {

	personID := c.Param("id")
	ctx := context.Background()
	var updatedPerson models.Person

	if personID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	err := c.ShouldBindJSON(&updatedPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password if it is provided
	if updatedPerson.Password != "" {
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(updatedPerson.Password), bcrypt.DefaultCost)
		if hashErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": hashErr.Error()})
			return
		}
		updatedPerson.Password = string(hashedPassword)
	}

	query := "UPDATE person SET titel = $1, vorname = $2, nachname = $3, email = $4, telefonnummer = $5, raum = $6, funktion = $7, password = $8 WHERE person_id = $9"
	_, execErr := database.DB.Exec(ctx, query, updatedPerson.Titel, updatedPerson.Vorname, updatedPerson.Nachname, updatedPerson.Email, updatedPerson.Telefonnummer, updatedPerson.Raum, updatedPerson.Funktion, updatedPerson.Password, personID)
	if execErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": execErr.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPerson)
}

// CreatePerson creates a new person in the database
func CreatePerson(c *gin.Context) {

	var newPerson models.Person

	err := c.ShouldBindJSON(&newPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(newPerson.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": hashErr.Error()})
		return
	}
	newPerson.Password = string(hashedPassword)

	ctx := context.Background()
	query := "INSERT INTO person (titel, vorname, nachname, email, telefonnummer, raum, funktion, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING person_id"
	dbErr := database.DB.QueryRow(ctx, query, newPerson.Titel, newPerson.Vorname, newPerson.Nachname, newPerson.Email, newPerson.Telefonnummer, newPerson.Raum, newPerson.Funktion, newPerson.Password).Scan(&newPerson.PersonID)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPerson)
}

// DeletePerson deletes a person from the database
func DeletePerson(c *gin.Context) {

	personID := c.Param("id")
	ctx := context.Background()

	if personID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person ID"})
		return
	}

	query := "DELETE FROM person WHERE person_id = $1"
	_, execErr := database.DB.Exec(ctx, query, personID)
	if execErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": execErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}
