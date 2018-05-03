> # Form API [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Data%20Collector%20as%20a%20Service&url=https://kamilsk.github.io/form-api/&via=ikamilsk&hashtags=go,service,data-collector)
> [![Analytics](https://ga-beacon.appspot.com/UA-109817251-15/form-api/readme?pixel)](https://kamilsk.github.io/form-api/)
> Data Collector as a Service &mdash; your personal server for HTML forms.

[![Patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://www.patreon.com/octolab)
[![Build Status](https://travis-ci.org/kamilsk/form-api.svg?branch=master)](https://travis-ci.org/kamilsk/form-api)
[![Coverage Status](https://coveralls.io/repos/github/kamilsk/form-api/badge.svg)](https://coveralls.io/github/kamilsk/form-api)
[![GoDoc](https://godoc.org/github.com/kamilsk/form-api?status.svg)](https://godoc.org/github.com/kamilsk/form-api)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Quick start

```bash
$ make up demo status

    Name                   Command               State                                  Ports
-----------------------------------------------------------------------------------------------------------------------------
env_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
env_server_1    /bin/sh -c envsubst '$SERV ...   Up      80/tcp, 0.0.0.0:80->8080/tcp
env_service_1   form-api run --with-profil ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp, 0.0.0.0:8091->8091/tcp

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

Use " [command] --help" for more information about a command.
```

## Installation

### Brew

```bash
$ brew install kamilsk/tap/form-api
```

### Binary

```bash
$ export VER=1.1.2      # all available versions are on https://github.com/kamilsk/form-api/releases
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
$ wget -q -O form-api.tar.gz \
       https://github.com/kamilsk/form-api/releases/download/"${VER}/form-api_${VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf form-api.tar.gz -C "${GOPATH}"/bin/ && rm form-api.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/form-api:latest
```

### From source code

```bash
$ egg github.com/kamilsk/form-api@^1.0.0 -- make generate test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/form-api@^1.0.0 -- make generate test install
```

> [egg](https://github.com/kamilsk/egg) is an `extended go get`.

#### Requirements

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

## Notes

- brief roadmap
  - [x] v1: MVP
  - [ ] v2: API v2
  - [ ] v3: CSI
  - [ ] v4: CRUD
  - [ ] v5: GUI
  - [ ] Forma, SaaS
- tested on Go 1.9 and 1.10

---

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kamilsk/form-api)
[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)
