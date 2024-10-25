
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
INSERT INTO person (vorname, nachname, titel, email, telefonnummer, raum, funktion, password)
VALUES 
    ('Jean-Alexander', 'Müller', 'Prof. Dr.-Ing.', 'jean-alexander.mueller@htwk-leipzig.de', '+49 341 3076-6638', 'ZU 123', 'Studiendekan', '$2a$14$coFdmFbPMNZKnYx/jG2/wOwKoLiGh7.obc8SV29q5iZ28bKngQDqK'),
    ('Sibylle', 'Schwarz', 'Prof. Dr. rer. nat.', 'sibylle.schwarz@htwk-leipzig.de', '+49 341 3076-6483', 'ZU 411', 'Professorin', '$2a$14$40LasxGj14cXsbizo/haMOEBV141/VrN1c1zdWt2/R2chY9i510Wi'),
    ('Jens', 'Wagner', 'Prof. Dr. rer. nat.', 'jens.wagner@htwk-leipzig.de', '+49 341 3076-6494', 'LI 015', 'Professor', '$2a$14$WtjFBU6FS.ulh/A1F6SS1e3/xJlRd8lumXBEZXgxHqeMO4OoE1To6'),
    ('Hanna', 'Brodowsky', NULL, 'hanna.brodowsky@htwk-leipzig.de', NULL, NULL, 'Dozentin', '$2a$14$awxB1ZjRnqKjXya/fkM4s.iTlkF6rX0Tjx5KI0SjMeXTnTJCie5bW'),
    ('Mario', 'Hlawitschka', 'Prof. Dr. rer. nat.', 'mario.hlawitschka@htwk-leipzig.de', '+49 341 3076-6493', 'ZU 224', 'Professor', '$2a$14$FvJ4osEf6iW42LNMidUDx.fDgzBjr/Kpu7MlR8dsXce88bFoOjlC6'),
    ('Martin', 'Grüttmüller', 'Prof. Dr. rer. nat. habil.', 'martin.gruettmueller@htwk-leipzig.de', '+49 341 3076-6487', 'ZU 412', 'Professor', '$2a$14$ko6zDJfW439N6Q3coxlR8.30fKKiGCm5S9seE4NARxhVEnGIvnAXS'),
    ('Karsten', 'Weicker', 'Prof. Dr. rer. nat.', 'karsten-weicker@htwk-leipzig.de', '+49 341 3076-6395', 'ZU 410', 'Professor', '$2a$14$jHkvqQUyM4pUM2gHSIh.Vuy9tAT3fynpNGzL5LLwphbPrzj6KI8gW'),
    ('Thomas', 'Kudraß', 'Prof. Dr.-Ing.', 'thomas.kudrass@htwk-leipzig.de', '+49 341 3076-6420', 'ZU 130', 'Professor', '$2a$14$D3vb1GbGut/UdrOPm/iYg.dKkFXKQQny6yxe9yRfV4bHVicvGjJSi'),
    ('Thomas', 'Riechert', 'Prof. Dr. rer. nat.', 'thomas.riechert@htwk-leipzig.de', '+49 341 3076-6413', 'ZU 507', 'Professor', '$2a$14$/OAmiBW5bGkdW.JkhQzYoOwHo22dR5HdWBhdxM7q/0rWqR3zaG7iK'),
    ('Antje', 'Tober-Nietner', 'Dr.', 'antje.tober@htwk-leipzig.de', NULL, NULL, 'Dozentin', '$2a$14$chIIQ1XgIoOT1kt7i5.5qOs6V/FOedfxTAJ6fp5qkv1V8tGqGWXAW'),
    ('Johannes', 'Waldmann', 'Prof. Dr. rer. nat.', 'johannes.waldmann@htwk-leipzig.de', '+49 341 3076-6479', 'ZU 129', 'Professor', '$2a$14$zEbhgM7ArQfOXv73gJVEkOuWq8M8iWqKYWBapNFnNhJ9dR9EAa.YW'),
    ('Andreas', 'Both', 'Prof. Dr. rer. nat.', 'andreas.both@htwk-leipzig.de', '+49 341 3076-6256', 'ZU 529', 'Professor', '$2a$14$d4EA1.cfmpf2D0.a5NnYfeHIFTi22dvTTh/fGYWdAfxDUP74uM0rq');

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
    ('Hochschulleitung'),
    ('Fakultätsverantwortlicher');

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
    kuerzel,                        -- 1: Kürzel des Moduls
    version,                        -- 2: Version des Moduls
    frueherer_schluessel,           -- 2.5: frühere Nummer des Moduls
    modultitel,                     -- 3: Titel des Moduls
    modultitel_englisch,            -- 4: Englischer Titel des Moduls
    kommentar,                      -- 5: Kommentar zum Modul
    niveau,                         -- 6: Studienniveau
    dauer,                          -- 7: Dauer in Semestern
    turnus,                         -- 8: Turnus (z.B. Sommer, Winter)
    studium_integrale,              -- 9: Ist das Modul Studium Integrale?
    sprachenzentrum,                -- 10: Ist das Modul im Sprachenzentrum?
    opal_link,                      -- 11: Link zu OPAL
    gruppengroesse_vorlesung,       -- 12: Gruppengröße Vorlesung
    gruppengroesse_uebung,          -- 13: Gruppengröße Übung
    gruppengroesse_praktikum,       -- 14: Gruppengröße Praktikum
    lehrform,                       -- 15: Lehrform
    medienform,                     -- 16: Medienform
    lehrinhalte,                    -- 17: Lehrinhalte
    qualifikationsziele,            -- 18: Qualifikationsziele
    sozial_und_selbstkompetenzen,   -- 19: Soziale und Selbstkompetenzen
    besondere_zulassungsvoraussetzungen, -- 20: Besondere Zulassungsvoraussetzungen
    empfohlene_voraussetzungen,     -- 21: Empfohlene Voraussetzungen
    fortsetzungsmoeglichkeiten,     -- 22: Fortsetzungsmöglichkeiten
    hinweise,                       -- 23: Hinweise
    ects_credits,                   -- 24: ECTS Credits
    praesenzeit_woche_vorlesung,    -- 25: Präsenzzeit pro Woche in Vorlesung
    praesenzeit_woche_uebung,       -- 26: Präsenzzeit pro Woche in Übung
    praesenzeit_woche_praktikum,    -- 27: Präsenzzeit pro Woche im Praktikum
    praesenzeit_woche_sonstiges,    -- 28: Präsenzzeit pro Woche in Sonstigem
    selbststudienzeit_aufschluesselung, -- 29: Aufschlüsselung der Selbststudienzeit
    aktuelle_lehrressourcen,        -- 30: Aktuelle Lehrressourcen
    literatur,                      -- 31: Literatur
    fakultaet_id,                   -- 32: Fakultät ID
    studienrichtung_id,             -- 33: Studienrichtung ID
    vertiefung_id,                  -- 34: Vertiefung ID
    parent_modul_kuerzel,           -- 35: parent_modul_kuerzel
    parent_modul_version,            -- 36: parent_modul_version
    vorheriger_zustand_id       -- 37: vorheriger_zustand_id
)
VALUES 
(   'C114',                         -- 1: kuerzel NN
    2,                              -- 2: version  NN
    NULL,                           -- 2.5 frühere Nummer
    'Modellierung',                 -- 3: modultitel NN
    'Modelling',                    -- 4: modultitel_englisch 
    'Kommentar Modellierung',       -- 5: kommentar
    'Bachelor',                     -- 6: niveau NN
    1,                              -- 7: dauer
    'Winter',                       -- 8: turnus
    FALSE,                          -- 9: studium_integrale
    FALSE,                          -- 10: sprachenzentrum
    NULL,                           -- 11: opal_link (kein Wert)
    120,                            -- 12: gruppengroesse_vorlesung
    30,                             -- 13: gruppengroesse_uebung
    NULL,                           -- 14: gruppengroesse_praktikum (kein Wert)
    'Vorlesung Übung Bearbeiten von Problemen und Lösungsfindung, Selbstudium anhand theoretischer und praktischer Übungsaufgaben', -- 15: lehrform
    NULL,                           -- 16: medienform (kein Wert)
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
    'Die Studierenden können mathematische und logische Grundkonzepte zur Modellierung praktischer Aufgabenstellungen anwenden.
    Sie können Anforderungen an Software und Systeme formal beschreiben und wissen, dass deren Korrektheit mit formalen Methoden nachweisbar ist.', -- 18: qualifikationsziele (gekürzt)
    NULL,                           -- 19: sozial_und_selbstkompetenzen (kein Wert)
    NULL,                           -- 20: besondere_zulassungsvoraussetzungen (kein Wert)
    'Fertigkeiten in der Programmierung (derzeit C-Programmierung)',  -- 21: empfohlene_voraussetzungen (kein Wert)
    NULL,                           -- 22: fortsetzungsmoeglichkeiten (kein Wert)
    'regelmäßiges erfolgreiches Lösen der praktischen Übungsaufgaben (PVB) und 3 Kurzvorträge zu schriftlichen Übungsaufgaben (PVP)',  -- 23: hinweise
    8.0,                            -- 24: ects_credits
    4,                              -- 25: praesenzeit_woche_vorlesung
    2,                              -- 26: praesenzeit_woche_uebung
    0,                              -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                              -- 28: praesenzeit_woche_sonstiges (kein Wert)
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
    NULL,                           -- 33: studienrichtung_id (kein Wert)
    NULL,                           -- 34: vertiefung_id (kein Wert)
    NULL,                           -- 35: parent_modul_kuerzel,
    NULL,                            -- parent_modul_version
    NULL                            -- vorheriger_zustand_id
),
(   'C287',                         -- 1: kuerzel NN
    3,                              -- 2: version  NN
    NULL,                           -- 2.5 frühere Nummer
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
    0,                              -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                              -- 28: praesenzeit_woche_sonstiges (kein Wert)
    NULL,                           -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    NULL,                           -- 30: aktuelle_lehrressourcen (kein Wert)
    NULL,                           -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                           -- 33: studienrichtung_id (kein Wert)
    NULL,                           -- 34: vertiefung_id (kein Wert)
    NULL,                           -- 35: parent_modul_kuerzel,
    NULL,                           -- parent_modul_version
    NULL                            -- vorheriger_zustand_id
),
(   'C287.1',                       -- 1: kuerzel NN
    3,                              -- 2: version  NN
    NULL,                           -- 2.5 frühere Nummer
    'Betriebssysteme',              -- 3: modultitel NN
    'Operating Systems',            -- 4: modultitel_englisch 
    'Kommentar Betriebssysteme',    -- 5: kommentar
    'Bachelor',                     -- 6: niveau NN
    1,                              -- 7: dauer
    'Winter',                       -- 8: turnus
    FALSE,                          -- 9: studium_integrale
    FALSE,                          -- 10: sprachenzentrum
    'https://bildungsportal.sachsen.de/opal/auth/RepositoryEntry/20425539587/CourseNode/99489065090706',    -- 11: opal_link (kein Wert)
    120,                            -- 12: gruppengroesse_vorlesung
    30,                             -- 13: gruppengroesse_uebung
    NULL,                           -- 14: gruppengroesse_praktikum (kein Wert)
    'Seminaristische Vorlesung Übungen zu Theorie und praktischen Fertigkeiten im Computerpool Übungsaufgaben für das Selbststudium',    -- 15: lehrform
    'Vorlesungen kombinieren vorbereitete Präsentationen und Erarbeitung von Themen an der Tafel Übungsaufgaben maßgeblich aus Standardwerken des Lehrgebiets',                           -- 16: medienform (kein Wert)
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
    0,                              -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                              -- 28: praesenzeit_woche_sonstiges (kein Wert)
    NULL,                           -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    NULL,                           -- 30: aktuelle_lehrressourcen (kein Wert)
    'W. Stallings: Operating Systems. Prentice Hall, New Jersey, 2003
    Silberschatz: Operating System Concepts, 9nd. Wiley, 2012
    M. Hailperin: "Operating Systems an Middleware, Supporting Controlled Intercation", CC BY-SA 3.0, Rev 1.3
    J. Plötner, S. Wendzel: "Linux - Das umfassende Handbuch", Rheinwerk Computing, 2012', -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                           -- 33: studienrichtung_id (kein Wert)
    NULL,                           -- 34: vertiefung_id (kein Wert)
    'C287',                         -- 35: parent_modul_kuerzel,
    3,                               -- parent_modul_version
    NULL                            -- vorheriger_zustand_id
),
(   'C287.2',                       -- 1: kuerzel NN
    3,                              -- 2: version  NN
    NULL,                           -- 2.5 frühere Nummer
    'Rechnernetze',                 -- 3: modultitel NN
    'Computer Networks',            -- 4: modultitel_englisch 
    'Kommentar Rechnernetze',       -- 5: kommentar
    'Bachelor',                     -- 6: niveau NN
    1,                              -- 7: dauer
    'Winter',                       -- 8: turnus
    FALSE,                          -- 9: studium_integrale
    FALSE,                          -- 10: sprachenzentrum
    'https://bildungsportal.sachsen.de/opal/auth/RepositoryEntry/20425539587/CourseNode/99489065090706',                           -- 11: opal_link (kein Wert)
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
        Carrier Ethernet, GMPLS',   -- 17: lehrinhalte (gekürzt)
    'Es besteht detailliertes, anwendungsfähiges Fachwissen auf dem Gebiet der Netzwerktechnologien, Strukturen und deren Grundprinzipien. Aufsetzend auf dem Verständnis der Grundprinzipien sowie der erworbenen praktischen Fähigkeiten sind sie in der Lage veränderte Methoden und Trends zu erkennen und deren Potential gegenüber etablierten Technologien zu ermitteln.',                           -- 18: qualifikationsziele (gekürzt)
    NULL,                           -- 19: sozial_und_selbstkompetenzen (kein Wert)
    NULL,                           -- 20: besondere_zulassungsvoraussetzungen (kein Wert)
    NULL,                           -- 21: empfohlene_voraussetzungen (kein Wert)
    NULL,                           -- 22: fortsetzungsmoeglichkeiten (kein Wert)
    NULL,                           -- 23: hinweise
    3.0,                            -- 24: ects_credits
    2,                              -- 25: praesenzeit_woche_vorlesung
    1,                              -- 26: praesenzeit_woche_uebung
    0,                              -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                              -- 28: praesenzeit_woche_sonstiges (kein Wert)
    NULL,                           -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    NULL,                           -- 30: aktuelle_lehrressourcen (kein Wert)
    'P. L. Dordal: "An Introduction to Computer Networks", CC BY-NC-ND 3.0, 2019.
    A. S. Tanenbaum, D. J. Wetherall: „Computer Netzworks“, Prentice Hall, 5. Auflage, 2010.
    K. R. Fall, W. R. Stevens: "TCP/IP Illustrated volume 1: The Protocols", Addison-Wesley, 2011.
    L. L. Peterson, B. S. Davie: "Computer Networks: A Systems Approach", Morgan Kaufmann, 5. Auflage, 2011.
    T. Nadeu, K. Gray: "SDN: Software Defined Networks", O`Reilly, 2013.
    „Ethernet“, Heise Verlag, 2008.', -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                           -- 33: studienrichtung_id (kein Wert)
    NULL,                           -- 34: vertiefung_id (kein Wert)
    'C287',                         -- 35: parent_modul_kuerzel,
    3,                               -- parent_modul_version
    NULL                            -- vorheriger_zustand_id
), -- ------------------------------
(   'C073',                         -- 1: kuerzel NN
    0,                              -- 2: version  NN
    NULL,                           -- 2.5 frühere Nummer
    'Softwareprojekt I',            -- 3: modultitel NN
    'Software Engineering Project I',                   -- 4: modultitel_englisch 
    'Kommentar Softwareprojekt I',      -- 5: kommentar
    'Bachelor',                     -- 6: niveau NN
    1,                              -- 7: dauer
    'Sommer',                       -- 8: turnus
    FALSE,                          -- 9: studium_integrale
    FALSE,                          -- 10: sprachenzentrum
    'https://bildungsportal.sachsen.de/opal/auth/RepositoryEntry/39268417537/CourseNode/77094584712802', -- 11: opal_link (kein Wert)
    120,                            -- 12: gruppengroesse_vorlesung
    30,                             -- 13: gruppengroesse_uebung
    NULL,                           -- 14: gruppengroesse_praktikum (kein Wert)
    'Begleitende Inputs werden als Vorlesung präsentiert
    Die Projekte werden in Teams selbstorganisiert bearbeitet
    ausgewählte Meetings und Zwischenstandspräsentationen werden duch die Betreuer organisiert und abgenommen', -- 15: lehrform
    NULL,                           -- 16: medienform (kein Wert)
    'Vorstellung der Anforderungen
    Teambildung
    Erstellung einer Anforderungsspezifikation und einer Architekturvision mit Präsentationen an Meilensteinen
    Erstellung einer produktiv einsetzbaren ersten Version der Software mit Präsentationen an Meilensteinen: erste Funktionalitäten sollten enthalten sein und prototypisch eine Vision für die Nutzungsoberfläche der gesamten Software vorhanden sein',  -- 17: lehrinhalte (gekürzt)
    'Die Studierenden können mathematische und logische Grundkonzepte zur Modellierung praktischer Aufgabenstellungen anwenden.
    Sie können Anforderungen an Software und Systeme formal beschreiben und wissen, dass deren Korrektheit mit formalen Methoden nachweisbar ist.', -- 18: qualifikationsziele (gekürzt)
    NULL,                           -- 19: sozial_und_selbstkompetenzen (kein Wert)
    NULL,                           -- 20: besondere_zulassungsvoraussetzungen (kein Wert)
    'Kompetenzen der Softwaretechnik und der Programmierung sollten soweit vorhanden sein, dass kleine Programme mit graphischer Benutzeroberfläche erstellt werden können)',  -- 21: empfohlene_voraussetzungen (kein Wert)
    NULL,                           -- 22: fortsetzungsmoeglichkeiten (kein Wert)
    NULL,                           -- 23: hinweise
    5.0,                            -- 24: ects_credits
    1,                              -- 25: praesenzeit_woche_vorlesung
    0,                              -- 26: praesenzeit_woche_uebung
    1,                              -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                              -- 28: praesenzeit_woche_sonstiges (kein Wert)
    NULL,                           -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    NULL,                           -- 30: aktuelle_lehrressourcen (kein Wert)
    'J. Ludewig, H. Lichter: "Software Engineering", dpunkt, in der aktuellen Auflage.
    C. Rupp et al.: "UML 2 glasklar. Praxiswissen für die UML-Modellierung", Hanser, in der aktuellen Auflage.
    H. Kellner: "Soziale Kompetenz für Ingenieure, Informatiker und Naturwissenschaftler", Hanser,2006.
    U. Vigenschow, B. Schneider: "Soft Skills für Softwareentwickler", dpunkt, in der aktuellen Auflage.
    R. Pichler: "Scrum - Agiles Projektmanagement erfolgreich einsetzen", dpunkt, 2007.', -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                           -- 33: studienrichtung_id (kein Wert)
    NULL,                           -- 34: vertiefung_id (kein Wert)
    NULL,                           -- 35: parent_modul_kuerzel,
    NULL,                            -- parent_modul_version
    NULL                            -- vorheriger_zustand_id
),
(   'C171',                         -- 1: kuerzel NN
    0,                              -- 2: version  NN
    NULL,                           -- 2.5 frühere Nummer
    'Softwareprojekt II',           -- 3: modultitel NN
    'Software Engineering Project II', -- 4: modultitel_englisch 
    'Kommentar Softwareprojekt II',      -- 5: kommentar
    'Bachelor',                     -- 6: niveau NN
    1,                              -- 7: dauer
    'Winter',                       -- 8: turnus
    FALSE,                          -- 9: studium_integrale
    FALSE,                          -- 10: sprachenzentrum
    'https://bildungsportal.sachsen.de/opal/auth/RepositoryEntry/41558605824?24', -- 11: opal_link (kein Wert)
    120,                            -- 12: gruppengroesse_vorlesung
    30,                             -- 13: gruppengroesse_uebung
    NULL,                           -- 14: gruppengroesse_praktikum (kein Wert)
    'begleitende Vorlesung mit Impulsreferaten
    Abschlussveranstaltungen inklusive kleiner Produktmesse
    Softwareentwicklung findet selbstorganisiert statt
    bei Meetings und Zwischenstandspräsentationen hospitieren die Betreuer',   -- 15: lehrform
    NULL,                           -- 16: medienform (kein Wert)
    'Erstellung einer Anforderungsspezifikation und einer Architekturvision mit Präsentationen an mehreren Meilensteinen
    Erstellung einer produktiv einsetzbaren Software mit Präsentationen an mehreren Meilensteinen
    Poster-Abschlusspräsentation
    Abschlusspräsentation als Vortrag',                           -- 17: lehrinhalte (gekürzt)
    'Studierende sind in der Lage ein in einem agilen Vorgehensmodell ein bestehendes Softwareentwicklungsprojekt fortzuführen und erfolgreich zu beenden, dass dem Kunden ein (zumindest partiell) funktionsfähiges Produkt ausgeliefert werden kann.
    Sie können fremden Quelltext lesen, darin Entwurfskonzepte erkennen sowie Änderungen durchführen. Sie erkennen selbständig Schnittstellen zu den Arbeitspaketen anderer Teammitglieder, können die Probleme benennen und selbständig Absprachen durchführen.
    Insbesondere sind sie in der Lage Teilmodule zu entwerfen und im Rahmen der Gesamtsoftware umzusetzen. Innerhalb des Projektkontexts beherrschen sie erfolgreich Strategien zur Qualitätssicherung, d.h. Fehlermanagement, Uni-Tests und Reviews. Die Qualität von Artefakten kann im Rahmen von Reviews beurteilt werden. Darüber hinaus werden im Projektkontext Probleme hinsichtlich der Planung und Durchführbarkeit erkannt sowie Maßnahmen vorgeschlagen.',  -- 18: qualifikationsziele (gekürzt)
    NULL,                           -- 19: sozial_und_selbstkompetenzen (kein Wert)
    NULL,                           -- 20: besondere_zulassungsvoraussetzungen (kein Wert)
    'Der Besuch des Softwareprojekts I im vorherigen Semester ist dringend anzuraten, da die Teams und die Projekte fortgeführt werden. Andernfalls ist der Arbeitsaufwand für die Einarbeitung ungleich höher.
    Analog zu "Softwareprojekt I" werden auch hier hinreichend ausgeprägte Programmier- und Softwareentwicklungskompetenzen erwartet.',                           -- 21: empfohlene_voraussetzungen (kein Wert)
    NULL,                           -- 22: fortsetzungsmoeglichkeiten (kein Wert)
    NULL,                           -- 23: hinweise
    5.0,                            -- 24: ects_credits
    0.5,                            -- 25: praesenzeit_woche_vorlesung
    0,                              -- 26: praesenzeit_woche_uebung
    1,                              -- 27: praesenzeit_woche_praktikum (kein Wert)
    0,                              -- 28: praesenzeit_woche_sonstiges (kein Wert)
    NULL,                           -- 29: selbststudienzeit_aufschluesselung (gekürzt)
    NULL,                           -- 30: aktuelle_lehrressourcen (kein Wert)
    'J. Ludewig, H. Lichter: "Software Engineering", dpunkt, in der aktuellen Auflage.
    C. Rupp et al.: "UML 2 glasklar. Praxiswissen für die UML-Modellierung", Hanser, in der aktuellen Auflage.
    H. Kellner: "Soziale Kompetenz für Ingenieure, Informatiker und Naturwissenschaftler", Hanser, 2006.
    U. Vigenschow, B. Schneider: "Soft Skills für Softwareentwickler", dpunkt, in der aktuellen Auflage.
    R. Pichler: "Scrum - Agiles Projektmanagement erfolgreich einsetzen", dpunkt, 2007.',  -- 31: literatur (gekürzt)
    (SELECT fakultaet_id FROM fakultaet WHERE kuerzel = 'IM'), -- 32: fakultaet_id
    NULL,                           -- 33: studienrichtung_id (kein Wert)
    NULL,                           -- 34: vertiefung_id (kein Wert)
    NULL,                           -- 35: parent_modul_kuerzel,
    NULL,                            -- parent_modul_version
    NULL                            -- vorheriger_zustand_id
);

INSERT INTO modul_voraussetzung (modul_kuerzel, modul_version, vorausgesetztes_modul_kuerzel, vorausgesetztes_modul_version)
VALUES('C171', 0, 'C073', 0);

-- Modulverantwortliche & Dozenten
-- Modellierung
INSERT INTO modul_person_rolle (modul_kuerzel, modul_version, person_id, rolle_id)
VALUES 
    ('C114', 2, 
    (SELECT person_id FROM person WHERE email = 'sibylle.schwarz@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher')),
    ('C114', 2, 
    (SELECT person_id FROM person WHERE email = 'sibylle.schwarz@htwk-leipzig.de'), 
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Dozent')),
    -- Betriebssysteme und Rechnernetze 
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
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Dozent')),
    ('C073', 0,    -- Softwareprojekt I
    (SELECT person_id FROM person WHERE email = 'karsten-weicker@htwk-leipzig.de'),
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher')),
    ('C073', 0, 
    (SELECT person_id FROM person WHERE email = 'andreas.both@htwk-leipzig.de'),
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Dozent')),
    ('C171', 0,    -- Softwareprojekt II
    (SELECT person_id FROM person WHERE email = 'karsten-weicker@htwk-leipzig.de'),
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher')),
    ('C171', 0, 
    (SELECT person_id FROM person WHERE email = 'karsten-weicker@htwk-leipzig.de'),
    (SELECT rolle_id FROM rolle WHERE bezeichnung = 'Dozent'));

-- Zuordnung der Module zu Studiengängen
INSERT INTO modul_studiengang (modul_kuerzel, modul_version, studiengang_id, modul_typ, semester)
VALUES
    ('C114', 2, 1, 'Pflichtmodul', 1),  -- Modellierung
    ('C114', 2, 2, 'Pflichtmodul', 1),
    ('C287', 3, 1, 'Pflichtmodul', 3),  -- Betriebssysteme und Rechnernetze
    ('C287', 3, 2, 'Pflichtmodul', 3),
    ('C287.1', 3, 1, 'Pflichtmodul', 3), -- Betriebssysteme
    ('C287.1', 3, 2, 'Pflichtmodul', 3),
    ('C287.2', 3, 1, 'Pflichtmodul', 3), -- Rechnernetze
    ('C287.2', 3, 2, 'Pflichtmodul', 3),
    ('C073', 0, 1, 'Pflichtmodul', 4),  -- Softwareprojekt I
    ('C073', 0, 2, 'Pflichtmodul', 4),
    ('C171', 0, 1, 'Pflichtmodul', 5),  -- Softwareprojekt II
    ('C171', 0, 2, 'Pflichtmodul', 5);


INSERT INTO taxonomie_kategorie (name, stufe, beschreibung) 
VALUES 
    ('Erinnern', 1, 'Auf relevantes Wissen im Langzeitgedächtnis zugreifen'),
    ('Verstehen', 2, 'Informationen in der Lerneinheit Bedeutung zuordnen, seien sie mündlich, schriftlich oder grafisch'),
    ('Anwenden', 3, 'Einen Handlungsablauf (ein Schema, eine Methode) in einer bestimmten Situation ausführen oder verwenden'),
    ('Analysieren', 4, 'Lerninhalte in ihre konstruierten Elemente zerlegen und bestimmen, wie diese untereinander zu einer übergreifenden Struktur oder einem übergreifende Zweck verbunden sind'),
    ('Beurteilen', 5, 'Urteile abgeben aufgrund von Kriterien oder Standards'),
    ('Kreieren', 6, 'Elemente zu einem kohärenten oder funktionierenden Ganzen zusammen setzen; Elemente zu einem neuen Muster oder einer neuen Struktur zusammenfügen');

INSERT INTO kognitiver_prozess (kategorie_id, name) 
VALUES 
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'Erkennen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'Erinnern'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'Interpretieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'Veranschaulichen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'Klassifizieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'Zusammenfassen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'Folgern'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'Vergleichen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'Erklären'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'Ausführen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'Implementieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'Differenzieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'Organisieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'Zuordnen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'Überprüfen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'Bewerten'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'Generieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'Planen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'Entwickeln');

INSERT INTO taxonomie_verb (kategorie_id, verb) 
VALUES 
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'schreiben'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'definieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'reproduzieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'auflisten'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'schildern'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'bezeichnen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'aufsagen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'angeben'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'aufzählen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'benennen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'zeichnen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'ausführen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'skizzieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Erinnern'), 'erzählen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'darstellen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'beschreiben'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'bestimmen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'demonstrieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'ableiten'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'diskutieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'erklären'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'formulieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'zusammenfassen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'lokalisieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'präsentieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'erläutern'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'übertragen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Verstehen'), 'wiederholen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'durchführen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'berechnen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'benutzen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'herausfinden'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'löschen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'ausfüllen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'eintragen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'drucken'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'anwenden'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'lösen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'planen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'illustrieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'formatieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Anwenden'), 'bearbeiten'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'testen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'kontrastieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'vergleichen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'isolieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'auswählen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'unterscheiden'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'gegenüberstellen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'kritisieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'analysieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'bestimmen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'experimentieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'sortieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'untersuchen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Analysieren'), 'kategorisieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'beurteilen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'argumentieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'voraussagen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'wählen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'evaluieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'begründen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'prüfen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'entscheiden'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'kritisieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'benoten'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'schätzen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'werten'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'unterstützen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Beurteilen'), 'klassifizieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'zusammensetzen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'sammeln'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'organisieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'konstruieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'präparieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'schreiben'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'entwerfen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'schlussfolgern'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'verbinden'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'konzipieren'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'zuordnen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'zusammenstellen'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'ableiten'),
    ((SELECT id FROM taxonomie_kategorie WHERE name = 'Kreieren'), 'entwickeln');

-- Berechtigungen einfügen
INSERT INTO berechtigung (bezeichnung)
VALUES 
    ('modul_erstellen'),
    ('modul_lesen'),
    ('modul_bearbeiten'),
    ('modul_loeschen'),
    ('modul_bearbeiter_zuweisen'),
    ('modul_archivieren'),
    ('studiengang_erstellen'),
    ('studiengang_lesen'),
    ('studiengang_bearbeiten'),
    ('studiengang_loeschen'),
    ('studiengang_bearbeiter_zuweisen'),
    ('studiengang_archivieren'),
    ('modul_zur_begutachtung_weiterleiten'),
    ('modul_zur_ueberarbeitung_weiterleiten'),
    ('studiengang_zur_begutachtung_weiterleiten'),
    ('studiengang_zur_ueberarbeitung_weiterleiten'),
    ('studiengang_zur_freigabe_weiterleiten'),
    ('modul_abnehmen'),
    ('pdf_dokumente_pflegen');

INSERT INTO rolle_berechtigung (rolle_id, berechtigung_id)
VALUES 
    -- Modulbearbeiter
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulbearbeiter'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulbearbeiter'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_zur_begutachtung_weiterleiten')),

    -- Modulverantwortlicher
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_zur_begutachtung_weiterleiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiter_zuweisen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_zur_ueberarbeitung_weiterleiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Modulverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_archivieren')),

    -- Studiengangbearbeiter
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangbearbeiter'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangbearbeiter'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_zur_begutachtung_weiterleiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangbearbeiter'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangbearbeiter'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiter_zuweisen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangbearbeiter'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_erstellen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangbearbeiter'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_zur_begutachtung_weiterleiten')),

    -- Studiengangverantwortlicher
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_zur_begutachtung_weiterleiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiter_zuweisen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_erstellen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_zur_begutachtung_weiterleiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_bearbeiter_zuweisen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_zur_ueberarbeitung_weiterleiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_zur_freigabe_weiterleiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_abnehmen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Studiengangverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_zur_ueberarbeitung_weiterleiten')),

    -- Fakultätsverantwortliche/-r für Modulux
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Fakultätsverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_erstellen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Fakultätsverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Fakultätsverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_bearbeiter_zuweisen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Fakultätsverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'modul_archivieren')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Fakultätsverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_erstellen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Fakultätsverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Fakultätsverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_bearbeiter_zuweisen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Fakultätsverantwortlicher'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_archivieren')),

    -- Prozesskontrolle
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Prozesskontrolle'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_lesen')),

    -- Hochschulleitung
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Hochschulleitung'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_lesen')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Hochschulleitung'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_bearbeiten')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Hochschulleitung'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'studiengang_archivieren')),
    ((SELECT rolle_id FROM rolle WHERE bezeichnung = 'Hochschulleitung'), (SELECT berechtigung_id FROM berechtigung WHERE bezeichnung = 'pdf_dokumente_pflegen'));