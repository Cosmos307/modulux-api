package controller

import (
	"context"
	"modulux/database"
	"modulux/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
)

// GetModules retrieves all modules from the database
func GetModules(c *gin.Context) {

	var modules []models.Module
	query := "SELECT * FROM modul"

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var module models.Module
		err := rows.Scan(
			&module.Kuerzel, &module.Version, &module.FruehererSchluessel, &module.Modultitel, &module.ModultitelEnglisch,
			&module.Kommentar, &module.Niveau, &module.Dauer, &module.Turnus, &module.StudiumIntegrale, &module.Sprachenzentrum,
			&module.OpalLink, &module.GruppengroesseVorlesung, &module.GruppengroesseUebung, &module.GruppengroessePraktikum,
			&module.Lehrform, &module.Medienform, &module.Lehrinhalte, &module.Qualifikationsziele, &module.SozialUndSelbstkompetenzen,
			&module.BesondereZulassungsvoraussetzungen, &module.EmpfohleneVoraussetzungen, &module.Fortsetzungsmoeglichkeiten,
			&module.Hinweise, &module.EctsCredits, &module.Workload, &module.PraesenzeitWocheVorlesung, &module.PraesenzeitWocheUebung,
			&module.PraesenzeitWochePraktikum, &module.PraesenzeitWocheSonstiges, &module.Selbststudienzeit, &module.SelbststudienzeitAufschluesselung,
			&module.AktuelleLehrressourcen, &module.Literatur, &module.ParentModulKuerzel, &module.ParentModulVersion, &module.FakultaetID,
			&module.StudienrichtungID, &module.VertiefungID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		modules = append(modules, module)
	}

	c.JSON(http.StatusOK, modules)
}

// GetModule retrieves a module by kuerzel and version from the database
func GetModule(c *gin.Context) {

	kuerzel := c.Param("kuerzel")
	versionStr := c.Param("version")
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version parameter must be a valid integer"})
		return
	}
	var module models.Module

	query := "SELECT * FROM modul WHERE kuerzel = $1 AND version = $2"
	err = database.DB.QueryRow(context.Background(), query, kuerzel, version).Scan(
		&module.Kuerzel, &module.Version, &module.FruehererSchluessel, &module.Modultitel, &module.ModultitelEnglisch,
		&module.Kommentar, &module.Niveau, &module.Dauer, &module.Turnus, &module.StudiumIntegrale, &module.Sprachenzentrum,
		&module.OpalLink, &module.GruppengroesseVorlesung, &module.GruppengroesseUebung, &module.GruppengroessePraktikum,
		&module.Lehrform, &module.Medienform, &module.Lehrinhalte, &module.Qualifikationsziele, &module.SozialUndSelbstkompetenzen,
		&module.BesondereZulassungsvoraussetzungen, &module.EmpfohleneVoraussetzungen, &module.Fortsetzungsmoeglichkeiten,
		&module.Hinweise, &module.EctsCredits, &module.Workload, &module.PraesenzeitWocheVorlesung, &module.PraesenzeitWocheUebung,
		&module.PraesenzeitWochePraktikum, &module.PraesenzeitWocheSonstiges, &module.Selbststudienzeit, &module.SelbststudienzeitAufschluesselung,
		&module.AktuelleLehrressourcen, &module.Literatur, &module.ParentModulKuerzel, &module.ParentModulVersion, &module.FakultaetID,
		&module.StudienrichtungID, &module.VertiefungID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, module)
}

// GetOpalLinks retrieves all modules with their opal links, kuerzel, and version from the database
func GetOpalLinks(c *gin.Context) {

	var modules []struct {
		Kuerzel  string      `json:"kuerzel"`
		Version  int         `json:"version"`
		OpalLink null.String `json:"opal_link"`
	}

	query := "SELECT kuerzel, version, opal_link FROM modul"
	rows, err := database.DB.Query(context.Background(), query)
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

// GetOpalLink retrieves the opal link of a module by kuerzel and version from the database
func GetOpalLink(c *gin.Context) {

	kuerzel := c.Param("kuerzel")
	versionStr := c.Param("version")
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version parameter must be a valid integer"})
		return
	}
	var opalLink null.String

	query := "SELECT opal_link FROM modul WHERE kuerzel = $1 AND version = $2"
	err = database.DB.QueryRow(context.Background(), query, kuerzel, version).Scan(&opalLink)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"opal_link": opalLink.ValueOrZero()})
}

