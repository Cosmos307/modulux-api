package models

import (
	"github.com/guregu/null"
)

type Person struct {
	PersonID      int         `json:"person_id"`
	Titel         null.String `json:"titel"`
	Vorname       string      `json:"vorname" `
	Nachname      string      `json:"nachname"`
	Email         string      `json:"email" `
	Telefonnummer null.String `json:"telefonnummer"`
	Raum          null.String `json:"raum"`
	Funktion      null.String `json:"funktion"`
	Password      string      `json:"password"`
}
