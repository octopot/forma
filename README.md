> # Form API [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Data%20Collector%20as%20a%20Service&url=https://kamilsk.github.io/form-api/&via=ikamilsk&hashtags=go,service,data-collector)
> [![Analytics](https://ga-beacon.appspot.com/UA-109817251-15/form-api/readme?pixel)](https://kamilsk.github.io/form-api/)
> Data Collector as a Service.

[![Patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://www.patreon.com/octolab)
[![Build Status](https://travis-ci.org/kamilsk/form-api.svg?branch=master)](https://travis-ci.org/kamilsk/semaphore)
[![Coverage Status](https://coveralls.io/repos/github/kamilsk/form-api/badge.svg)](https://coveralls.io/github/kamilsk/form-api)
[![GoDoc](https://godoc.org/github.com/kamilsk/form-api?status.svg)](https://godoc.org/github.com/kamilsk/form-api)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Usage

```bash
$ make deps up demo && make status

     Name                    Command               State                Ports             
------------------------------------------------------------------------------------------
env_db_1          docker-entrypoint.sh postgres    Up       0.0.0.0:5432->5432/tcp        
env_migration_1   form-api migrate up              Exit 0                                 
env_server_1      /bin/sh -c envsubst '$SERV ...   Up       80/tcp, 0.0.0.0:8080->8080/tcp
env_service_1     form-api run --with-profiler     Up       0.0.0.0:8081->8080/tcp        

$ curl http://localhost:8080/api/v1/41ca5e09-3ce2-4094-b108-3ecc257c6fa4
$ curl -H "Content-Type: application/x-www-form-urlencoded" \
       --data-urlencode "email=test@my.email" \
       http://localhost:8080/api/v1/41ca5e09-3ce2-4094-b108-3ecc257c6fa4
```

## Specification

### API

You can find API specification [here](env/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest)
HTTP client to work with the API - you can import data for it from the [file](env/insomnia.json).

### CLI

```bash
$ form-api --help
Form API

Usage:
   [command]

Available Commands:
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for this command
```

## Installation

### Requirements

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

### From source

```bash
$ go get -d -u github.com/kamilsk/form-api
$ cd ${GOPATH}/src/github.com/kamilsk/form-api
$ make deps generate test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/form-api
```

> [egg](https://github.com/kamilsk/egg) is an `extended go get`.

## Feedback

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kamilsk/form-api)
[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

## Notes

- brief roadmap
  - [x] v1: MVP
  - [ ] v2: CRUD
  - [ ] v3: GUI
  - [ ] v4: API v2
  - [ ] v5: Scalability
  - [ ] v6: Integrability
  - [ ] v7: Performance and Redundancy
  - [ ] v8: Complexity
- tested on Go 1.9
- made with ❤️ by [OctoLab](https://www.octolab.org/)
