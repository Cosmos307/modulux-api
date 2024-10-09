package controllers

import (
	"context"
	"modulux/database"
	"modulux/models"
	"net/http"

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
