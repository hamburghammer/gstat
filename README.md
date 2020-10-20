# gstat

[![Build Status](https://cloud.drone.io/api/badges/hamburghammer/gstat/status.svg?ref=refs/heads/master)](https://cloud.drone.io/hamburghammer/gstat)
[![Go Report Card](https://goreportcard.com/badge/github.com/hamburghammer/gstat)](https://goreportcard.com/report/github.com/hamburghammer/gstat)

Is a tool to get the system stats in a parsable format.
The tool is part of the competition with [Niklas](https://github.com/nhh).

Checkout his solution: [cstat](https://github.com/nhh/cstat)

## Terms of the competition

### Allgemein:

Name des Programms: 
    - cstat 
    - gstat

Das Tool kann mehrere Metriken auf einmal zurückgeben.

### Anforderung:

    - Aktuelle CPU Auslastung
    - Gesamtverbrauch aller CPU Kerne in %
    - Die nach cpu sortierten Prozesse als Liste
    - Aktueller Speicherplatzverbrauch
        -used / free in megabyte
    - Aktueller RAM Verbrauch
        - used / free in megabyte
    - Healthchecks mit Latency
        - http / https / *ICMP*
        - GET /
    - *Aktueller Network IO (optional)*
    - *Disk IO (optional)*
    - JSON Output
    - Datum und Uhrzeit im ISO Format

### Kriterien:

    - Single Executable
    - Linux

### Bewertungs:

1. Wie groß ist das Binary
2. Performance von Befehlen
3. Cpu Auslastung
4. Ram Auslastung

### Kommandozeilenaufrufe:

Alphabetische Reihenfolge

-c -h -d // --healtheck=http://example.com --disk --format=json

[https://de.wikipedia.org/wiki/Uniform_Resource_Identifier](https://de.wikipedia.org/wiki/Uniform_Resource_Identifier)

`$ cstat/gstat --cpu --format=json`
`$ cstat/gstat --metric=disk`
`$ cstat/gstat --check`

#### Goals:

cstat/gstat -c -d -i -h https://my-server.com > log.json

curl -X POST -d $(cstat/gstat -c -d -i -h https://my-server.com) https://my-logging.com/logs

