package controllers

import (
	"context"
	"modulux/database"
	"modulux/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetModulVoraussetzungen retrieves all modul voraussetzungen from the database
func GetAllModulVoraussetzungen(c *gin.Context) {

	var voraussetzungen []models.ModulVoraussetzung
	query := "SELECT * FROM modul_voraussetzung"

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var voraussetzung models.ModulVoraussetzung
		err := rows.Scan(
			&voraussetzung.ModulKuerzel, &voraussetzung.ModulVersion, &voraussetzung.VorausgesetztesModulKuerzel, &voraussetzung.VorausgesetztesModulVersion,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		voraussetzungen = append(voraussetzungen, voraussetzung)
	}

	c.JSON(http.StatusOK, voraussetzungen)
}

// GetModulVoraussetzungen retrieves all vorausgesetzte module for a given modul kuerzel and version from the database
func GetModulVoraussetzungen(c *gin.Context) {

	modulKuerzel := c.Param("modul_kuerzel")
	modulVersionStr := c.Param("modul_version")

	modulVersion, err := strconv.Atoi(modulVersionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Modul version parameter must be a valid integer"})
		return
	}

	var voraussetzungen []models.ModulVoraussetzung

	query := "SELECT vorausgesetztes_modul_kuerzel, vorausgesetztes_modul_version FROM modul_voraussetzung WHERE modul_kuerzel = $1 AND modul_version = $2"
	rows, err := database.DB.Query(context.Background(), query, modulKuerzel, modulVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving vorausgesetzte module"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var voraussetzung models.ModulVoraussetzung
		err := rows.Scan(&voraussetzung.VorausgesetztesModulKuerzel, &voraussetzung.VorausgesetztesModulVersion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning vorausgesetzte module"})
			return
		}
		voraussetzungen = append(voraussetzungen, voraussetzung)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating through vorausgesetzte module"})
		return
	}

	c.JSON(http.StatusOK, voraussetzungen)
}

// CreateModulVoraussetzung creates a new modul voraussetzung in the database
func CreateModulVoraussetzung(c *gin.Context) {
	var newVoraussetzung models.ModulVoraussetzung
	if err := c.ShouldBindJSON(&newVoraussetzung); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Retrieve semester information for module and prerequisite module in the same study program
	query := `SELECT ms1.studiengang_id, ms1.semester, ms2.semester 
              FROM modul_studiengang ms1
              JOIN modul_studiengang ms2 
              ON ms1.studiengang_id = ms2.studiengang_id
              WHERE ms1.modul_kuerzel = $1 AND ms1.modul_version = $2
              AND ms2.modul_kuerzel = $3 AND ms2.modul_version = $4`
	rows, err := database.DB.Query(context.Background(), query,
		newVoraussetzung.ModulKuerzel, newVoraussetzung.ModulVersion,
		newVoraussetzung.VorausgesetztesModulKuerzel, newVoraussetzung.VorausgesetztesModulVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving semester information for the modules"})
		return
	}
	defer rows.Close()

	var problemModules []gin.H
	var modulSemester, vorausgesetztesModulSemester int
	for rows.Next() {
		var studiengangID int
		err := rows.Scan(&studiengangID, &modulSemester, &vorausgesetztesModulSemester)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning semester information"})
			return
		}

		// Check if the prerequisite module is in a later semester
		if vorausgesetztesModulSemester >= modulSemester {
			problemModules = append(problemModules, gin.H{
				"studiengang_id":                 studiengangID,
				"modul_kuerzel":                  newVoraussetzung.ModulKuerzel,
				"modul_version":                  newVoraussetzung.ModulVersion,
				"modul_semester":                 modulSemester,
				"vorausgesetztes_modul_kuerzel":  newVoraussetzung.VorausgesetztesModulKuerzel,
				"vorausgesetztes_modul_version":  newVoraussetzung.VorausgesetztesModulVersion,
				"vorausgesetztes_modul_semester": vorausgesetztesModulSemester,
			})
		}
	}

	if len(problemModules) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "The prerequisite module cannot be in the same or a later semester than the module in any study program",
			"problem_modules": problemModules,
		})
		return
	}

	// Insert the new modul voraussetzung
	query = `INSERT INTO modul_voraussetzung (modul_kuerzel, modul_version, vorausgesetztes_modul_kuerzel, vorausgesetztes_modul_version) 
    VALUES ($1, $2, $3, $4)`
	_, err = database.DB.Exec(context.Background(), query,
		newVoraussetzung.ModulKuerzel, newVoraussetzung.ModulVersion, newVoraussetzung.VorausgesetztesModulKuerzel, newVoraussetzung.VorausgesetztesModulVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newVoraussetzung)
}

// DeleteModulVoraussetzung deletes a modul voraussetzung by modul kuerzel and version from the database
func DeleteModulVoraussetzung(c *gin.Context) {

	modulKuerzel := c.Param("modul_kuerzel")
	modulVersionStr := c.Param("modul_version")
	vorausgesetztesModulKuerzel := c.Param("vorausgesetztes_modul_kuerzel")
	vorausgesetztesModulVersionStr := c.Param("vorausgesetztes_modul_version")

	modulVersion, err := strconv.Atoi(modulVersionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Modul version parameter must be a valid integer"})
		return
	}
	vorausgesetztesModulVersion, err := strconv.Atoi(vorausgesetztesModulVersionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vorausgesetztes modul version parameter must be a valid integer"})
		return
	}

	query := "DELETE FROM modul_voraussetzung WHERE modul_kuerzel = $1 AND modul_version = $2 AND vorausgesetztes_modul_kuerzel = $3 AND vorausgesetztes_modul_version = $4"
	result, err := database.DB.Exec(context.Background(), query, modulKuerzel, modulVersion, vorausgesetztesModulKuerzel, vorausgesetztesModulVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No modul voraussetzung found with the given parameters"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Modul voraussetzung deleted successfully"})
}
