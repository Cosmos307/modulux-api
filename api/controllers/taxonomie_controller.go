package controllers

import (
	"context"
	"modulux/database"
	"net/http"
	"strings"

	"modulux/models"

	"github.com/gin-gonic/gin"
)

// GetVerbsByCategory retrieves verbes from the database by category
func GetVerbsByCategory(c *gin.Context) {

	category := c.Param("kategorie")

	verbs, err := retrieveVerbsFromDB(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"kategorie": category, "verben": verbs})
}

// GetTaxonomieFeedback returns feedback based on the number of verbs in the text
func GetTaxonomieFeedback(c *gin.Context) {

	var taxonomie models.Taxonomie
	err := c.ShouldBindJSON(&taxonomie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verbCount, err := CountVerbs(taxonomie.Text, taxonomie.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var status string
	if verbCount >= 5 {
		status = "grün"
	} else if verbCount >= 3 {
		status = "gelb"
	} else {
		status = "rot"
	}

	response := map[string]interface{}{
		"kategorie":    taxonomie.Category,
		"verbenanzahl": verbCount,
		"ampelzustand": status,
	}

	c.JSON(http.StatusOK, response)
}

// CountVerbs counts the number of verbs in a text by categorie
func CountVerbs(text, category string) (int, error) {

	verbs, err := retrieveVerbsFromDB(category)
	if err != nil {
		return 0, err
	}

	verbSet := make(map[string]struct{})
	for _, verb := range verbs {
		verbSet[verb] = struct{}{}
	}

	words := strings.Fields(text)
	verbCount := 0
	for _, word := range words {
		_, exists := verbSet[word]
		if exists {
			verbCount++
		}
	}

	return verbCount, nil
}

// getVerbsByCategory retrieves verbs for a specific category from the database
func retrieveVerbsFromDB(category string) ([]string, error) {

	rows, err := database.DB.Query(context.Background(), "SELECT verb FROM taxonomie_verb WHERE kategorie_id = (SELECT id FROM taxonomie_kategorie WHERE name = $1)", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var verbs []string
	for rows.Next() {
		var verb string
		err := rows.Scan(&verb)
		if err != nil {
			return nil, err
		}
		verbs = append(verbs, verb)
	}
	return verbs, nil
}
