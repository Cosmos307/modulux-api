package models

import (
	"github.com/guregu/null"
)

type Module struct {
	Kuerzel                            string      `json:"kuerzel" binding:"required"`
	Version                            int         `json:"version" binding:"required"`
	FruehererSchluessel                null.String `json:"frueherer_schluessel"`
	Modultitel                         string      `json:"modultitel" binding:"required"`
	ModultitelEnglisch                 null.String `json:"modultitel_englisch"`
	Kommentar                          null.String `json:"kommentar"`
	Niveau                             string      `json:"niveau" binding:"required"`
	Dauer                              int         `json:"dauer" binding:"required"`
	Turnus                             string      `json:"turnus" binding:"required"`
	StudiumIntegrale                   bool        `json:"studium_integrale"`
	Sprachenzentrum                    bool        `json:"sprachenzentrum"`
	OpalLink                           null.String `json:"opal_link"`
	GruppengroesseVorlesung            null.Int    `json:"gruppengroesse_vorlesung"`
	GruppengroesseUebung               null.Int    `json:"gruppengroesse_uebung"`
	GruppengroessePraktikum            null.Int    `json:"gruppengroesse_praktikum"`
	Lehrform                           null.String `json:"lehrform"`
	Medienform                         null.String `json:"medienform"`
	Lehrinhalte                        null.String `json:"lehrinhalte"`
	Qualifikationsziele                null.String `json:"qualifikationsziele"`
	SozialUndSelbstkompetenzen         null.String `json:"sozial_und_selbstkompetenzen"`
	BesondereZulassungsvoraussetzungen null.String `json:"besondere_zulassungsvoraussetzungen"`
	EmpfohleneVoraussetzungen          null.String `json:"empfohlene_voraussetzungen"`
	Fortsetzungsmoeglichkeiten         null.String `json:"fortsetzungsmoeglichkeiten"`
	Hinweise                           null.String `json:"hinweise"`
	EctsCredits                        float64     `json:"ects_credits" binding:"required"`
	Workload                           int         `json:"workload"`
	PraesenzeitWocheVorlesung          float64     `json:"praesenzeit_woche_vorlesung"`
	PraesenzeitWocheUebung             float64     `json:"praesenzeit_woche_uebung"`
	PraesenzeitWochePraktikum          float64     `json:"praesenzeit_woche_praktikum"`
	PraesenzeitWocheSonstiges          float64     `json:"praesenzeit_woche_sonstiges"`
	Selbststudienzeit                  int         `json:"selbststudienzeit"`
	SelbststudienzeitAufschluesselung  null.String `json:"selbststudienzeit_aufschluesselung"`
	AktuelleLehrressourcen             null.String `json:"aktuelle_lehrressourcen"`
	Literatur                          null.String `json:"literatur"`
	ParentModulKuerzel                 null.String `json:"parent_modul_kuerzel"`
	ParentModulVersion                 null.Int    `json:"parent_modul_version"`
	FakultaetID                        null.Int    `json:"fakultaet_id"`
	StudienrichtungID                  null.Int    `json:"studienrichtung_id"`
	VertiefungID                       null.Int    `json:"vertiefung_id"`
	VorherigerZustandID                null.Int    `json:"vorheriger_zustand_id"`
}
