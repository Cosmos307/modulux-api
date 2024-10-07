package models

import (
	"database/sql"
	"encoding/json"
)

type Person struct {
	PersonID      int            `json:"person_id"`
	Titel         sql.NullString `json:"titel"`
	Vorname       string         `json:"vorname" binding:"required"`
	Nachname      string         `json:"nachname" binding:"required"`
	Email         string         `json:"email" binding:"required,email"`
	Telefonnummer sql.NullString `json:"telefonnummer"`
	Raum          sql.NullString `json:"raum"`
	Funktion      sql.NullString `json:"funktion" binding:"required"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	type Alias Person
	return json.Marshal(&struct {
		Titel         string `json:"titel"`
		Telefonnummer string `json:"telefonnummer"`
		Raum          string `json:"raum"`
		Funktion      string `json:"funktion"`
		Alias
	}{
		Titel:         nullStringToString(p.Titel),
		Telefonnummer: nullStringToString(p.Telefonnummer),
		Raum:          nullStringToString(p.Raum),
		Funktion:      nullStringToString(p.Funktion),
		Alias:         (Alias)(p),
	})
}

func (p *Person) UnmarshalJSON(data []byte) error {
	type Alias Person
	aux := &struct {
		Titel         *string `json:"titel"`
		Telefonnummer *string `json:"telefonnummer"`
		Raum          *string `json:"raum"`
		Funktion      *string `json:"funktion"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	p.Titel = stringToNullString(aux.Titel)
	p.Telefonnummer = stringToNullString(aux.Telefonnummer)
	p.Raum = stringToNullString(aux.Raum)
	p.Funktion = stringToNullString(aux.Funktion)

	return nil
}

func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func stringToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{Valid: false}
}