// UpdateModule updates an existing module in the database
func UpdateModule(c *gin.Context) {

	kuerzel := c.Param("kuerzel")

	versionStr := c.Param("version")
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version parameter must be a valid integer"})
		return
	}

	var updatedModule models.Module
	if err := c.ShouldBindJSON(&updatedModule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE modul SET frueherer_schluessel = $1, modultitel = $2, modultitel_englisch = $3, kommentar = $4, 
	niveau = $5, dauer = $6, turnus = $7, studium_integrale = $8, sprachenzentrum = $9, 
	opal_link = $10, gruppengroesse_vorlesung = $11, gruppengroesse_uebung = $12, gruppengroesse_praktikum = $13, 
	lehrform = $14, medienform = $15, lehrinhalte = $16, qualifikationsziele = $17, sozial_und_selbstkompetenzen = $18, 
	besondere_zulassungsvoraussetzungen = $19, empfohlene_voraussetzungen = $20, fortsetzungsmoeglichkeiten = $21, 
	hinweise = $22, ects_credits = $23, workload = $24, praesenzeit_woche_vorlesung = $25, praesenzeit_woche_uebung = $26, 
	praesenzeit_woche_praktikum = $27, praesenzeit_woche_sonstiges = $28, selbststudienzeit = $29, selbststudienzeit_aufschluesselung = $30, 
	aktuelle_lehrressourcen = $31, literatur = $32, parent_modul_kuerzel = $33, parent_modul_version = $34, fakultaet_id = $35, 
	studienrichtung_id = $36, vertiefung_id = $37 WHERE kuerzel = $38 AND version = $39`
	_, err = database.DB.Exec(context.Background(), query,
		updatedModule.FruehererSchluessel, updatedModule.Modultitel, updatedModule.ModultitelEnglisch, updatedModule.Kommentar,
		updatedModule.Niveau, updatedModule.Dauer, updatedModule.Turnus, updatedModule.StudiumIntegrale, updatedModule.Sprachenzentrum,
		updatedModule.OpalLink, updatedModule.GruppengroesseVorlesung, updatedModule.GruppengroesseUebung, updatedModule.GruppengroessePraktikum,
		updatedModule.Lehrform, updatedModule.Medienform, updatedModule.Lehrinhalte, updatedModule.Qualifikationsziele, updatedModule.SozialUndSelbstkompetenzen,
		updatedModule.BesondereZulassungsvoraussetzungen, updatedModule.EmpfohleneVoraussetzungen, updatedModule.Fortsetzungsmoeglichkeiten,
		updatedModule.Hinweise, updatedModule.EctsCredits, updatedModule.Workload, updatedModule.PraesenzeitWocheVorlesung, updatedModule.PraesenzeitWocheUebung,
		updatedModule.PraesenzeitWochePraktikum, updatedModule.PraesenzeitWocheSonstiges, updatedModule.Selbststudienzeit, updatedModule.SelbststudienzeitAufschluesselung,
		updatedModule.AktuelleLehrressourcen, updatedModule.Literatur, updatedModule.ParentModulKuerzel, updatedModule.ParentModulVersion, updatedModule.FakultaetID,
		updatedModule.StudienrichtungID, updatedModule.VertiefungID, kuerzel, version,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedModule)
}

// CreateModule creates a new module in the database
func CreateModule(c *gin.Context) {

	var newModule models.Module
	err := c.ShouldBindJSON(&newModule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO modul (kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_integrale, sprachenzentrum, opal_link, gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen, besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, workload, praesenzeit_woche_vorlesung, praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, literatur, parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38)`
	_, err = database.DB.Exec(context.Background(), query,
		newModule.Kuerzel, newModule.Version, newModule.FruehererSchluessel, newModule.Modultitel, newModule.ModultitelEnglisch,
		newModule.Kommentar, newModule.Niveau, newModule.Dauer, newModule.Turnus, newModule.StudiumIntegrale, newModule.Sprachenzentrum,
		newModule.OpalLink, newModule.GruppengroesseVorlesung, newModule.GruppengroesseUebung, newModule.GruppengroessePraktikum,
		newModule.Lehrform, newModule.Medienform, newModule.Lehrinhalte, newModule.Qualifikationsziele, newModule.SozialUndSelbstkompetenzen,
		newModule.BesondereZulassungsvoraussetzungen, newModule.EmpfohleneVoraussetzungen, newModule.Fortsetzungsmoeglichkeiten,
		newModule.Hinweise, newModule.EctsCredits, newModule.Workload, newModule.PraesenzeitWocheVorlesung, newModule.PraesenzeitWocheUebung,
		newModule.PraesenzeitWochePraktikum, newModule.PraesenzeitWocheSonstiges, newModule.Selbststudienzeit, newModule.SelbststudienzeitAufschluesselung,
		newModule.AktuelleLehrressourcen, newModule.Literatur, newModule.ParentModulKuerzel, newModule.ParentModulVersion, newModule.FakultaetID,
		newModule.StudienrichtungID, newModule.VertiefungID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newModule)
}

// DeleteModule deletes a module by kuerzel and version from the database
func DeleteModule(c *gin.Context) {

	kuerzel := c.Param("kuerzel")
	versionStr := c.Param("version")
	version, err := strconv.Atoi(versionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version parameter must be a valid integer"})
		return
	}

	query := "DELETE FROM modul WHERE kuerzel = $1 AND version = $2"
	_, err = database.DB.Exec(context.Background(), query, kuerzel, version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Module deleted successfully"})
}
