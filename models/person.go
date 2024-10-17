package models

import (
	"github.com/guregu/null"
)

type Person struct {
	PersonID      int         `json:"person_id"`
	Titel         null.String `json:"titel"`
	Vorname       string      `json:"vorname" binding:"required"`
	Nachname      string      `json:"nachname" binding:"required"`
	Email         string      `json:"email" binding:"required,email"`
	Telefonnummer null.String `json:"telefonnummer"`
	Raum          null.String `json:"raum"`
	Funktion      null.String `json:"funktion" binding:"required"`
}
