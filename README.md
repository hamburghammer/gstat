# gstat

[![Build Status](https://cloud.drone.io/api/badges/hamburghammer/gstat/status.svg?ref=refs/heads/master)](https://cloud.drone.io/hamburghammer/gstat)
[![Go Report Card](https://goreportcard.com/badge/github.com/hamburghammer/gstat)](https://goreportcard.com/report/github.com/hamburghammer/gstat)

Is a cli tool to get some system stats in a machine parsable format (JSON). It supports only Linux but might also run on some UNIX based operation systems.


This tool is part of a competition with [nhh](https://github.com/nhh) -> [Details](docs/completion.md)

**WIP: expect some changes for the stats gathering**

## Features
- Easy to use
- Single executable
- Runs on Linux

## Installation
For the time there are no binaries provided this means you need to install it through go.

Requirements:
- Go is installed.
- You have the `$GOPATH` defined.
- The `$GOPATH/bin` directory is in your `$PATH`.

Install and update it with `go get -u github.com/hamburghammer/gstat`.

## Usage
```
Usage:
  gstat [OPTIONS]

Application Options:
  -c, --cpu     Include the total CPU consumption.
  -m, --mem     Include the RAM usage.
  -d, --disk    Include the Disk usage.
  -p, --proc    Include the top 10 running processes with the highest CPU consumption.
      --health= Make a healthcheck call against the URI.

Help Options:
  -h, --help    Show this help message
```
*not all flags a jet supported or fully implemented!

example output:


`gstat -cmd`
```json
{"Date":"2020-11-21T16:32:18+01:00","CPU":3.49999999997029,"mem":{"used":5777,"total":16022},"disk":{"used":90319,"total":224323}}
```
