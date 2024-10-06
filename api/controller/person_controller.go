package controller

import (
	"context"
	"database/sql"
	"modulux/database"
	"modulux/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
