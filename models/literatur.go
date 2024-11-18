package models

import (
	"github.com/guregu/null"
)

type Literatur struct {
	LiteraturID null.Int    `json:"literatur_id"`
	Titel       null.String `json:"titel"`
	Autor       null.String `json:"autor"`
	Jahr        null.Int    `json:"jahr"`
	Verlag      null.String `json:"verlag"`
	ISBN        null.String `json:"isbn"`
	Link        null.String `json:"link"`
	DOI         null.String `json:"doi"`
}
