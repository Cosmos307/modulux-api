# modulux-api

Die **modulux-api** ist eine prototypische RESTful API zur Verwaltung von Modulen, Studiengängen und deren Beziehungen und stellt den praktischen Teil meiner Bachelorarbeit dar. Sie entstand aus dem Bedarf, die bestehende monolithische Webanwendung der Moduldatenbank Modulux – bisher als TYPO3-Modul im Einsatz – durch eine moderne, flexible und erweiterbare Lösung abzulösen.

Modulux wird an der Hochschule für Technik, Wirtschaft und Kultur Leipzig (HTWK Leipzig) zur Organisation und Verwaltung von Studiengängen und Modulen eingesetzt. Die bisherige monolithische Architektur erschwert jedoch die Integration mit anderen Systemen, die Erweiterbarkeit sowie die Nachverfolgbarkeit von Änderungen. Zudem fehlen Schnittstellen für externe Zugriffe, eine zeitgemäße Versionsverwaltung sowie diverse Unterstützungen für Bearbeiter.

Mit der modulux-api wird eine Grundlage geschaffen, um diese Schwächen zu adressieren: Die API bietet klare Schnittstellen, unterstützt Versionierung, leistet Unterstützungen und Kontrollen für Datenkonsistenz, ermöglicht parallele Arbeitsstände und setzt ein Rollenkonzept zur gezielten Zugriffssteuerung um. Ziel ist es, die Verwaltung von Curricula effizienter, transparenter und zukunftssicher zu gestalten – sowohl für Hochschulen als auch für vergleichbare Bildungseinrichtungen.

## Architektur
![Softwarearchitektur](images/Softwarearchitektur-Diagramm.png)
![Datenbankarchitektur](images/Datenbank-Diagramm.png)


## Motivation

Die API adressiert typische Herausforderungen monolithischer Systeme:
- **Fehlende Schnittstellen:** Erschwerte Integration und Erweiterbarkeit.
- **Keine moderne Versionsverwaltung:** Eingeschränkte Nachverfolgbarkeit von Änderungen.
- **Eingeschränkte Arbeitsstände:** Begrenzte Möglichkeiten zur parallelen Bearbeitung und Sicherung.
- **Hoher Wartungsaufwand:** Geringe Flexibilität und schwierige Anpassbarkeit.
- **Keine Unterstützung:** Bearbeiter können im aktuellen Modulux System Freitext eintragen, ohne Verifikation und Kontrolle der Eingaben sind  menschliche Fehler vorprogrammiert.

Mit modulux-api sollen Verwaltungsprozesse effizienter, transparenter und zukunftssicher gestaltet werden.

## Features der programmierten API

- Verwaltung von Modulen (CRUD, Versionierung, Rollback)
- Verwaltung von Studiengängen und deren Modulen
- Verwaltung von Modul-Voraussetzungen inkl. Konsistenzprüfung und Vorschlagsfunktion von vorherigen Modulen
- Literaturverwaltung für Module mit automatisierten Literaturvorschlägen (CrossRef-API-Abfragen)
- Unterstützung bei der Formulierung von Qualifikationszielen (Taxonomie-Feedback, Verbvorschläge)
- Benutzer- und Rollenmanagement mit Authentifizierung (JWT) und Autorisierung 
- Umfangreiche Validierungen und Fehlerbehandlung
- Bereitstellung von Kalender-relevanten Modulinformationen (z.B. Opal-Links) für externe Anwendungen

## Projektstruktur
```
todo-app/
│
├── api/
│   ├── Dockerfile          # Docker-Build-Definition für das API-Backend
│   ├── go.mod, go.sum      # Go Modules Definition & Checksummen
│   ├── cmd/
│   │   └── server/
│   │       └── main.go     # Einstiegspunkt (main.go)
│   └── internal/
│       ├── config/
│       ├── database/       # Datenbankverbindung und -initialisierung      
│       ├── handlers/
│       ├── models/
│       ├── repository/
│       │   └── mocks/      # Mock-Repositories für Tests
│       ├── routes/         # Definition der API-Routen
│       └── tests/          # Unit- und Integrationstests
├── database/   	        # SQL-Skripte für Migrationen und Seed-Daten
├── docker-compose.yml      # Container-Orchestrierung
├── .gitignore              # Git Ignore-Datei
└── README.md
```
## Wichtige Komponenten

### Controller

- [`api/controllers/modul_controller.go`](modulux-api/api/controllers/modul_controller.go): Modulverwaltung, Versionierung, Rollback, Literatur
- [`api/controllers/modul_voraussetzung_controller.go`](modulux-api/api/controllers/modul_voraussetzung_controller.go): Verwaltung von Modul-Voraussetzungen
- [`api/controllers/modul_studiengang_controller.go`](modulux-api/api/controllers/modul_studiengang_controller.go): Zuordnung von Modulen zu Studiengängen
- [`api/controllers/studiengang_controller.go`](modulux-api/api/controllers/studiengang_controller.go): Studiengangverwaltung
- ...

### Middleware

- [`api/middleware/`](modulux-api/api/middleware/):  
  - Authentifizierung (Login, JWT-Token)
  - Autorisierung (z.B. Rollenüberprüfung für aufgerufene Endpunkte)
  - Logging und Fehlerbehandlung

### Models

- [`models/modul.go`](modulux-api/models/modul.go): Modulstruktur mit allen Eigenschaften
- Weitere Modelle für Studiengänge, Literatur, Benutzer etc.

### Routen

- [`api/routes/modul_routes.go`](modulux-api/api/routes/modul_routes.go): Endpunkte für Module
- [`api/routes/modul_voraussetzung_routes.go`](modulux-api/api/routes/modul_voraussetzung_routes.go): Endpunkte für Modul-Voraussetzungen
- ...

## Beispiel-Endpunkte

- `GET /modul/` – Alle Module abrufen
- `POST /modul/` – Neues Modul anlegen
- `PUT /modul/:kuerzel/:version` – Modul aktualisieren (mit Auth)
- `POST /modul/:kuerzel/:version/reset` – Letzte Änderung zurücksetzen (Rollback)
- `GET /modul_voraussetzungen/:studiengang_id/:modul_kuerzel/:modul_version` – Voraussetzungen eines Moduls abrufen

## Authentifizierung & Autorisierung

- Middleware prüft JWT-Token und Rollen
- Nur berechtigte Nutzer können bestimmte Aktionen (z.B. Modul bearbeiten) durchführen

## Datenbank

- PostgreSQL
- Migrationen und Seed-Daten in [`database/`](modulux-api/database/)

## Deployment

- Dockerfile und docker-compose für lokalen und produktiven Betrieb
- Umgebungsvariablen für Konfiguration

## Entwicklung & Tests

- Go Modules (`go.mod`)
- Unit- und Integrationstests (siehe Testverzeichnisse in den Controllern)
