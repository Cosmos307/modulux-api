
-- Fakultäten
INSERT INTO fakultaet (name, kuerzel)
VALUES 
    ('Architektur und Sozialwissenschaften', 'AS'),
    ('Digitale Transformation', 'DIT'),
    ('Ingenieurwissenschaften', 'ING'),
    ('Informatik und Medien', 'IM'),
    ('Bauwesen', 'B'),
    ('Wirtschaftswissenschaft und Wirtschaftsingenieurwesen', 'WW');

-- Personen
INSERT INTO person (vorname, nachname, titel, email, telefonnummer, raum, funktion)
VALUES 
    ('Jean-Alexander', 'Müller', 'Prof. Dr.-Ing.', 'jean-alexander.mueller@htwk-leipzig.de', '+49 341 3076-6638', 'ZU 123', 'Studiendekan'),
    ('Sibylle', 'Schwarz', 'Prof. Dr. rer. nat.', 'sibylle.schwarz@htwk-leipzig.de', '+49 341 3076-6483', 'ZU 411', 'Professorin'),
    ('Jens', 'Wagner', 'Prof. Dr. rer. nat.', 'jens.wagner@htwk-leipzig.de', '+49 341 3076-6494', 'LI 015', 'Professor'),
    ('Hanna', 'Brodowsky', NULL, 'hanna.brodowsky@htwk-leipzig.de', NULL, NULL, 'Dozentin'),
    ('Mario', 'Hlawitschka', 'Prof. Dr. rer. nat.', 'mario.hlawitschka@htwk-leipzig.de', '+49 341 3076-6493', 'ZU 224', 'Professor'),
    ('Martin', 'Grüttmüller', 'Prof. Dr. rer. nat. habil.', 'martin.gruettmueller@htwk-leipzig.de', '+49 341 3076-6487', 'ZU 412', 'Professor'),
    ('Karsten', 'Weicker', 'Prof. Dr. rer. nat.', 'karsten.weicker@htwk-leipzig.de', '+49 341 3076-6395', 'ZU 410', 'Professor'),
    ('Thomas', 'Kudraß', 'Prof. Dr.-Ing.', 'thomas.kudrass@htwk-leipzig.de', '+49 341 3076-6420', 'ZU 130', 'Professor'),
    ('Thomas', 'Riechert', 'Prof. Dr. rer. nat.', 'thomas.riechert@htwk-leipzig.de', '+49 341 3076-6413', 'ZU 507', 'Professor'),
    ('Antje', 'Tober-Nietner', 'Dr.', 'antje.tober@htwk-leipzig.de', NULL, NULL, 'Dozentin'),
    ('Johannes', 'Waldmann', 'Prof. Dr. rer. nat.', 'johannes.waldmann@htwk-leipzig.de', '+49 341 3076-6479', 'ZU 129', 'Professor');

-- Rollen 
INSERT INTO rolle (bezeichnung)
VALUES 
    ('Studiengangverantwortlicher'),
    ('Studiengangbearbeiter'),
    ('Modulverantwortlicher'),
    ('Modulbearbeiter'),
    ('Dozent'),
    ('Moduluxverantwortlicher'),
    ('Prozesskontrolle'),
    ('Hochschulleitung');

-- Studiengang 
INSERT INTO studiengang (studiengang_id, kuerzel, nummern_im_studienablaufplan, studiengangstitel, studiengangstitel_englisch, kommentar, abschluss, erste_immatrikulation, erforderliche_credits, kapazitaet, in_vollzeit_studierbar, in_teilzeit_studierbar, fakultaet_id, teasertext, mobilitaetsfenster, website)
VALUES 
    (1, '20INB', '1125153', 'Informatik | Bachelor', 'Computer Science | Bachelor', 'Kommentar 20INB', 'Bachelor', '2020-01-01', 180, 120, TRUE, FALSE, (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), 'Teasertext 20INB', '5. Semester', 'www.htwk-leipzig.de/inb'),
    (2, '20MIB', '1234214', 'Medieninformatik | Bachelor', 'Media Informatics | Bachelor', 'Kommentar 20MNB', 'Bachelor', '2020-01-01', 180, 120, TRUE, FALSE, (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), 'Teasertext 20MNB', '5. Semester', 'www.htwk-leipzig.de/mib');

