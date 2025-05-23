package controllers

import (
	"context"
	"database/sql"
	"modulux/database"
	"modulux/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
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

// DeleteStudiengang deletes a studiengang by ID from the database
func DeleteStudiengang(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter must be a valid integer"})
		return
	}

	query := "DELETE FROM studiengang WHERE studiengang_id = $1"
	_, err = database.DB.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Studiengang deleted successfully"})
}

// GetModulverantwortlicheByStudiengang retrieves all modulverantwortliche for a specific studiengang_id
func GetModulverantwortlicheByStudiengang(c *gin.Context) {
	studiengangID := c.Param("id")
	var modulverantwortliche []models.Person

	query := `
        SELECT p.person_id, p.titel, p.vorname, p.nachname, p.email, p.telefonnummer, p.raum, p.funktion
        FROM person p
        JOIN modul_person_rolle mpr ON p.person_id = mpr.person_id
        JOIN modul_studiengang ms ON mpr.modul_kuerzel = ms.modul_kuerzel AND mpr.modul_version = ms.modul_version
        JOIN rolle r ON mpr.rolle_id = r.rolle_id
        WHERE ms.studiengang_id = $1 AND r.bezeichnung = 'Modulverantwortlicher'
    `
	rows, err := database.DB.Query(context.Background(), query, studiengangID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var person models.Person
		if err := rows.Scan(
			&person.PersonID,
			&person.Titel,
			&person.Vorname,
			&person.Nachname,
			&person.Email,
			&person.Telefonnummer,
			&person.Raum,
			&person.Funktion,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		modulverantwortliche = append(modulverantwortliche, person)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, modulverantwortliche)
}

// GetOpalLinks retrieves all modules with their opal links, kuerzel, and version from the database
func GetOpalLinks(c *gin.Context) {
	studiengang := c.Param("id")

	var modules []struct {
		Kuerzel  string      `json:"kuerzel"`
		Version  int         `json:"version"`
		OpalLink null.String `json:"opal_link"`
	}

	query := `
		SELECT m.kuerzel, m.version, m.opal_link
		FROM modul m
		JOIN modul_studiengang sm ON m.kuerzel = sm.modul_kuerzel AND m.version = sm.modul_version
		JOIN studiengang s ON sm.studiengang_id = s.studiengang_id
		WHERE s.studiengang_id = $1
	`
	rows, err := database.DB.Query(context.Background(), query, studiengang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var module struct {
			Kuerzel  string      `json:"kuerzel"`
			Version  int         `json:"version"`
			OpalLink null.String `json:"opal_link"`
		}
		err := rows.Scan(&module.Kuerzel, &module.Version, &module.OpalLink)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		modules = append(modules, module)
	}

	c.JSON(http.StatusOK, modules)
}

// GetModuleGoalsByStudiengangID retrieves all modules for a specific studiengang_id
func GetModuleGoalsByStudiengangID(c *gin.Context) {
	studiengangID := c.Param("id")
	var modules []struct {
		Kuerzel             string `json:"kuerzel"`
		Version             int    `json:"version"`
		Titel               string `json:"titel"`
		Zielqualifikationen string `json:"zielqualifikationen"`
	}

	query := `
        SELECT m.kuerzel, m.version, m.modultitel, m.qualifikationsziele
        FROM modul m
        JOIN modul_studiengang ms ON m.kuerzel = ms.modul_kuerzel AND m.version = ms.modul_version
        WHERE ms.studiengang_id = $1
    `
	rows, err := database.DB.Query(context.Background(), query, studiengangID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var module struct {
			Kuerzel             string `json:"kuerzel"`
			Version             int    `json:"version"`
			Titel               string `json:"titel"`
			Zielqualifikationen sql.NullString
		}
		err := rows.Scan(&module.Kuerzel, &module.Version, &module.Titel, &module.Zielqualifikationen)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		modules = append(modules, struct {
			Kuerzel             string `json:"kuerzel"`
			Version             int    `json:"version"`
			Titel               string `json:"titel"`
			Zielqualifikationen string `json:"zielqualifikationen"`
		}{
			Kuerzel:             module.Kuerzel,
			Version:             module.Version,
			Titel:               module.Titel,
			Zielqualifikationen: module.Zielqualifikationen.String,
		})
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "Keine Module gefunden"})
		return
	}

	c.JSON(http.StatusOK, modules)
}

// GetModuleWithLiteratureByStudiengangID retrieves all modules with their literature for a specific studiengang_id, excluding specified modules
func GetModuleWithLiteratureByStudiengangID(c *gin.Context) {
	studiengangIDStr := c.Param("id")
	studiengangID, err := strconv.Atoi(studiengangIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid studiengang_id"})
		return
	}

	type ExcludeModulesRequest struct {
		Exclude []string `json:"auszuschliessende_module"`
	}
	var request ExcludeModulesRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var modules []models.Module

	query := `
        SELECT m.kuerzel, m.version, m.modultitel, m.qualifikationsziele,
            l.literatur_id, l.titel, l.autor, l.jahr, l.verlag, l.isbn, l.link, l.doi
        FROM modul m
        JOIN modul_studiengang ms ON m.kuerzel = ms.modul_kuerzel AND m.version = ms.modul_version
        LEFT JOIN modul_literatur ml ON m.kuerzel = ml.modul_kuerzel AND m.version = ml.modul_version
        LEFT JOIN literatur l ON ml.literatur_id = l.literatur_id
        WHERE ms.studiengang_id = $1
    `

	// Wenn Module ausgeschlossen werden sollen, füge die entsprechenden Bedingungen zur Abfrage hinzu
	if len(request.Exclude) > 0 {
		excludeConditions := make([]string, len(request.Exclude))
		for i, exclude := range request.Exclude {
			parts := strings.Split(exclude, "-")
			if len(parts) == 2 {
				excludeConditions[i] = "(m.kuerzel != '" + parts[0] + "' OR m.version != " + parts[1] + ")"
			}
		}
		query += " AND " + strings.Join(excludeConditions, " AND ")
	}

	rows, err := database.DB.Query(context.Background(), query, studiengangID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	moduleMap := make(map[string]*models.Module)

	for rows.Next() {
		var module models.Module
		var literatur models.Literatur
		err := rows.Scan(&module.Kuerzel, &module.Version, &module.Modultitel, &module.Qualifikationsziele,
			&literatur.LiteraturID, &literatur.Titel, &literatur.Autor, &literatur.Jahr, &literatur.Verlag, &literatur.ISBN, &literatur.Link, &literatur.DOI)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		key := module.Kuerzel + "-" + strconv.Itoa(module.Version)
		if existingModule, exists := moduleMap[key]; exists {
			existingModule.Literatur = append(existingModule.Literatur, literatur)
		} else {
			if literatur.LiteraturID != (null.Int{}) {
				module.Literatur = []models.Literatur{literatur}
			}
			moduleMap[key] = &module
		}
	}

	for _, module := range moduleMap {
		modules = append(modules, *module)
	}
	c.JSON(http.StatusOK, modules)
}
