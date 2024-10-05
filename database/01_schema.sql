CREATE TABLE IF NOT EXISTS rolle (
    rolle_id SERIAL PRIMARY KEY,
    bezeichnung VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS berechtigung (
    berechtigung_id SERIAL PRIMARY KEY,
    bezeichnung VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS rolle_berechtigung (
    rolle_id INT REFERENCES rolle(rolle_id) ON DELETE CASCADE,
    berechtigung_id INT REFERENCES berechtigung(berechtigung_id) ON DELETE CASCADE,
    PRIMARY KEY (rolle_id, berechtigung_id)
);

CREATE TABLE IF NOT EXISTS person (
    person_id SERIAL PRIMARY KEY,
    titel VARCHAR(255),
    vorname VARCHAR(255) NOT NULL,
    nachname VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    telefonnummer VARCHAR(255),
    raum VARCHAR(255),
    funktion VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS fakultaet (
    fakultaet_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    kuerzel VARCHAR(50) NOT NULL
);

CREATE TYPE abschluss_typ AS ENUM ('Bachelor', 'Master', 'Diplom');
CREATE TYPE studienniveau AS ENUM ('Bachelor', 'Master', 'Diplom');
CREATE TYPE semester_turnus AS ENUM ('Sommer', 'Winter', 'Sommer und Winter');


CREATE TABLE IF NOT EXISTS studiengang (
    studiengang_id INT PRIMARY KEY,
    kuerzel VARCHAR(10) NOT NULL UNIQUE,
    nummern_im_studienablaufplan VARCHAR(30) NOT NULL,
    studiengangstitel VARCHAR(50) NOT NULL,
    studiengangstitel_englisch VARCHAR(50),
    kommentar TEXT,
    abschluss abschluss_typ NOT NULL,
    erste_immatrikulation DATE,
    erforderliche_credits INT NOT NULL,
    kapazitaet INT NOT NULL,
    in_vollzeit_studierbar BOOLEAN DEFAULT FALSE,
    in_teilzeit_studierbar BOOLEAN DEFAULT FALSE,
    kooperativer_studiengang BOOLEAN DEFAULT FALSE,
    doppelabschlussprogramm BOOLEAN DEFAULT FALSE,
    fernstudium BOOLEAN DEFAULT FALSE,
    englischsprachig BOOLEAN DEFAULT FALSE,
    teasertext TEXT,
    mobilitaetsfenster VARCHAR(100),
    website VARCHAR(100),
    imagefilm VARCHAR(100),
    werbeflyer VARCHAR(100),
    akkreditierungsurkunde VARCHAR(100),

    fakultaet_id INT REFERENCES fakultaet(fakultaet_id)
);

CREATE TABLE IF NOT EXISTS studienrichtung (
    studienrichtung_id SERIAL PRIMARY KEY,
    bezeichnung VARCHAR(255) NOT NULL,
    studiengang_id INT REFERENCES studiengang(studiengang_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS vertiefung (
    vertiefung_id SERIAL PRIMARY KEY,
    bezeichnung VARCHAR(255) NOT NULL,
    studienrichtung_id INT REFERENCES studienrichtung(studienrichtung_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS block (
    block_id SERIAL PRIMARY KEY,
    bezeichnung VARCHAR(255) NOT NULL,
    beschreibung TEXT,
    studiengang_id INT REFERENCES studiengang(studiengang_id) ON DELETE CASCADE,
    studienrichtung_id INT REFERENCES studienrichtung(studienrichtung_id) ON DELETE CASCADE,
    vertiefung_id INT REFERENCES vertiefung(vertiefung_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS modul (
    kuerzel VARCHAR(6) NOT NULL,
    version INT NOT NULL,
    frueherer_schluessel VARCHAR(6),
    modultitel VARCHAR(100) NOT NULL,
    modultitel_englisch VARCHAR(100),
    kommentar TEXT,
    niveau studienniveau NOT NULL,
    dauer INT NOT NULL,
    turnus semester_turnus NOT NULL,
    studium_integrale BOOLEAN NOT NULL DEFAULT FALSE,   --
    sprachenzentrum BOOLEAN NOT NULL DEFAULT FALSE,     --
    opal_link VARCHAR(255),
    gruppengroesse_vorlesung INT,
    gruppengroesse_uebung INT,
    gruppengroesse_praktikum INT,                       --                  
    lehrform TEXT,
    medienform TEXT,
    lehrinhalte TEXT,
    qualifikationsziele TEXT,
    sozial_und_selbstkompetenzen TEXT,                  --
    besondere_zulassungsvoraussetzungen TEXT,
    empfohlene_voraussetzungen TEXT,    
    fortsetzungsmoeglichkeiten TEXT,                    --
    hinweise TEXT,
    ects_credits DECIMAL(3,1) NOT NULL,
    workload INT GENERATED ALWAYS AS (ROUND(ects_credits * 30)) STORED,
    praesenzeit_woche_vorlesung INT DEFAULT 0 NOT NULL,
    praesenzeit_woche_uebung INT DEFAULT 0 NOT NULL,
    praesenzeit_woche_praktikum INT DEFAULT 0 NOT NULL,
    praesenzeit_woche_sonstiges INT DEFAULT 0 NOT NULL,
    selbststudienzeit INT GENERATED ALWAYS AS (
        (ROUND(ects_credits * 30)) - (14 * (praesenzeit_woche_vorlesung + praesenzeit_woche_uebung + praesenzeit_woche_praktikum + praesenzeit_woche_sonstiges))
    ) STORED,
    selbststudienzeit_aufschluesselung TEXT,
    aktuelle_lehrressourcen TEXT,
    literatur TEXT,
    
    parent_modul_kuerzel VARCHAR(7),
    parent_modul_version INT,  

    fakultaet_id INT REFERENCES fakultaet(fakultaet_id),
    studienrichtung_id INT REFERENCES studienrichtung(studienrichtung_id) ON DELETE CASCADE,
    vertiefung_id INT REFERENCES vertiefung(vertiefung_id) ON DELETE CASCADE,
    FOREIGN KEY (parent_modul_kuerzel, parent_modul_version) REFERENCES modul(kuerzel, version) ON DELETE CASCADE,

    PRIMARY KEY (kuerzel, version)
);

CREATE TABLE IF NOT EXISTS modul_person_rolle (
    modul_kuerzel VARCHAR(6),
    modul_version INT,
    person_id INT REFERENCES person(person_id) ON DELETE CASCADE,
    rolle_id INT REFERENCES rolle(rolle_id) ON DELETE CASCADE,
    FOREIGN KEY (modul_kuerzel, modul_version) REFERENCES modul(kuerzel, version) ON DELETE CASCADE,
    PRIMARY KEY (modul_kuerzel, modul_version, person_id, rolle_id)
);

CREATE TABLE IF NOT EXISTS studiengang_person_rolle (
    studiengang_id INT,
    person_id INT REFERENCES person(person_id) ON DELETE CASCADE,
    rolle_id INT REFERENCES rolle(rolle_id) ON DELETE CASCADE,
    PRIMARY KEY (studiengang_id, person_id, rolle_id)
);


CREATE TYPE modul_typ_enum AS ENUM ('Wahlpflichtmodul', 'Pflichtmodul');

CREATE TABLE IF NOT EXISTS modul_studiengang (
    modul_kuerzel VARCHAR(6),
    modul_version INT,
    studiengang_id INT,
    semester INT,
    modul_typ modul_typ_enum NOT NULL,
    
    FOREIGN KEY (modul_kuerzel, modul_version) REFERENCES modul(kuerzel, version) ON DELETE CASCADE,

    PRIMARY KEY (modul_kuerzel, modul_version, studiengang_id)
);


CREATE TABLE IF NOT EXISTS modul_block (
    block_id INT REFERENCES block(block_id) ON DELETE CASCADE,
    modul_kuerzel VARCHAR(6),
    modul_version INT,
    FOREIGN KEY (modul_kuerzel, modul_version) REFERENCES modul(kuerzel, version) ON DELETE CASCADE,

    PRIMARY KEY (modul_kuerzel, modul_version, block_id)
);

