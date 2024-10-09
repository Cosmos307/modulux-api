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

// UpdateStudiengang updates an existing studiengang in the database
func UpdateStudiengang(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter must be a valid integer"})
		return
	}

	var updatedStudiengang models.Studiengang
	err = c.ShouldBindJSON(&updatedStudiengang)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE studiengang SET kuerzel = $1, nummern_im_studienablaufplan = $2, studiengangstitel = $3, studiengangstitel_englisch = $4, 
	kommentar = $5, abschluss = $6, erste_immatrikulation = $7, erforderliche_credits = $8, 
	kapazitaet = $9, in_vollzeit_studierbar = $10, in_teilzeit_studierbar = $11, kooperativer_studiengang = $12, 
	doppelabschlussprogramm = $13, fernstudium = $14, englischsprachig = $15, teasertext = $16, 
	mobilitaetsfenster = $17, website = $18, imagefilm = $19, werbeflyer = $20, akkreditierungsurkunde = $21, fakultaet_id = $22 WHERE studiengang_id = $23`
	_, err = database.DB.Exec(context.Background(), query,
		updatedStudiengang.Kuerzel, updatedStudiengang.NummernImStudienablaufplan, updatedStudiengang.Studiengangstitel, updatedStudiengang.StudiengangstitelEnglisch,
		updatedStudiengang.Kommentar, updatedStudiengang.Abschluss, updatedStudiengang.ErsteImmatrikulation, updatedStudiengang.ErforderlicheCredits,
		updatedStudiengang.Kapazitaet, updatedStudiengang.InVollzeitStudierbar, updatedStudiengang.InTeilzeitStudierbar, updatedStudiengang.KooperativerStudiengang,
		updatedStudiengang.Doppelabschlussprogramm, updatedStudiengang.Fernstudium, updatedStudiengang.Englischsprachig, updatedStudiengang.Teasertext,
		updatedStudiengang.Mobilitaetsfenster, updatedStudiengang.Website, updatedStudiengang.Imagefilm, updatedStudiengang.Werbeflyer, updatedStudiengang.Akkreditierungsurkunde,
		updatedStudiengang.FakultaetID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStudiengang)
}

// CreateStudiengang creates a new studiengang in the database
func CreateStudiengang(c *gin.Context) {

	var newStudiengang models.Studiengang
	if err := c.ShouldBindJSON(&newStudiengang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO studiengang (kuerzel, nummern_im_studienablaufplan, studiengangstitel, studiengangstitel_englisch, kommentar, abschluss, erste_immatrikulation, erforderliche_credits, kapazitaet, in_vollzeit_studierbar, in_teilzeit_studierbar, kooperativer_studiengang, doppelabschlussprogramm, fernstudium, englischsprachig, teasertext, mobilitaetsfenster, website, imagefilm, werbeflyer, akkreditierungsurkunde, fakultaet_id) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) RETURNING studiengang_id`
	err := database.DB.QueryRow(context.Background(), query,
		newStudiengang.Kuerzel, newStudiengang.NummernImStudienablaufplan, newStudiengang.Studiengangstitel, newStudiengang.StudiengangstitelEnglisch,
		newStudiengang.Kommentar, newStudiengang.Abschluss, newStudiengang.ErsteImmatrikulation, newStudiengang.ErforderlicheCredits,
		newStudiengang.Kapazitaet, newStudiengang.InVollzeitStudierbar, newStudiengang.InTeilzeitStudierbar, newStudiengang.KooperativerStudiengang,
		newStudiengang.Doppelabschlussprogramm, newStudiengang.Fernstudium, newStudiengang.Englischsprachig, newStudiengang.Teasertext,
		newStudiengang.Mobilitaetsfenster, newStudiengang.Website, newStudiengang.Imagefilm, newStudiengang.Werbeflyer, newStudiengang.Akkreditierungsurkunde,
		newStudiengang.FakultaetID).Scan(&newStudiengang.StudiengangID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newStudiengang)
}