-- Person dem Studiengang als Verantwortlicher zuordnen
INSERT INTO studiengang_person_rolle (studiengang_id, person_id, rolle_id)
VALUES 
    (1, 
    (SELECT person_id FROM person WHERE email = 'jean-alexander.mueller@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher')),
    (2,
    (SELECT person_id FROM person WHERE email = 'thomas.riechert@htwk-leipzig.de'),
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'));

INSERT INTO modul (
    kuerzel,                     -- 1: Kürzel des Moduls
    version,                     -- 2: Version des Moduls
    frueherer_schluessel,        -- 2.5: frühere Nummer des Moduls
    modultitel,                  -- 3: Titel des Moduls
    modultitel_englisch,          -- 4: Englischer Titel des Moduls
    kommentar,                   -- 5: Kommentar zum Modul
    niveau,                      -- 6: Studienniveau
    dauer,                       -- 7: Dauer in Semestern
    turnus,                      -- 8: Turnus (z.B. Sommer, Winter)
    studium_integrale,           -- 9: Ist das Modul Studium Integrale?
    sprachenzentrum,             -- 10: Ist das Modul im Sprachenzentrum?
    opal_link,                   -- 11: Link zu OPAL
    gruppengroesse_vorlesung,     -- 12: Gruppengröße Vorlesung
    gruppengroesse_uebung,        -- 13: Gruppengröße Übung
    gruppengroesse_praktikum,     -- 14: Gruppengröße Praktikum
    lehrform,                    -- 15: Lehrform
    medienform,                  -- 16: Medienform
    lehrinhalte,                 -- 17: Lehrinhalte
    qualifikationsziele,         -- 18: Qualifikationsziele
    sozial_und_selbstkompetenzen, -- 19: Soziale und Selbstkompetenzen
    besondere_zulassungsvoraussetzungen, -- 20: Besondere Zulassungsvoraussetzungen
    empfohlene_voraussetzungen,  -- 21: Empfohlene Voraussetzungen
    fortsetzungsmoeglichkeiten,  -- 22: Fortsetzungsmöglichkeiten
    hinweise,                    -- 23: Hinweise
    ects_credits,                -- 24: ECTS Credits
    praesenzeit_woche_vorlesung, -- 25: Präsenzzeit pro Woche in Vorlesung
    praesenzeit_woche_uebung,    -- 26: Präsenzzeit pro Woche in Übung
    praesenzeit_woche_praktikum, -- 27: Präsenzzeit pro Woche im Praktikum
    praesenzeit_woche_sonstiges, -- 28: Präsenzzeit pro Woche in Sonstigem
    selbststudienzeit_aufschluesselung, -- 29: Aufschlüsselung der Selbststudienzeit
    aktuelle_lehrressourcen,     -- 30: Aktuelle Lehrressourcen
    literatur,                   -- 31: Literatur
    fakultaet_id,                -- 32: Fakultät ID
    studienrichtung_id,          -- 33: Studienrichtung ID
    vertiefung_id,               -- 34: Vertiefung ID
    parent_modul_kuerzel,
    parent_modul_version  
)
VALUES 
(   'C114',                        -- 1: kuerzel NN
    2,                             -- 2: version  NN
    NULL,                          -- 2.5 frühere Nummer
    'Modellierung',                -- 3: modultitel NN
    'Modelling',                   -- 4: modultitel_englisch 
    'Kommentar Modellierung',      -- 5: kommentar
    'Bachelor',                    -- 6: niveau NN
    1,                             -- 7: dauer
    'Winter',                      -- 8: turnus
    FALSE,                         -- 9: studium_integrale
    FALSE,                         -- 10: sprachenzentrum
    NULL,                          -- 11: opal_link (kein Wert)
    120,                           -- 12: gruppengroesse_vorlesung
    30,                            -- 13: gruppengroesse_uebung
    NULL,                          -- 14: gruppengroesse_praktikum (kein Wert)
    'Vorlesung Übung Bearbeiten von Problemen und Lösungsfindung, Selbstudium anhand theoretischer und praktischer Übungsaufgaben', -- 15: lehrform
    NULL,                          -- 16: medienform (kein Wert)
    'Modellierung und formale Darstellung von Daten durch Mengen; Mengenoperationen
    Zusammenhängen durch Relationen, Funktionen, Äquivalenz- Ordnungsrelationen, Graphen
    strukturierten Daten durch Wörter, Texte, Sprachen, Bäume, Signaturen, Terme, strukturelle Induktion, algebraische Strukturen
    Eigenschaften und Anforderungen in Logiken (jeweils Syntax, Semantik, Folgern, Schließen)
    Software-Schnittstellen durch abstrakte Datentypen
    Abläufen und Berechnungen durch Zustandsübergangssysteme jeweils mit praktischen Modellierungsbeispielen;',  -- 17: lehrinhalte (gekürzt)
    'Die Studierenden können mathematische und logische Grundkonzepte zur Modellierung praktischer Aufgabenstellungen anwenden.
    Sie können Anforderungen an Software und Systeme formal beschreiben und wissen, dass deren Korrektheit mit formalen Methoden nachweisbar ist.', -- 18: qualifikationsziele (gekürzt)
    NULL,                          -- 19: sozial_und_selbstkompetenzen (kein Wert)
    NULL,                          -- 20: besondere_zulassungsvoraussetzungen (kein Wert)
    'Fertigkeiten in der Programmierung (derzeit C-Programmierung)',  -- 21: empfohlene_voraussetzungen (kein Wert)
    NULL,                          -- 22: fortsetzungsmoeglichkeiten (kein Wert)
    'regelmäßiges erfolgreiches Lösen der praktischen Übungsaufgaben (PVB) und 3 Kurzvorträge zu schriftlichen Übungsaufgaben (PVP)',  -- 23: hinweise
    8.0,                           -- 24: ects_credits
    4,                             -- 25: praesenzeit_woche_vorlesung
    2,                             -- 26: praesenzeit_woche_uebung
    0,                          -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                          -- 28: praesenzeit_woche_sonstiges (kein Wert)
    '156 Stunden
    28 Stunden Vorbereitung Lehrveranstaltung
    28 Stunden E-Learning
    84 Stunden Bearbeitung Prüfungsvorleistung
    16 Stunden Vorbereitung Prüfung',  -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    'Lehrmaterial und aktuelle Informationen: https://informatik.htwk-leipzig.de/schwarz', -- 30: aktuelle_lehrressourcen (kein Wert)
    'U. Kastens, H. Kleine Büning: "Modellierung: Grundlagen und formale Methoden", Hanser, 2008.
    M. Huth, M. Ryan: "Logic in Computer Science”, Cambridge University Press, 2010.
    U. Schöning: "Theoretische Informatik - kurzgefasst", Spektrum, in der aktuellen Auflage.
    M. Broy, R. Steinbrüggen: "Modellbildung in der Informatik", Springer, 2004', -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                         -- 33: studienrichtung_id (kein Wert)
    NULL,                         -- 34: vertiefung_id (kein Wert)
    NULL,                          -- 35: parent_modul_kuerzel,
    NULL                            -- parent_modul_version
),
(   'C287',                        -- 1: kuerzel NN
    3,                             -- 2: version  NN
    NULL,                          -- 2.5 frühere Nummer
    'Betriebssysteme und Rechnernetze', -- 3: modultitel NN
    'Operating Systems and Computer Networks', -- 4: modultitel_englisch 
    'Kommentar Betriebssysteme und Rechnernetze',      -- 5: kommentar
    'Bachelor',                     -- 6: niveau NN
    1,                              -- 7: dauer
    'Winter',                       -- 8: turnus
    FALSE,                          -- 9: studium_integrale
    FALSE,                          -- 10: sprachenzentrum
    NULL,                           -- 11: opal_link (kein Wert)
    120,                            -- 12: gruppengroesse_vorlesung
    30,                             -- 13: gruppengroesse_uebung
    NULL,                           -- 14: gruppengroesse_praktikum (kein Wert)
    NULL,                           -- 15: lehrform
    NULL,                           -- 16: medienform (kein Wert)
    NULL,                           -- 17: lehrinhalte (gekürzt)
    NULL,                           -- 18: qualifikationsziele (gekürzt)
    NULL,                           -- 19: sozial_und_selbstkompetenzen (kein Wert)
    NULL,                           -- 20: besondere_zulassungsvoraussetzungen (kein Wert)
    NULL,                           -- 21: empfohlene_voraussetzungen (kein Wert)
    NULL,                           -- 22: fortsetzungsmoeglichkeiten (kein Wert)
    NULL,                           -- 23: hinweise
    6.0,                            -- 24: ects_credits
    4,                              -- 25: praesenzeit_woche_vorlesung
    2,                              -- 26: praesenzeit_woche_uebung
    0,                           -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                           -- 28: praesenzeit_woche_sonstiges (kein Wert)
    NULL,                           -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    NULL,                           -- 30: aktuelle_lehrressourcen (kein Wert)
    NULL,                           -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                         -- 33: studienrichtung_id (kein Wert)
    NULL,                         -- 34: vertiefung_id (kein Wert)
    NULL,                          -- 35: parent_modul_kuerzel,
    NULL                            -- parent_modul_version
),
(   'C287.1',                        -- 1: kuerzel NN
    3,                             -- 2: version  NN
    NULL,                          -- 2.5 frühere Nummer
    'Betriebssysteme', -- 3: modultitel NN
    'Operating Systems', -- 4: modultitel_englisch 
    'Kommentar Betriebssysteme',      -- 5: kommentar
    'Bachelor',                     -- 6: niveau NN
    1,                              -- 7: dauer
    'Winter',                       -- 8: turnus
    FALSE,                          -- 9: studium_integrale
    FALSE,                          -- 10: sprachenzentrum
    NULL,                           -- 11: opal_link (kein Wert)
    120,                            -- 12: gruppengroesse_vorlesung
    30,                             -- 13: gruppengroesse_uebung
    NULL,                           -- 14: gruppengroesse_praktikum (kein Wert)
    'Seminaristische Vorlesung
    Übungen zu Theorie und praktischen Fertigkeiten im Computerpool
    Übungsaufgaben für das Selbststudium',                           -- 15: lehrform
    'Vorlesungen kombinieren vorbereitete Präsentationen und Erarbeitung von Themen an der Tafel
    Übungsaufgaben maßgeblich aus Standardwerken des Lehrgebiets',                           -- 16: medienform (kein Wert)
    'Aufgabenstellung und Begriffsbestimmung
    Entwicklung von Betriebssystemen
    Klassifikation und Methodik
    Prozesse: Konzept, Beschreibung, Kontrolle von Prozessen
        Speicherverwaltung
        Interprozesskommunikation: Signale, Pipes, Sockets, System V IPC (Message Queues, Semaphore, Shared Memory)
        Prozesskoordination: Concurrency, kritische Bereiche, Lösungsansätze
        Scheduling: Typen, Bursts, Prozess-Scheduling, Schedulingalgorithmen
        Virtualierungskonzepte
        Dateisysteme
        Sicherheitmechanismen
    PC-Betriebssysteme als Beispiel
        Prozesse, Dateisysteme, Nutzer
        Kommandoprozeduren unter Linux
        parallele Prozesse unter Linux
        einfache Formen der Kommunikation paralleler Prozesse
        praktische Übungen zur Programmierung von Kommandoprozeduren und parallelen Prozessen',  -- 17: lehrinhalte (gekürzt)
    'Die Studierenden können Grundkonzepte von modernen Betriebssystemen formal und sprachlich korrekt beschreiben und sind in der Lage, sie auf PC-Plattformen anzuwenden und nutzbar zu machen. Sie können selbständig und mit angemessenen Mitteln Betriebssysteme auf PC-Plattformen installieren und anpassen. Sowohl die Erstellung von Unix-spezifischen Anwendungsprogrammen unter Einsatz der Unix-API wie auch die Programmierung von Kommandoprozeduren kann selbständig unter Nutzung der vorhandenen Systemdokumentationen durchgeführt werden.',                           -- 18: qualifikationsziele (gekürzt)
    NULL,                           -- 19: sozial_und_selbstkompetenzen (kein Wert)
    NULL,                           -- 20: besondere_zulassungsvoraussetzungen (kein Wert)
    'Fertigkeiten in der Programmierung (derzeit C-Programmierung)', -- 21: empfohlene_voraussetzungen (kein Wert)
    NULL,                           -- 22: fortsetzungsmoeglichkeiten (kein Wert)
    NULL,                           -- 23: hinweise
    3.0,                            -- 24: ects_credits
    2,                              -- 25: praesenzeit_woche_vorlesung
    1,                              -- 26: praesenzeit_woche_uebung
    0,                           -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                           -- 28: praesenzeit_woche_sonstiges (kein Wert)
    NULL,                           -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    NULL,                           -- 30: aktuelle_lehrressourcen (kein Wert)
    'W. Stallings: Operating Systems. Prentice Hall, New Jersey, 2003
    Silberschatz: Operating System Concepts, 9nd. Wiley, 2012
    M. Hailperin: "Operating Systems an Middleware, Supporting Controlled Intercation", CC BY-SA 3.0, Rev 1.3
    J. Plötner, S. Wendzel: "Linux - Das umfassende Handbuch", Rheinwerk Computing, 2012', -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                         -- 33: studienrichtung_id (kein Wert)
    NULL,                         -- 34: vertiefung_id (kein Wert)
    'C287',                          -- 35: parent_modul_kuerzel,
    3                            -- parent_modul_version
),
(   'C287.2',                        -- 1: kuerzel NN
    3,                             -- 2: version  NN
    NULL,                          -- 2.5 frühere Nummer
    'Rechnernetze', -- 3: modultitel NN
    'Computer Networks', -- 4: modultitel_englisch 
    'Kommentar Rechnernetze',      -- 5: kommentar
    'Bachelor',                     -- 6: niveau NN
    1,                              -- 7: dauer
    'Winter',                       -- 8: turnus
    FALSE,                          -- 9: studium_integrale
    FALSE,                          -- 10: sprachenzentrum
    NULL,                           -- 11: opal_link (kein Wert)
    120,                            -- 12: gruppengroesse_vorlesung
    30,                             -- 13: gruppengroesse_uebung
    NULL,                           -- 14: gruppengroesse_praktikum (kein Wert)
    'Seminaristische Vorlesung
    Übungen zu Theorie und praktischen Fertigkeiten im Computerpool
    Übungsaufgaben für das Selbststudium',  -- 15: lehrform
    'Vorlesungen kombinieren vorbereitete Präsentationen und Erarbeitung von Themen an der Tafel
    Übungsaufgaben maßgeblich aus Standardwerken des Lehrgebiets',   -- 16: medienform (kein Wert)
    'Einführung in Netzwerktechnologien und Strukturen
        Datacenter / Vernetzung in Rechenzentren
        Lokale Netze bis zum Intranet
        Das Internet und andere Weitverkehrsnetze
        Überblick zu Mobil- und Zugangsnetzen
    Architektur und Grundprinzipien
        Paketvermittlung, Referenzmodelle und Betriebsverfahren
        Scheduling und Planung
        Direktverbindungsnetze
        Vermittlungsprinzipien, Routingverfahren
        Tunnel, Overlay
        Sicherheitsaspekte
    Technologien
        Internet Protocol (v4, v6, vX)
        IEEE 802-Technologien
        Virtualisierung, SDN, OpenFlow
        Carrier Ethernet, GMPLS',  -- 17: lehrinhalte (gekürzt)
    'Es besteht detailliertes, anwendungsfähiges Fachwissen auf dem Gebiet der Netzwerktechnologien, Strukturen und deren Grundprinzipien. Aufsetzend auf dem Verständnis der Grundprinzipien sowie der erworbenen praktischen Fähigkeiten sind sie in der Lage veränderte Methoden und Trends zu erkennen und deren Potential gegenüber etablierten Technologien zu ermitteln.',                           -- 18: qualifikationsziele (gekürzt)
    NULL,                           -- 19: sozial_und_selbstkompetenzen (kein Wert)
    NULL,                           -- 20: besondere_zulassungsvoraussetzungen (kein Wert)
    NULL,                           -- 21: empfohlene_voraussetzungen (kein Wert)
    NULL,                           -- 22: fortsetzungsmoeglichkeiten (kein Wert)
    NULL,                           -- 23: hinweise
    3.0,                            -- 24: ects_credits
    2,                              -- 25: praesenzeit_woche_vorlesung
    1,                              -- 26: praesenzeit_woche_uebung
    0,                           -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                           -- 28: praesenzeit_woche_sonstiges (kein Wert)
    NULL,                           -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    NULL,                           -- 30: aktuelle_lehrressourcen (kein Wert)
    'P. L. Dordal: "An Introduction to Computer Networks", CC BY-NC-ND 3.0, 2019.
    A. S. Tanenbaum, D. J. Wetherall: „Computer Netzworks“, Prentice Hall, 5. Auflage, 2010.
    K. R. Fall, W. R. Stevens: "TCP/IP Illustrated volume 1: The Protocols", Addison-Wesley, 2011.
    L. L. Peterson, B. S. Davie: "Computer Networks: A Systems Approach", Morgan Kaufmann, 5. Auflage, 2011.
    T. Nadeu, K. Gray: "SDN: Software Defined Networks", O`Reilly, 2013.
    „Ethernet“, Heise Verlag, 2008.', -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                         -- 33: studienrichtung_id (kein Wert)
    NULL,                         -- 34: vertiefung_id (kein Wert)
    'C287',                          -- 35: parent_modul_kuerzel,
    3                            -- parent_modul_version
);

-- Modulverantwortliche & Dozenten
-- Modellierung
INSERT INTO modul_person_rolle (modul_kuerzel, modul_version, person_id, rolle_id)
VALUES 
    ('C114', 2, 
    (SELECT person_id FROM person WHERE email = 'sibylle.schwarz@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher')),
    ('C114', 2, 
    (SELECT person_id FROM person WHERE email = 'sibylle.schwarz@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Dozent'));

-- Betriebssysteme und Rechnernetze 
INSERT INTO modul_person_rolle (modul_kuerzel, modul_version, person_id, rolle_id)
VALUES 
    ('C287', 3, 
    (SELECT person_id FROM person WHERE email = 'jean-alexander.mueller@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher')),
    ('C287.1', 3,  -- Betriebssysteme
    (SELECT person_id FROM person WHERE email = 'thomas.kudrass@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher')),
    ('C287.1', 3,  
    (SELECT person_id FROM person WHERE email = 'thomas.kudrass@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Dozent')),
    ('C287.2', 3,  -- Rechnernetze
    (SELECT person_id FROM person WHERE email = 'jean-alexander.mueller@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher')),
    ('C287.2', 3,  
    (SELECT person_id FROM person WHERE email = 'jean-alexander.mueller@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Dozent'));

-- Zuordnung der Module zu Studiengängen
INSERT INTO modul_studiengang (modul_kuerzel, modul_version, studiengang_id, semester ,modul_typ)
VALUES
    ('C114', 2, 1, 1, 'Pflichtmodul'),  -- Modellierung
    ('C114', 2, 2, 1, 'Pflichtmodul'),
    ('C287', 3, 1, 3, 'Pflichtmodul'),  -- Betriebssysteme und Rechnernetze
    ('C287', 3, 2, 3, 'Pflichtmodul'),
    ('C287.1', 3, 1, 3, 'Pflichtmodul'), -- Betriebssysteme
    ('C287.1', 3, 2, 3, 'Pflichtmodul'),
    ('C287.2', 3, 1, 3, 'Pflichtmodul'), -- Rechnernetze
    ('C287.2', 3, 2, 3, 'Pflichtmodul');