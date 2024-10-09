package controller

import (
	"context"
	"modulux/database"
	"modulux/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetStudiengaenge retrieves all studiengaenge from the database
func GetStudiengaenge(c *gin.Context) {

	var studiengaenge []models.Studiengang
	query := "SELECT * FROM studiengang"

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var studiengang models.Studiengang
		err := rows.Scan(
			&studiengang.StudiengangID, &studiengang.Kuerzel, &studiengang.NummernImStudienablaufplan, &studiengang.Studiengangstitel,
			&studiengang.StudiengangstitelEnglisch, &studiengang.Kommentar, &studiengang.Abschluss, &studiengang.ErsteImmatrikulation,
			&studiengang.ErforderlicheCredits, &studiengang.Kapazitaet, &studiengang.InVollzeitStudierbar, &studiengang.InTeilzeitStudierbar,
			&studiengang.KooperativerStudiengang, &studiengang.Doppelabschlussprogramm, &studiengang.Fernstudium, &studiengang.Englischsprachig,
			&studiengang.Teasertext, &studiengang.Mobilitaetsfenster, &studiengang.Website, &studiengang.Imagefilm, &studiengang.Werbeflyer,
			&studiengang.Akkreditierungsurkunde, &studiengang.FakultaetID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		studiengaenge = append(studiengaenge, studiengang)
	}

	c.JSON(http.StatusOK, studiengaenge)
}

// GetStudiengang retrieves a studiengang by ID from the database
func GetStudiengang(c *gin.Context) {

	var studiengang models.Studiengang
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "StudiengangID parameter must be a valid integer"})
		return
	}

	query := "SELECT * FROM studiengang WHERE studiengang_id = $1"
	err = database.DB.QueryRow(context.Background(), query, id).Scan(
		&studiengang.StudiengangID, &studiengang.Kuerzel, &studiengang.NummernImStudienablaufplan, &studiengang.Studiengangstitel,
		&studiengang.StudiengangstitelEnglisch, &studiengang.Kommentar, &studiengang.Abschluss, &studiengang.ErsteImmatrikulation,
		&studiengang.ErforderlicheCredits, &studiengang.Kapazitaet, &studiengang.InVollzeitStudierbar, &studiengang.InTeilzeitStudierbar,
		&studiengang.KooperativerStudiengang, &studiengang.Doppelabschlussprogramm, &studiengang.Fernstudium, &studiengang.Englischsprachig,
		&studiengang.Teasertext, &studiengang.Mobilitaetsfenster, &studiengang.Website, &studiengang.Imagefilm, &studiengang.Werbeflyer,
		&studiengang.Akkreditierungsurkunde, &studiengang.FakultaetID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, studiengang)
}
