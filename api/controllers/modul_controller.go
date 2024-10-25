package controllers

import (
	"context"
	"fmt"
	"modulux/database"
	"modulux/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/jackc/pgx/v5"
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
	version, err := strconv.Atoi(c.Param("version"))
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
	version, err := strconv.Atoi(c.Param("version"))
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

// CreateModule creates a new module in the database
func CreateModule(c *gin.Context) {
	var newModule models.Module
	err := c.ShouldBindJSON(&newModule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO modul (
	kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_integrale, 
	sprachenzentrum, opal_link, gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, 
	qualifikationsziele, sozial_und_selbstkompetenzen, besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, 
	hinweise, ects_credits, praesenzeit_woche_vorlesung, praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, 
	selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, literatur, parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id)  
    VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37
	)`
	_, err = database.DB.Exec(context.Background(), query,
		newModule.Kuerzel, newModule.Version, newModule.FruehererSchluessel, newModule.Modultitel, newModule.ModultitelEnglisch,
		newModule.Kommentar, newModule.Niveau, newModule.Dauer, newModule.Turnus, newModule.StudiumIntegrale, newModule.Sprachenzentrum,
		newModule.OpalLink, newModule.GruppengroesseVorlesung, newModule.GruppengroesseUebung, newModule.GruppengroessePraktikum,
		newModule.Lehrform, newModule.Medienform, newModule.Lehrinhalte, newModule.Qualifikationsziele, newModule.SozialUndSelbstkompetenzen,
		newModule.BesondereZulassungsvoraussetzungen, newModule.EmpfohleneVoraussetzungen, newModule.Fortsetzungsmoeglichkeiten,
		newModule.Hinweise, newModule.EctsCredits, newModule.PraesenzeitWocheVorlesung, newModule.PraesenzeitWocheUebung,
		newModule.PraesenzeitWochePraktikum, newModule.PraesenzeitWocheSonstiges, newModule.SelbststudienzeitAufschluesselung,
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
	version, err := strconv.Atoi(c.Param("version"))
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
	c.JSON(http.StatusOK, gin.H{"message": "Modul deleted successfully"})
}

// UpdateOrCreateModuleVersion updates or creates a new version of a module
func UpdateOrCreateModuleVersion(c *gin.Context) {
	kuerzel := c.Param("kuerzel")
	version, err := strconv.Atoi(c.Param("version"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version parameter must be a valid integer"})
		return
	}

	var updatedModule models.Module
	if err := c.ShouldBindJSON(&updatedModule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer tx.Rollback(context.Background())

	previousHistoryID, err := saveModuleHistory(tx, kuerzel, version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	currentModule, err := getCurrentModule(tx, kuerzel, version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if hasSignificantChanges(currentModule, updatedModule) {
		maxVersion, err := getMaxModuleVersion(kuerzel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newVersion := maxVersion + 1
		newModule := updatedModule
		newModule.Version = newVersion

		err = InsertModule(tx, kuerzel, newVersion, newModule, previousHistoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	} else if hasMinorChanges(currentModule, updatedModule) {
		err = updateModule(tx, kuerzel, version, updatedModule, previousHistoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedModule)
}

// getMaxModuleVersion returns the highest version number of a module
func getMaxModuleVersion(kuerzel string) (int, error) {
	var maxVersion int
	err := database.DB.QueryRow(context.Background(), "SELECT COALESCE(MAX(version), 0) FROM modul WHERE kuerzel = $1", kuerzel).Scan(&maxVersion)
	if err != nil {
		return 0, fmt.Errorf("failed to get max module version: %w", err)
	}
	return maxVersion, nil
}

// InsertModule inserts a new module into the database
func InsertModule(tx pgx.Tx, kuerzel string, version int, module models.Module, vorherigerZustandID int) error {
	query := `
        INSERT INTO modul (
            kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_integrale, sprachenzentrum, opal_link,
            gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
            besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
            praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, literatur,
            parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id, vorheriger_zustand_id
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38
        )
	`

	_, err := tx.Exec(context.Background(), query,
		kuerzel,
		version,
		module.FruehererSchluessel,
		module.Modultitel,
		module.ModultitelEnglisch,
		module.Kommentar,
		module.Niveau,
		module.Dauer,
		module.Turnus,
		module.StudiumIntegrale,
		module.Sprachenzentrum,
		module.OpalLink,
		module.GruppengroesseVorlesung,
		module.GruppengroesseUebung,
		module.GruppengroessePraktikum,
		module.Lehrform,
		module.Medienform,
		module.Lehrinhalte,
		module.Qualifikationsziele,
		module.SozialUndSelbstkompetenzen,
		module.BesondereZulassungsvoraussetzungen,
		module.EmpfohleneVoraussetzungen,
		module.Fortsetzungsmoeglichkeiten,
		module.Hinweise,
		module.EctsCredits,
		module.PraesenzeitWocheVorlesung,
		module.PraesenzeitWocheUebung,
		module.PraesenzeitWochePraktikum,
		module.PraesenzeitWocheSonstiges,
		module.SelbststudienzeitAufschluesselung,
		module.AktuelleLehrressourcen,
		module.Literatur,
		module.ParentModulKuerzel,
		module.ParentModulVersion,
		module.FakultaetID,
		module.StudienrichtungID,
		module.VertiefungID,
		vorherigerZustandID,
	)
	if err != nil {
		return fmt.Errorf("failed to insert module: %w", err)
	}
	return nil
}

// saveModuleHistory saves the current state of a module to the modul_historie table
func saveModuleHistory(tx pgx.Tx, kuerzel string, version int) (int, error) {

	var historyID int
	query := `
        INSERT INTO modul_historie (
            kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_integrale, sprachenzentrum, opal_link,
            gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
            besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
            praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, literatur,
            parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id, vorheriger_zustand_id
        )
        SELECT
            kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_integrale, sprachenzentrum, opal_link,
            gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
            besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
            praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, literatur,
            parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id, vorheriger_zustand_id
        FROM modul
        WHERE kuerzel = $1 AND version = $2
        RETURNING id
    `
	err := tx.QueryRow(context.Background(), query, kuerzel, version).Scan(&historyID)
	if err != nil {
		return 0, err
	}
	return historyID, nil
}

// getCurrentModule retrieves the current module from the database
func getCurrentModule(tx pgx.Tx, kuerzel string, version int) (models.Module, error) {
	var module models.Module
	query := `
		SELECT kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_integrale, sprachenzentrum, opal_link,
			gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
			besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
			praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, literatur,
			parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id, vorheriger_zustand_id
		FROM modul
		WHERE kuerzel = $1 AND version = $2
	`
	err := tx.QueryRow(context.Background(), query, kuerzel, version).Scan(
		&module.Kuerzel, &module.Version, &module.FruehererSchluessel, &module.Modultitel, &module.ModultitelEnglisch, &module.Kommentar,
		&module.Niveau, &module.Dauer, &module.Turnus, &module.StudiumIntegrale, &module.Sprachenzentrum, &module.OpalLink,
		&module.GruppengroesseVorlesung, &module.GruppengroesseUebung, &module.GruppengroessePraktikum, &module.Lehrform, &module.Medienform,
		&module.Lehrinhalte, &module.Qualifikationsziele, &module.SozialUndSelbstkompetenzen, &module.BesondereZulassungsvoraussetzungen,
		&module.EmpfohleneVoraussetzungen, &module.Fortsetzungsmoeglichkeiten, &module.Hinweise, &module.EctsCredits, &module.PraesenzeitWocheVorlesung,
		&module.PraesenzeitWocheUebung, &module.PraesenzeitWochePraktikum, &module.PraesenzeitWocheSonstiges, &module.Selbststudienzeit,
		&module.SelbststudienzeitAufschluesselung, &module.AktuelleLehrressourcen, &module.Literatur, &module.ParentModulKuerzel, &module.ParentModulVersion,
		&module.FakultaetID, &module.StudienrichtungID, &module.VertiefungID, &module.VorherigerZustandID,
	)
	if err != nil {
		return models.Module{}, fmt.Errorf("failed to get current module: %w", err)
	}
	return module, nil
}

// hasSignificantChanges checks if there are significant changes between the current and updated module
func hasSignificantChanges(currentModule, updatedModule models.Module) bool {
	return updatedModule.Qualifikationsziele != currentModule.Qualifikationsziele ||
		updatedModule.Dauer != currentModule.Dauer ||
		updatedModule.EctsCredits != currentModule.EctsCredits ||
		updatedModule.PraesenzeitWocheVorlesung != currentModule.PraesenzeitWocheVorlesung ||
		updatedModule.PraesenzeitWocheUebung != currentModule.PraesenzeitWocheUebung ||
		updatedModule.PraesenzeitWochePraktikum != currentModule.PraesenzeitWochePraktikum ||
		updatedModule.PraesenzeitWocheSonstiges != currentModule.PraesenzeitWocheSonstiges ||
		updatedModule.Lehrform != currentModule.Lehrform ||
		updatedModule.Medienform != currentModule.Medienform ||
		updatedModule.BesondereZulassungsvoraussetzungen != currentModule.BesondereZulassungsvoraussetzungen ||
		updatedModule.EmpfohleneVoraussetzungen != currentModule.EmpfohleneVoraussetzungen ||
		updatedModule.Fortsetzungsmoeglichkeiten != currentModule.Fortsetzungsmoeglichkeiten
}

// hasMinorChanges checks if there are minor changes between the current and updated module
func hasMinorChanges(currentModule, updatedModule models.Module) bool {
	return updatedModule.FruehererSchluessel != currentModule.FruehererSchluessel ||
		updatedModule.Modultitel != currentModule.Modultitel ||
		updatedModule.ModultitelEnglisch != currentModule.ModultitelEnglisch ||
		updatedModule.Kommentar != currentModule.Kommentar ||
		updatedModule.Niveau != currentModule.Niveau ||
		updatedModule.Turnus != currentModule.Turnus ||
		updatedModule.StudiumIntegrale != currentModule.StudiumIntegrale ||
		updatedModule.Sprachenzentrum != currentModule.Sprachenzentrum ||
		updatedModule.OpalLink != currentModule.OpalLink ||
		updatedModule.GruppengroesseVorlesung != currentModule.GruppengroesseVorlesung ||
		updatedModule.GruppengroesseUebung != currentModule.GruppengroesseUebung ||
		updatedModule.GruppengroessePraktikum != currentModule.GruppengroessePraktikum ||
		updatedModule.Lehrinhalte != currentModule.Lehrinhalte ||
		updatedModule.SozialUndSelbstkompetenzen != currentModule.SozialUndSelbstkompetenzen ||
		updatedModule.Hinweise != currentModule.Hinweise ||
		updatedModule.SelbststudienzeitAufschluesselung != currentModule.SelbststudienzeitAufschluesselung ||
		updatedModule.AktuelleLehrressourcen != currentModule.AktuelleLehrressourcen ||
		updatedModule.Literatur != currentModule.Literatur ||
		updatedModule.ParentModulKuerzel != currentModule.ParentModulKuerzel ||
		updatedModule.ParentModulVersion != currentModule.ParentModulVersion ||
		updatedModule.FakultaetID != currentModule.FakultaetID ||
		updatedModule.StudienrichtungID != currentModule.StudienrichtungID ||
		updatedModule.VertiefungID != currentModule.VertiefungID
}

// updateModule updates a module in the database
func updateModule(tx pgx.Tx, kuerzel string, version int, updatedModule models.Module, vorherigerZustandID int) error {
	query := `
        UPDATE modul
        SET 
            frueherer_schluessel = $1,
            modultitel = $2,
            modultitel_englisch = $3,
            kommentar = $4,
            niveau = $5,
            turnus = $6,
            studium_integrale = $7,
            sprachenzentrum = $8,
            opal_link = $9,
            gruppengroesse_vorlesung = $10,
            gruppengroesse_uebung = $11,
            gruppengroesse_praktikum = $12,
            lehrinhalte = $13,
            sozial_und_selbstkompetenzen = $14,
            hinweise = $15,
            selbststudienzeit_aufschluesselung = $16,
            aktuelle_lehrressourcen = $17,
            literatur = $18,
            parent_modul_kuerzel = $19,
            parent_modul_version = $20,
            fakultaet_id = $21,
            studienrichtung_id = $22,
            vertiefung_id = $23,
            vorheriger_zustand_id = $24
        WHERE kuerzel = $25 AND version = $26
    `

	_, err := tx.Exec(context.Background(), query,
		updatedModule.FruehererSchluessel,
		updatedModule.Modultitel,
		updatedModule.ModultitelEnglisch,
		updatedModule.Kommentar,
		updatedModule.Niveau,
		updatedModule.Turnus,
		updatedModule.StudiumIntegrale,
		updatedModule.Sprachenzentrum,
		updatedModule.OpalLink,
		updatedModule.GruppengroesseVorlesung,
		updatedModule.GruppengroesseUebung,
		updatedModule.GruppengroessePraktikum,
		updatedModule.Lehrinhalte,
		updatedModule.SozialUndSelbstkompetenzen,
		updatedModule.Hinweise,
		updatedModule.SelbststudienzeitAufschluesselung,
		updatedModule.AktuelleLehrressourcen,
		updatedModule.Literatur,
		updatedModule.ParentModulKuerzel,
		updatedModule.ParentModulVersion,
		updatedModule.FakultaetID,
		updatedModule.StudienrichtungID,
		updatedModule.VertiefungID,
		vorherigerZustandID,
		kuerzel,
		version,
	)
	if err != nil {
		return err
	}

	return nil
}

// ResetModuleToPreviousState resets the last change of a module
func ResetModuleToPreviousState(c *gin.Context) {
	kuerzel := c.Param("kuerzel")
	version, err := strconv.Atoi(c.Param("version"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Version parameter must be a valid integer"})
		return
	}

	ctx := context.Background()
	tx, err := database.DB.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer tx.Rollback(ctx)

	// Get the ID of the previous state
	var previousStateID int
	err = tx.QueryRow(ctx, `
        SELECT vorheriger_zustand_id
        FROM modul
        WHERE kuerzel = $1 AND version = $2
    `, kuerzel, version).Scan(&previousStateID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the previous module from the history
	var previousModule models.Module
	err = tx.QueryRow(ctx, `
        SELECT kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_integrale, sprachenzentrum, opal_link,
            gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
            besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
            praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, literatur,
            parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id, vorheriger_zustand_id
        FROM modul_historie
        WHERE id = $1
    `, previousStateID).Scan(
		&previousModule.Kuerzel,
		&previousModule.Version,
		&previousModule.FruehererSchluessel,
		&previousModule.Modultitel,
		&previousModule.ModultitelEnglisch,
		&previousModule.Kommentar,
		&previousModule.Niveau,
		&previousModule.Dauer,
		&previousModule.Turnus,
		&previousModule.StudiumIntegrale,
		&previousModule.Sprachenzentrum,
		&previousModule.OpalLink,
		&previousModule.GruppengroesseVorlesung,
		&previousModule.GruppengroesseUebung,
		&previousModule.GruppengroessePraktikum,
		&previousModule.Lehrform,
		&previousModule.Medienform,
		&previousModule.Lehrinhalte,
		&previousModule.Qualifikationsziele,
		&previousModule.SozialUndSelbstkompetenzen,
		&previousModule.BesondereZulassungsvoraussetzungen,
		&previousModule.EmpfohleneVoraussetzungen,
		&previousModule.Fortsetzungsmoeglichkeiten,
		&previousModule.Hinweise,
		&previousModule.EctsCredits,
		&previousModule.PraesenzeitWocheVorlesung,
		&previousModule.PraesenzeitWocheUebung,
		&previousModule.PraesenzeitWochePraktikum,
		&previousModule.PraesenzeitWocheSonstiges,
		&previousModule.SelbststudienzeitAufschluesselung,
		&previousModule.AktuelleLehrressourcen,
		&previousModule.Literatur,
		&previousModule.ParentModulKuerzel,
		&previousModule.ParentModulVersion,
		&previousModule.FakultaetID,
		&previousModule.StudienrichtungID,
		&previousModule.VertiefungID,
		&previousModule.VorherigerZustandID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the module to the previous state
	_, err = tx.Exec(ctx, `
        UPDATE modul
        SET 
            frueherer_schluessel = $1,
            modultitel = $2,
            modultitel_englisch = $3,
            kommentar = $4,
            niveau = $5,
            dauer = $6,
            turnus = $7,
            studium_integrale = $8,
            sprachenzentrum = $9,
            opal_link = $10,
            gruppengroesse_vorlesung = $11,
            gruppengroesse_uebung = $12,
            gruppengroesse_praktikum = $13,
            lehrform = $14,
            medienform = $15,
            lehrinhalte = $16,
            qualifikationsziele = $17,
            sozial_und_selbstkompetenzen = $18,
            besondere_zulassungsvoraussetzungen = $19,
            empfohlene_voraussetzungen = $20,
            fortsetzungsmoeglichkeiten = $21,
            hinweise = $22,
            ects_credits = $23,
            praesenzeit_woche_vorlesung = $24,
            praesenzeit_woche_uebung = $25,
            praesenzeit_woche_praktikum = $26,
            praesenzeit_woche_sonstiges = $27,
            selbststudienzeit_aufschluesselung = $28,
            aktuelle_lehrressourcen = $29,
            literatur = $30,
            parent_modul_kuerzel = $31,
            parent_modul_version = $32,
            fakultaet_id = $33,
            studienrichtung_id = $34,
            vertiefung_id = $35
			
        WHERE kuerzel = $36 AND version = $37
    `,
		previousModule.FruehererSchluessel,
		previousModule.Modultitel,
		previousModule.ModultitelEnglisch,
		previousModule.Kommentar,
		previousModule.Niveau,
		previousModule.Dauer,
		previousModule.Turnus,
		previousModule.StudiumIntegrale,
		previousModule.Sprachenzentrum,
		previousModule.OpalLink,
		previousModule.GruppengroesseVorlesung,
		previousModule.GruppengroesseUebung,
		previousModule.GruppengroessePraktikum,
		previousModule.Lehrform,
		previousModule.Medienform,
		previousModule.Lehrinhalte,
		previousModule.Qualifikationsziele,
		previousModule.SozialUndSelbstkompetenzen,
		previousModule.BesondereZulassungsvoraussetzungen,
		previousModule.EmpfohleneVoraussetzungen,
		previousModule.Fortsetzungsmoeglichkeiten,
		previousModule.Hinweise,
		previousModule.EctsCredits,
		previousModule.PraesenzeitWocheVorlesung,
		previousModule.PraesenzeitWocheUebung,
		previousModule.PraesenzeitWochePraktikum,
		previousModule.PraesenzeitWocheSonstiges,
		previousModule.SelbststudienzeitAufschluesselung,
		previousModule.AktuelleLehrressourcen,
		previousModule.Literatur,
		previousModule.ParentModulKuerzel,
		previousModule.ParentModulVersion,
		previousModule.FakultaetID,
		previousModule.StudienrichtungID,
		previousModule.VertiefungID,
		previousModule.VorherigerZustandID,
		kuerzel,
		version,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Module successfully reset to previous state"})
}

// GetUserRoles retrieves the roles of the logged-in user
func GetUserRoles(c *gin.Context) {
	personID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	query := `
        SELECT r.bezeichnung
        FROM modul_person_rolle mpr
        JOIN rolle r ON mpr.rolle_id = r.rolle_id
        WHERE mpr.person_id = $1
    `

	rows, err := database.DB.Query(context.Background(), query, personID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query roles"})
		return
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan role"})
			return
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}
