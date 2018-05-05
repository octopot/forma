> # Form API [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Data%20Collector%20as%20a%20Service&url=https://kamilsk.github.io/form-api/&via=ikamilsk&hashtags=go,service,data-collector)
> [![Analytics](https://ga-beacon.appspot.com/UA-109817251-15/form-api/readme?pixel)](https://kamilsk.github.io/form-api/)
> Data Collector as a Service &mdash; your personal server for HTML forms.

[![Patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://www.patreon.com/octolab)
[![Build Status](https://travis-ci.org/kamilsk/form-api.svg?branch=master)](https://travis-ci.org/kamilsk/form-api)
[![Code Coverage](https://scrutinizer-ci.com/g/kamilsk/form-api/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/kamilsk/form-api/?branch=master)
[![Code Quality](https://scrutinizer-ci.com/g/kamilsk/form-api/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/kamilsk/form-api/?branch=master)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Quick start

Requirements: 

- Docker 17.09.0-ce or above
- Docker Compose 1.16.1 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

```bash
$ make up demo status

       Name                     Command               State                                  Ports
----------------------------------------------------------------------------------------------------------------------------------
form-api_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
form-api_server_1    /bin/sh -c envsubst '$SERV ...   Up      80/tcp, 0.0.0.0:80->8080/tcp
form-api_service_1   form-api run --with-profil ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp, 0.0.0.0:8091->8091/tcp

$ curl http://localhost:8080/api/v1/41ca5e09-3ce2-4094-b108-3ecc257c6fa4
# <form id="41ca5e09-3ce2-4094-b108-3ecc257c6fa4" lang="en" title="Email subscription"
#       action="http://localhost/api/v1/41ca5e09-3ce2-4094-b108-3ecc257c6fa4" method="post"
#       enctype="application/x-www-form-urlencoded">
#       <input id="41ca5e09-3ce2-4094-b108-3ecc257c6fa4_email" name="email" type="email" title="Email"
#              maxlength="64" required="true"></input>
# </form>
$ curl -v -H "Content-Type: application/x-www-form-urlencoded" \
       --data-urlencode "email=test@my.email" \
       http://localhost:8080/api/v1/41ca5e09-3ce2-4094-b108-3ecc257c6fa4
# > POST /api/v1/41ca5e09-3ce2-4094-b108-3ecc257c6fa4 HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/7.54.0
# > Accept: */*
# > Content-Type: application/x-www-form-urlencoded
# > Content-Length: 21
# >
# < HTTP/1.1 302 Found
# < Location: https://kamil.samigullin.info/?41ca5e09-3ce2-4094-b108-3ecc257c6fa4=success
# < Date: Sat, 05 May 2018 09:34:47 GMT
# < Content-Length: 0
# <
$
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
  form-api [command]

Available Commands:
  completion  Print Bash or Zsh completion
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for form-api

Use "form-api [command] --help" for more information about a command.
```

## Installation

### Brew

```bash
$ brew install kamilsk/tap/form-api
```

### Binary

```bash
$ export VER=1.0.0      # all available versions are on https://github.com/kamilsk/form-api/releases
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
$ egg github.com/kamilsk/form-api@^1.0.0 -- make test install
```

#### Mirror

```bash
$ egg bitbucket.org/kamilsk/form-api@^1.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg) is an `extended go get`.

### Bash and Zsh completions

You can find completion files [here](https://github.com/kamilsk/shared/tree/dotfiles/bash_completion.d) or
build your own using these commands

```bash
$ form-api completion bash > /path/to/bash_completion.d/form-api.sh
$ form-api completion zsh  > /path/to/zsh-completions/_form-api.zsh
```

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
