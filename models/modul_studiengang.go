package models

import (
	"github.com/guregu/null"
)

type ModulStudiengang struct {
	ModulKuerzel  string   `json:"modul_kuerzel"`
	ModulVersion  int      `json:"modul_version"`
	StudiengangID int      `json:"studiengang_id"`
	Semester      null.Int `json:"semester"`
	ModulTyp      string   `json:"modul_typ"`
}
