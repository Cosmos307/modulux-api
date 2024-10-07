package controller

import (
	"context"
	"modulux/database"
	"modulux/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
