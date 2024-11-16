package controllers

import (
	"context"
	"database/sql"
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
			&module.Kommentar, &module.Niveau, &module.Dauer, &module.Turnus, &module.StudiumGenerale, &module.Sprachenzentrum,
			&module.OpalLink, &module.GruppengroesseVorlesung, &module.GruppengroesseUebung, &module.GruppengroessePraktikum,
			&module.Lehrform, &module.Medienform, &module.Lehrinhalte, &module.Qualifikationsziele, &module.SozialUndSelbstkompetenzen,
			&module.BesondereZulassungsvoraussetzungen, &module.EmpfohleneVoraussetzungen, &module.Fortsetzungsmoeglichkeiten,
			&module.Hinweise, &module.EctsCredits, &module.Workload, &module.PraesenzeitWocheVorlesung, &module.PraesenzeitWocheUebung,
			&module.PraesenzeitWochePraktikum, &module.PraesenzeitWocheSonstiges, &module.Selbststudienzeit, &module.SelbststudienzeitAufschluesselung,
			&module.AktuelleLehrressourcen, &module.ParentModulKuerzel, &module.ParentModulVersion, &module.FakultaetID,
			&module.StudienrichtungID, &module.VertiefungID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		modules = append(modules, module)

		// Abfrage der Literaturdaten für das Modul
		literaturQuery := `
	 		SELECT l.literatur_id, l.titel, l.autor, l.jahr, l.verlag, l.isbn, l.link, l.doi
	 		FROM literatur l
			JOIN modul_literatur ml ON l.literatur_id = ml.literatur_id
			WHERE ml.modul_kuerzel = $1 AND ml.modul_version = $2
		`
		literaturRows, err := database.DB.Query(context.Background(), literaturQuery, module.Kuerzel, module.Version)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer literaturRows.Close()

		for literaturRows.Next() {
			var literatur models.Literatur
			err := literaturRows.Scan(
				&literatur.LiteraturID, &literatur.Titel, &literatur.Autor, &literatur.Jahr, &literatur.Verlag,
				&literatur.ISBN, &literatur.Link, &literatur.DOI,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			module.Literatur = append(module.Literatur, literatur)
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
		&module.Kommentar, &module.Niveau, &module.Dauer, &module.Turnus, &module.StudiumGenerale, &module.Sprachenzentrum,
		&module.OpalLink, &module.GruppengroesseVorlesung, &module.GruppengroesseUebung, &module.GruppengroessePraktikum,
		&module.Lehrform, &module.Medienform, &module.Lehrinhalte, &module.Qualifikationsziele, &module.SozialUndSelbstkompetenzen,
		&module.BesondereZulassungsvoraussetzungen, &module.EmpfohleneVoraussetzungen, &module.Fortsetzungsmoeglichkeiten,
		&module.Hinweise, &module.EctsCredits, &module.Workload, &module.PraesenzeitWocheVorlesung, &module.PraesenzeitWocheUebung,
		&module.PraesenzeitWochePraktikum, &module.PraesenzeitWocheSonstiges, &module.Selbststudienzeit, &module.SelbststudienzeitAufschluesselung,
		&module.AktuelleLehrressourcen, &module.ParentModulKuerzel, &module.ParentModulVersion, &module.FakultaetID,
		&module.StudienrichtungID, &module.VertiefungID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Abfrage der Literaturdaten für das Modul
	literaturQuery := `
		SELECT l.literatur_id, l.titel, l.autor, l.jahr, l.verlag, l.isbn, l.link, l.doi
		FROM literatur l
		JOIN modul_literatur ml ON l.literatur_id = ml.literatur_id
		WHERE ml.modul_kuerzel = $1 AND ml.modul_version = $2
	`
	literaturRows, err := database.DB.Query(context.Background(), literaturQuery, module.Kuerzel, module.Version)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer literaturRows.Close()

	for literaturRows.Next() {
		var literatur models.Literatur
		err := literaturRows.Scan(
			&literatur.LiteraturID, &literatur.Titel, &literatur.Autor, &literatur.Jahr, &literatur.Verlag,
			&literatur.ISBN, &literatur.Link, &literatur.DOI,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		module.Literatur = append(module.Literatur, literatur)
	}

	c.JSON(http.StatusOK, module)
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

	personID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(context.Background())
			panic(p)
		} else if err != nil {
			tx.Rollback(context.Background())
		} else {
			err = tx.Commit(context.Background())
		}
	}()

	query := `INSERT INTO modul (
	kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_generale, 
	sprachenzentrum, opal_link, gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, 
	qualifikationsziele, sozial_und_selbstkompetenzen, besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, 
	hinweise, ects_credits, praesenzeit_woche_vorlesung, praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, 
	selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id)  
    VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36
	)`
	_, err = tx.Exec(context.Background(), query,
		newModule.Kuerzel, newModule.Version, newModule.FruehererSchluessel, newModule.Modultitel, newModule.ModultitelEnglisch,
		newModule.Kommentar, newModule.Niveau, newModule.Dauer, newModule.Turnus, newModule.StudiumGenerale, newModule.Sprachenzentrum,
		newModule.OpalLink, newModule.GruppengroesseVorlesung, newModule.GruppengroesseUebung, newModule.GruppengroessePraktikum,
		newModule.Lehrform, newModule.Medienform, newModule.Lehrinhalte, newModule.Qualifikationsziele, newModule.SozialUndSelbstkompetenzen,
		newModule.BesondereZulassungsvoraussetzungen, newModule.EmpfohleneVoraussetzungen, newModule.Fortsetzungsmoeglichkeiten,
		newModule.Hinweise, newModule.EctsCredits, newModule.PraesenzeitWocheVorlesung, newModule.PraesenzeitWocheUebung,
		newModule.PraesenzeitWochePraktikum, newModule.PraesenzeitWocheSonstiges, newModule.SelbststudienzeitAufschluesselung,
		newModule.AktuelleLehrressourcen, newModule.ParentModulKuerzel, newModule.ParentModulVersion, newModule.FakultaetID,
		newModule.StudienrichtungID, newModule.VertiefungID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	roleQuery := `
        INSERT INTO modul_person_rolle (modul_kuerzel, modul_version, person_id, rolle_id)
        VALUES ($1, $2, $3, (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher'))
    `
	_, err = tx.Exec(context.Background(), roleQuery, newModule.Kuerzel, newModule.Version, personID)
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
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(context.Background())
			panic(p)
		} else if err != nil {
			tx.Rollback(context.Background())
		} else {
			err = tx.Commit(context.Background())
		}
	}()

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

		personIDStr, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
			return
		}

		personID, err := strconv.Atoi(personIDStr.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
			return
		}

		newVersion := maxVersion + 1
		newModule := updatedModule
		newModule.Version = newVersion

		err = InsertNewVersionModule(tx, kuerzel, newVersion, newModule, personID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	} else if hasMinorChanges(currentModule, updatedModule) {
		previousHistoryID, err := saveModuleHistory(tx, kuerzel, version)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

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

// InsertNewModule inserts a new module into the database
func InsertNewVersionModule(tx pgx.Tx, kuerzel string, version int, module models.Module, personID int) error {
	query := `
        INSERT INTO modul (
            kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_generale, sprachenzentrum, opal_link,
            gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
            besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
            praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen,
            parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36
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
		module.StudiumGenerale,
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
		module.ParentModulKuerzel,
		module.ParentModulVersion,
		module.FakultaetID,
		module.StudienrichtungID,
		module.VertiefungID,
	)
	if err != nil {
		return fmt.Errorf("failed to insert module: %w", err)
	}

	// Assign role "Modulverantwortlicher" to the person who initiated the significant change
	roleQuery := `
        INSERT INTO modul_person_rolle (modul_kuerzel, modul_version, person_id, rolle_id)
        VALUES ($1, $2, $3, (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher'))
    `
	_, err = tx.Exec(context.Background(), roleQuery, kuerzel, version, personID)

	if err != nil {
		return fmt.Errorf("failed to assign role to person: %w", err)
	}

	return nil
}

// saveModuleHistory saves the current state of a module to the modul_historie table
func saveModuleHistory(tx pgx.Tx, kuerzel string, version int) (int, error) {
	var historyID int
	query := `
        INSERT INTO modul_historie (
            kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_generale, sprachenzentrum, opal_link,
            gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
            besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
            praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, 
            parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id, vorheriger_zustand_id
        )
        SELECT
            kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_generale, sprachenzentrum, opal_link,
            gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
            besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
            praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen, 
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
		SELECT kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_generale, sprachenzentrum, opal_link,
			gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
			besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
			praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen,
			parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id, vorheriger_zustand_id
		FROM modul
		WHERE kuerzel = $1 AND version = $2
	`
	err := tx.QueryRow(context.Background(), query, kuerzel, version).Scan(
		&module.Kuerzel, &module.Version, &module.FruehererSchluessel, &module.Modultitel, &module.ModultitelEnglisch, &module.Kommentar,
		&module.Niveau, &module.Dauer, &module.Turnus, &module.StudiumGenerale, &module.Sprachenzentrum, &module.OpalLink,
		&module.GruppengroesseVorlesung, &module.GruppengroesseUebung, &module.GruppengroessePraktikum, &module.Lehrform, &module.Medienform,
		&module.Lehrinhalte, &module.Qualifikationsziele, &module.SozialUndSelbstkompetenzen, &module.BesondereZulassungsvoraussetzungen,
		&module.EmpfohleneVoraussetzungen, &module.Fortsetzungsmoeglichkeiten, &module.Hinweise, &module.EctsCredits, &module.PraesenzeitWocheVorlesung,
		&module.PraesenzeitWocheUebung, &module.PraesenzeitWochePraktikum, &module.PraesenzeitWocheSonstiges, &module.Selbststudienzeit,
		&module.SelbststudienzeitAufschluesselung, &module.AktuelleLehrressourcen, &module.ParentModulKuerzel, &module.ParentModulVersion,
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
		updatedModule.StudiumGenerale != currentModule.StudiumGenerale ||
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
            studium_generale = $7,
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
            parent_modul_kuerzel = $18,
            parent_modul_version = $19,
            fakultaet_id = $20,
            studienrichtung_id = $21,
            vertiefung_id = $22,
            vorheriger_zustand_id = $23
        WHERE kuerzel = $24 AND version = $25
    `

	_, err := tx.Exec(context.Background(), query,
		updatedModule.FruehererSchluessel,
		updatedModule.Modultitel,
		updatedModule.ModultitelEnglisch,
		updatedModule.Kommentar,
		updatedModule.Niveau,
		updatedModule.Turnus,
		updatedModule.StudiumGenerale,
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
func RollbackLastChange(c *gin.Context) {
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
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var lastChangeType string
	err = tx.QueryRow(ctx, `
        SELECT change_type
        FROM (
            SELECT 'module' AS change_type, aenderungsdatum
            FROM modul_historie
            WHERE kuerzel = $1 AND version = $2
            UNION ALL
            SELECT 'literature' AS change_type, aenderungsdatum
            FROM modul_literatur_historie
            WHERE modul_kuerzel = $1 AND modul_version = $2
        ) AS combined_changes
        ORDER BY aenderungsdatum DESC
        LIMIT 1
    `, kuerzel, version).Scan(&lastChangeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get last change type"})
		return
	}

	if lastChangeType == "module" {
		// Rollback module change
		err = rollbackModuleChange(tx, kuerzel, version)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rollback module change"})
			return
		}
	} else if lastChangeType == "literature" {
		// Rollback literature change
		err = rollbackLiteratureChange(tx, kuerzel, version)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rollback literature change", "details": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Module or literature successfully reset to previous state"})
}

// rollbackModuleChange rollbacks the last change of a module
func rollbackModuleChange(tx pgx.Tx, kuerzel string, version int) error {
	// Get the ID of the previous state
	var previousStateID sql.NullInt32
	err := tx.QueryRow(context.Background(), `
        SELECT vorheriger_zustand_id
        FROM modul
        WHERE kuerzel = $1 AND version = $2
    `, kuerzel, version).Scan(&previousStateID)
	if err != nil {
		return fmt.Errorf("failed to get previous state ID: %w", err)
	}

	if !previousStateID.Valid {
		return fmt.Errorf("no previous state found for module %s version %d", kuerzel, version)
	}

	// Get the previous module from the history
	var previousModule models.Module
	err = tx.QueryRow(context.Background(), `
        SELECT kuerzel, version, frueherer_schluessel, modultitel, modultitel_englisch, kommentar, niveau, dauer, turnus, studium_generale, sprachenzentrum, opal_link,
            gruppengroesse_vorlesung, gruppengroesse_uebung, gruppengroesse_praktikum, lehrform, medienform, lehrinhalte, qualifikationsziele, sozial_und_selbstkompetenzen,
            besondere_zulassungsvoraussetzungen, empfohlene_voraussetzungen, fortsetzungsmoeglichkeiten, hinweise, ects_credits, praesenzeit_woche_vorlesung,
            praesenzeit_woche_uebung, praesenzeit_woche_praktikum, praesenzeit_woche_sonstiges, selbststudienzeit, selbststudienzeit_aufschluesselung, aktuelle_lehrressourcen,
            parent_modul_kuerzel, parent_modul_version, fakultaet_id, studienrichtung_id, vertiefung_id, vorheriger_zustand_id
        FROM modul_historie
        WHERE id = $1
    `, previousStateID.Int32).Scan(
		&previousModule.Kuerzel, &previousModule.Version, &previousModule.FruehererSchluessel, &previousModule.Modultitel, &previousModule.ModultitelEnglisch, &previousModule.Kommentar,
		&previousModule.Niveau, &previousModule.Dauer, &previousModule.Turnus, &previousModule.StudiumGenerale, &previousModule.Sprachenzentrum, &previousModule.OpalLink,
		&previousModule.GruppengroesseVorlesung, &previousModule.GruppengroesseUebung, &previousModule.GruppengroessePraktikum, &previousModule.Lehrform, &previousModule.Medienform,
		&previousModule.Lehrinhalte, &previousModule.Qualifikationsziele, &previousModule.SozialUndSelbstkompetenzen, &previousModule.BesondereZulassungsvoraussetzungen,
		&previousModule.EmpfohleneVoraussetzungen, &previousModule.Fortsetzungsmoeglichkeiten, &previousModule.Hinweise, &previousModule.EctsCredits, &previousModule.PraesenzeitWocheVorlesung,
		&previousModule.PraesenzeitWocheUebung, &previousModule.PraesenzeitWochePraktikum, &previousModule.PraesenzeitWocheSonstiges, &previousModule.Selbststudienzeit,
		&previousModule.SelbststudienzeitAufschluesselung, &previousModule.AktuelleLehrressourcen, &previousModule.ParentModulKuerzel, &previousModule.ParentModulVersion,
		&previousModule.FakultaetID, &previousModule.StudienrichtungID, &previousModule.VertiefungID, &previousModule.VorherigerZustandID,
	)
	if err != nil {
		return fmt.Errorf("failed to get previous module state: %w", err)
	}

	// Update the module to the previous state
	_, err = tx.Exec(context.Background(), `
        UPDATE modul
        SET 
            frueherer_schluessel = $1,
            modultitel = $2,
            modultitel_englisch = $3,
            kommentar = $4,
            niveau = $5,
            turnus = $6,
            studium_generale = $7,
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
            parent_modul_kuerzel = $18,
            parent_modul_version = $19,
            fakultaet_id = $20,
            studienrichtung_id = $21,
            vertiefung_id = $22,
            vorheriger_zustand_id = $23
        WHERE kuerzel = $24 AND version = $25
    `, previousModule.FruehererSchluessel, previousModule.Modultitel, previousModule.ModultitelEnglisch, previousModule.Kommentar,
		previousModule.Niveau, previousModule.Turnus, previousModule.StudiumGenerale, previousModule.Sprachenzentrum, previousModule.OpalLink,
		previousModule.GruppengroesseVorlesung, previousModule.GruppengroesseUebung, previousModule.GruppengroessePraktikum, previousModule.Lehrinhalte,
		previousModule.SozialUndSelbstkompetenzen, previousModule.Hinweise, previousModule.SelbststudienzeitAufschluesselung, previousModule.AktuelleLehrressourcen,
		previousModule.ParentModulKuerzel, previousModule.ParentModulVersion, previousModule.FakultaetID, previousModule.StudienrichtungID, previousModule.VertiefungID,
		previousModule.VorherigerZustandID, kuerzel, version)
	if err != nil {
		return fmt.Errorf("failed to update module to previous state: %w", err)
	}

	// Delete the copied last state from modul_historie
	_, err = tx.Exec(context.Background(), `
		DELETE FROM modul_historie
		WHERE id = $1
	`, previousStateID.Int32)
	if err != nil {
		return fmt.Errorf("failed to delete previous module state from history: %w", err)
	}

	return nil
}

// rollbackLiteratureChange rollbacks the last change of a module's literature
func rollbackLiteratureChange(tx pgx.Tx, kuerzel string, version int) error {
	// Get the ID of the previous snapshot
	var previousSnapshotID sql.NullInt32
	err := tx.QueryRow(context.Background(), `
        SELECT vorheriger_snapshot_id
        FROM modul_literatur
        WHERE modul_kuerzel = $1 AND modul_version = $2
    `, kuerzel, version).Scan(&previousSnapshotID)
	if err != nil {
		fmt.Println("Error getting previous snapshot ID:", err)
		return fmt.Errorf("failed to get previous snapshot ID: %w", err)
	}

	if !previousSnapshotID.Valid {
		return fmt.Errorf("no previous snapshot found for module %s version %d", kuerzel, version)
	}

	// Delete the current literature references
	_, err = tx.Exec(context.Background(), `
        DELETE FROM modul_literatur
        WHERE modul_kuerzel = $1 AND modul_version = $2
    `, kuerzel, version)
	if err != nil {
		fmt.Println("Error deleting current literature references:", err)
		return fmt.Errorf("failed to delete current literature references: %w", err)
	}

	// Insert the literature references from the previous snapshot
	_, err = tx.Exec(context.Background(), `
        INSERT INTO modul_literatur (modul_kuerzel, modul_version, literatur_id, vorheriger_snapshot_id)
        SELECT modul_kuerzel, modul_version, literatur_id, 
        CASE WHEN vorheriger_snapshot_id IS NULL THEN NULL ELSE vorheriger_snapshot_id END
        FROM modul_literatur_historie
        WHERE modul_kuerzel = $1 AND modul_version = $2 AND snapshot_id = $3
    `, kuerzel, version, previousSnapshotID.Int32)
	if err != nil {
		fmt.Println("Error inserting literature references from previous snapshot:", err)
		return fmt.Errorf("failed to insert literature references from previous snapshot: %w", err)
	}

	// Delete the literature references from the history table
	_, err = tx.Exec(context.Background(), `
        DELETE FROM modul_literatur_historie
        WHERE modul_kuerzel = $1 AND modul_version = $2 AND snapshot_id = $3
    `, kuerzel, version, previousSnapshotID.Int32)
	if err != nil {
		fmt.Println("Error deleting literature references from history table:", err)
		return fmt.Errorf("failed to delete literature references from history table: %w", err)
	}

	return nil
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
