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
			&voraussetzung.StudiengangID, &voraussetzung.ModulKuerzel, &voraussetzung.ModulVersion, &voraussetzung.VorausgesetztesModulKuerzel, &voraussetzung.VorausgesetztesModulVersion,
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
	studiengangIDStr := c.Param("studiengang_id")
	modulKuerzel := c.Param("modul_kuerzel")
	modulVersionStr := c.Param("modul_version")

	studiengangID, err := strconv.Atoi(studiengangIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Studiengang ID parameter must be a valid integer"})
		return
	}
	modulVersion, err := strconv.Atoi(modulVersionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Modul version parameter must be a valid integer"})
		return
	}

	var voraussetzungen []models.ModulVoraussetzung

	query := "SELECT vorausgesetztes_modul_kuerzel, vorausgesetztes_modul_version FROM modul_voraussetzung WHERE studiengang_id = $1 AND modul_kuerzel = $2 AND modul_version = $3"
	rows, err := database.DB.Query(context.Background(), query, studiengangID, modulKuerzel, modulVersion)
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

		voraussetzung.StudiengangID = studiengangID
		voraussetzung.ModulKuerzel = modulKuerzel
		voraussetzung.ModulVersion = modulVersion
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
	err := c.ShouldBindJSON(&newVoraussetzung)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if the prerequisite module exists in the same study program
	existsQuery := `
        SELECT COUNT(*) 
        FROM modul_studiengang 
        WHERE studiengang_id = $1 AND modul_kuerzel = $2 AND modul_version = $3
    `
	var count int
	err = database.DB.QueryRow(context.Background(), existsQuery, newVoraussetzung.StudiengangID, newVoraussetzung.VorausgesetztesModulKuerzel, newVoraussetzung.VorausgesetztesModulVersion).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existence of the prerequisite module"})
		return
	}
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Das vorrauszusetzende Modul existiert nicht in dem Studiengang"})
		return
	}

	// Retrieve semester information for module and prerequisite module in the same study program
	query := `
		SELECT ms1.studiengang_id, ms1.semester, ms2.semester 
		FROM modul_studiengang ms1
		JOIN modul_studiengang ms2 
		ON ms1.studiengang_id = ms2.studiengang_id
		WHERE ms1.modul_kuerzel = $1 AND ms1.modul_version = $2
		AND ms2.modul_kuerzel = $3 AND ms2.modul_version = $4 AND ms1.studiengang_id = $5
	`
	rows, err := database.DB.Query(context.Background(), query,
		newVoraussetzung.ModulKuerzel, newVoraussetzung.ModulVersion,
		newVoraussetzung.VorausgesetztesModulKuerzel, newVoraussetzung.VorausgesetztesModulVersion, newVoraussetzung.StudiengangID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler bei der Abfrage der Semesternummern"})
		return
	}
	defer rows.Close()

	var problemModules []gin.H
	var modulSemester, vorausgesetztesModulSemester, studiengangID int

	for rows.Next() {
		err := rows.Scan(&studiengangID, &modulSemester, &vorausgesetztesModulSemester)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning semester information"})
			return
		}

		// Check if the prerequisite module is in a later semester
		if vorausgesetztesModulSemester >= modulSemester {
			problemModules = append(problemModules, gin.H{
				"studiengang_id": studiengangID,
				"modul_kuerzel":  newVoraussetzung.ModulKuerzel,
				"modul_version":  newVoraussetzung.ModulVersion,
				"modul_semester": modulSemester,
			})
			problemModules = append(problemModules, gin.H{
				"studiengang_id": studiengangID,
				"modul_kuerzel":  newVoraussetzung.VorausgesetztesModulKuerzel,
				"modul_version":  newVoraussetzung.VorausgesetztesModulVersion,
				"modul_semester": vorausgesetztesModulSemester,
			})
		}
	}

	if len(problemModules) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "Das vorrauszusetzende Modul darf nicht im gleichen oder einem späteren Semester im Studiengang stattfinden",
			"problem_modules": problemModules,
		})
		return
	}

	// Insert the new modul voraussetzung
	query = `INSERT INTO modul_voraussetzung (studiengang_id, modul_kuerzel, modul_version, vorausgesetztes_modul_kuerzel, vorausgesetztes_modul_version) 
    VALUES ($1, $2, $3, $4, $5)`
	_, err = database.DB.Exec(context.Background(), query,
		studiengangID, newVoraussetzung.ModulKuerzel, newVoraussetzung.ModulVersion, newVoraussetzung.VorausgesetztesModulKuerzel, newVoraussetzung.VorausgesetztesModulVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "erfolgreich eingefügt"})
}

