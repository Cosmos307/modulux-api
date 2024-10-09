package models

import (
	"github.com/guregu/null"
)

type Studiengang struct {
	StudiengangID              int         `json:"studiengang_id"`
	Kuerzel                    string      `json:"kuerzel"`
	NummernImStudienablaufplan null.String `json:"nummern_im_studienablaufplan"`
	Studiengangstitel          string      `json:"studiengangstitel"`
	StudiengangstitelEnglisch  null.String `json:"studiengangstitel_englisch"`
	Kommentar                  null.String `json:"kommentar"`
	Abschluss                  string      `json:"abschluss"`
	ErsteImmatrikulation       null.String `json:"erste_immatrikulation"`
	ErforderlicheCredits       int         `json:"erforderliche_credits"`
	Kapazitaet                 int         `json:"kapazitaet"`
	InVollzeitStudierbar       bool        `json:"in_vollzeit_studierbar"`
	InTeilzeitStudierbar       bool        `json:"in_teilzeit_studierbar"`
	KooperativerStudiengang    bool        `json:"kooperativer_studiengang"`
	Doppelabschlussprogramm    bool        `json:"doppelabschlussprogramm"`
	Fernstudium                bool        `json:"fernstudium"`
	Englischsprachig           bool        `json:"englischsprachig"`
	Teasertext                 null.String `json:"teasertext"`
	Mobilitaetsfenster         null.String `json:"mobilitaetsfenster"`
	Website                    null.String `json:"website"`
	Imagefilm                  null.String `json:"imagefilm"`
	Werbeflyer                 null.String `json:"werbeflyer"`
	Akkreditierungsurkunde     null.String `json:"akkreditierungsurkunde"`
	FakultaetID                int         `json:"fakultaet_id"`
}
