package controller

import (
	"context"
	"log"
	"modulux/database"
	"modulux/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPersons retrieves all persons from the database
func GetPersons(c *gin.Context) {

	ctx := context.Background()
	var persons []models.Person

	// Debugging: Log the query
	log.Println("Executing query: SELECT person_id, titel, vorname, nachname, email, telefonnummer, raum, funktion FROM person")

	rows, err := database.DB.Query(ctx, "SELECT person_id, titel, vorname, nachname, email, telefonnummer, raum, funktion FROM person")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.PersonID, &person.Titel, &person.Vorname, &person.Nachname, &person.Email, &person.Telefonnummer, &person.Raum, &person.Funktion); err != nil {
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