// GetModulVorschlaege retrieves all modules that are suggested for a given module in a study program
func GetModulVorschlaege(c *gin.Context) {
	studiengangIDStr := c.Param("studiengang_id")
	modulKuerzel := c.Param("modul_kuerzel")
	modulVersionStr := c.Param("modul_version")

	studiengangID, err := strconv.Atoi(studiengangIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Studiengang ID"})
		return
	}

	modulVersion, err := strconv.Atoi(modulVersionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Modulversion"})
		return
	}

	// query to get the semester of the given module
	var modulSemester int
	query := `
        SELECT semester
        FROM modul_studiengang
        WHERE studiengang_id = $1 AND modul_kuerzel = $2 AND modul_version = $3
    `
	err = database.DB.QueryRow(context.Background(), query, studiengangID, modulKuerzel, modulVersion).Scan(&modulSemester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Abrufen des Semesters"})
		return
	}

	// query to get all modules that are in the same study program and have a lower semester number then the given module
	query = `
        SELECT m.kuerzel, m.version, m.modultitel, ms.semester
        FROM modul m
        JOIN modul_studiengang ms ON m.kuerzel = ms.modul_kuerzel AND m.version = ms.modul_version
        WHERE ms.studiengang_id = $1 AND ms.semester < $2
    `
	rows, err := database.DB.Query(context.Background(), query, studiengangID, modulSemester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Abrufen der Module"})
		return
	}
	defer rows.Close()

	var moduleVorschlaege []gin.H
	for rows.Next() {
		var kuerzel string
		var version int
		var modultitel string
		var semester int
		err := rows.Scan(&kuerzel, &version, &modultitel, &semester)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Scannen der Module"})
			return
		}
		moduleVorschlaege = append(moduleVorschlaege, gin.H{
			"modul_kuerzel": kuerzel,
			"modul_version": version,
			"modul_titel":   modultitel,
			"semester":      semester,
		})
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Iterieren durch die Module"})
		return
	}

	c.JSON(http.StatusOK, moduleVorschlaege)
}

// DeleteModulVoraussetzung deletes a modul voraussetzung by modul kuerzel and version from the database
func DeleteModulVoraussetzung(c *gin.Context) {
	studiengangIDStr := c.Param("studiengang_id")
	modulKuerzel := c.Param("modul_kuerzel")
	modulVersionStr := c.Param("modul_version")
	vorausgesetztesModulKuerzel := c.Param("vorausgesetztes_modul_kuerzel")
	vorausgesetztesModulVersionStr := c.Param("vorausgesetztes_modul_version")

	studiengangID, err := strconv.Atoi(studiengangIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Studiengang ID parameter must be a valid integer"})
		return
	}

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

	query := "DELETE FROM modul_voraussetzung WHERE studiengang_id = $1 AND modul_kuerzel = $2 AND modul_version = $3 AND vorausgesetztes_modul_kuerzel = $4 AND vorausgesetztes_modul_version = $5"
	result, err := database.DB.Exec(context.Background(), query, studiengangID, modulKuerzel, modulVersion, vorausgesetztesModulKuerzel, vorausgesetztesModulVersion)
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
