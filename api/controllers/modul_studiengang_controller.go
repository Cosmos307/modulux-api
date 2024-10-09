package controllers

import (
	"context"
	"modulux/database"
	"modulux/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllModulStudiengang retrieves all modul_studiengang entries
func GetAllModulStudiengang(c *gin.Context) {

	var moduleStudiengaenge []models.ModulStudiengang
	query := `SELECT modul_kuerzel, modul_version, studiengang_id, semester, modul_typ FROM modul_studiengang`

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var modulStudiengang models.ModulStudiengang
		err := rows.Scan(&modulStudiengang.ModulKuerzel, &modulStudiengang.ModulVersion, &modulStudiengang.StudiengangID, &modulStudiengang.Semester, &modulStudiengang.ModulTyp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		moduleStudiengaenge = append(moduleStudiengaenge, modulStudiengang)
	}

	c.JSON(http.StatusOK, moduleStudiengaenge)
}

// GetModulStudiengang retrieves a specific modul_studiengang entry
func GetSpecificModulStudiengang(c *gin.Context) {

	modulKuerzel := c.Param("modul_kuerzel")
	modulVersion := c.Param("modul_version")
	studiengangID := c.Param("studiengang_id")
	var modulStudiengang models.ModulStudiengang

	query := `SELECT modul_kuerzel, modul_version, studiengang_id, semester, modul_typ FROM modul_studiengang 
              WHERE modul_kuerzel = $1 AND modul_version = $2 AND studiengang_id = $3`
	err := database.DB.QueryRow(context.Background(), query, modulKuerzel, modulVersion, studiengangID).Scan(
		&modulStudiengang.ModulKuerzel, &modulStudiengang.ModulVersion, &modulStudiengang.StudiengangID,
		&modulStudiengang.Semester, &modulStudiengang.ModulTyp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	c.JSON(http.StatusOK, modulStudiengang)
}

// GetModulStudiengangByModul retrieves all modul_studiengang entries for a specific modul_kuerzel and modul_version
func GetModulStudiengangByModul(c *gin.Context) {

	modulKuerzel := c.Param("modul_kuerzel")
	modulVersionStr := c.Param("modul_version")
	modulVersion, err := strconv.Atoi(modulVersionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Modul version must be a valid integer"})
		return
	}
	var moduleStudiengaenge []models.ModulStudiengang

	query := `SELECT modul_kuerzel, modul_version, studiengang_id, semester, modul_typ FROM modul_studiengang 
	WHERE modul_kuerzel = $1 AND modul_version = $2`
	rows, err := database.DB.Query(context.Background(), query, modulKuerzel, modulVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var modulStudiengang models.ModulStudiengang
		if err := rows.Scan(&modulStudiengang.ModulKuerzel, &modulStudiengang.ModulVersion, &modulStudiengang.StudiengangID,
			&modulStudiengang.Semester, &modulStudiengang.ModulTyp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		moduleStudiengaenge = append(moduleStudiengaenge, modulStudiengang)
	}

	c.JSON(http.StatusOK, moduleStudiengaenge)
}

// GetModulStudiengangByStudiengang retrieves all modul_studiengang entries for a specific studiengang_id
func GetModulStudiengangByStudiengang(c *gin.Context) {

	studiengangID := c.Param("studiengang_id")
	var moduleStudiengaenge []models.ModulStudiengang

	query := `SELECT modul_kuerzel, modul_version, studiengang_id, semester, modul_typ FROM modul_studiengang 
              WHERE studiengang_id = $1`
	rows, err := database.DB.Query(context.Background(), query, studiengangID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var modulStudiengang models.ModulStudiengang
		if err := rows.Scan(&modulStudiengang.ModulKuerzel, &modulStudiengang.ModulVersion, &modulStudiengang.StudiengangID,
			&modulStudiengang.Semester, &modulStudiengang.ModulTyp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		moduleStudiengaenge = append(moduleStudiengaenge, modulStudiengang)
	}

	c.JSON(http.StatusOK, moduleStudiengaenge)
}
