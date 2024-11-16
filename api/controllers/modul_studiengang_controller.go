package controllers

import (
	"context"
	"database/sql"
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
	modulVersionStr := c.Param("modul_version")
	studiengangID := c.Param("studiengang_id")
	var modulStudiengang models.ModulStudiengang

	modulVersion, err := strconv.Atoi(modulVersionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Modul version must be a valid integer"})
		return
	}

	query := `SELECT modul_kuerzel, modul_version, studiengang_id, semester, modul_typ FROM modul_studiengang 
              WHERE modul_kuerzel = $1 AND modul_version = $2 AND studiengang_id = $3`
	err = database.DB.QueryRow(context.Background(), query, modulKuerzel, modulVersion, studiengangID).Scan(
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
	var moduleStudiengaenge []models.ModulStudiengang

	modulVersion, err := strconv.Atoi(modulVersionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Modul version must be a valid integer"})
		return
	}

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
func GetModulStudiengangByStudiengangID(c *gin.Context) {

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

// AddModulStudiengang fügt einen Eintrag in die Tabelle modul_studiengang hinzu
func AddModulStudiengang(c *gin.Context) {
	var modulStudiengang models.ModulStudiengang

	if err := c.ShouldBindJSON(&modulStudiengang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Überprüfen, ob es vorausgesetzte Module gibt
	query := `SELECT vorausgesetztes_modul_kuerzel, vorausgesetztes_modul_version FROM modul_voraussetzung 
            WHERE modul_kuerzel = $1 AND modul_version = $2`
	rows, err := database.DB.Query(context.Background(), query, modulStudiengang.ModulKuerzel, modulStudiengang.ModulVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query modul_voraussetzung"})
		return
	}
	defer rows.Close()

	var vorausgesetzteModule []struct {
		Kuerzel string
		Version int
	}

	for rows.Next() {
		var kuerzel string
		var version int
		err := rows.Scan(&kuerzel, &version)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan modul_voraussetzung"})
			return
		}
		vorausgesetzteModule = append(vorausgesetzteModule, struct {
			Kuerzel string
			Version int
		}{kuerzel, version})
	}

	err = rows.Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over modul_voraussetzung"})
		return
	}

	// Wenn es vorausgesetzte Module gibt, überprüfen, ob jedes vorausgesetzte Modul in einem vorherigen Semester stattfindet
	for _, modul := range vorausgesetzteModule {
		var vorausgesetztesModulSemester int
		query = `SELECT semester FROM modul_studiengang 
                WHERE modul_kuerzel = $1 AND modul_version = $2 AND studiengang_id = $3`
		err = database.DB.QueryRow(context.Background(), query, modul.Kuerzel, modul.Version, modulStudiengang.StudiengangID).Scan(&vorausgesetztesModulSemester)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Prerequisite module not found in modul_studiengang"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query modul_studiengang for prerequisite module"})
			}
			return
		}

		if modulStudiengang.Semester.Valid {
			if vorausgesetztesModulSemester >= int(modulStudiengang.Semester.Int64) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Prerequisite module must be in a previous semester"})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Semester value is not valid"})
			return
		}
	}

	// Eintrag in die Tabelle modul_studiengang hinzufügen
	query = `INSERT INTO modul_studiengang (modul_kuerzel, modul_version, studiengang_id, semester, modul_typ) 
            VALUES ($1, $2, $3, $4, $5)`
	_, err = database.DB.Exec(context.Background(), query, modulStudiengang.ModulKuerzel, modulStudiengang.ModulVersion, modulStudiengang.StudiengangID, modulStudiengang.Semester, modulStudiengang.ModulTyp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert into modul_studiengang"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Eintrag erfolgreich hinzugefügt"})
}
