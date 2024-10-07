package controller

import (
	"context"
	"modulux/database"
	"modulux/models"
	"net/http"

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
	version := c.Param("version")
	var module models.Module

	query := "SELECT * FROM modul WHERE kuerzel = $1 AND version = $2"
	err := database.DB.QueryRow(context.Background(), query, kuerzel, version).Scan(
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
	version := c.Param("version")
	var opalLink null.String

	query := "SELECT opal_link FROM modul WHERE kuerzel = $1 AND version = $2"
	err := database.DB.QueryRow(context.Background(), query, kuerzel, version).Scan(&opalLink)
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
