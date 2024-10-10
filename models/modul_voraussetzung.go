package models

type ModulVoraussetzung struct {
	ModulKuerzel                string `json:"modul_kuerzel"`
	ModulVersion                int    `json:"modul_version"`
	VorausgesetztesModulKuerzel string `json:"vorausgesetztes_modul_kuerzel"`
	VorausgesetztesModulVersion int    `json:"vorausgesetztes_modul_version"`
}
