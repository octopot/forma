> # Forma [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Data%20Collector%20as%20a%20Service&url=https://kamilsk.github.io/form-api/&via=ikamilsk&hashtags=go,service,data-collector,form-handler)
> [![Analytics](https://ga-beacon.appspot.com/UA-109817251-15/form-api/readme?pixel)](https://kamilsk.github.io/form-api/)
> Data Collector as a Service &mdash; your personal server for HTML forms.

[![Patreon](https://img.shields.io/badge/patreon-donate-orange.svg)](https://www.patreon.com/octolab)
[![Build Status](https://travis-ci.org/kamilsk/form-api.svg?branch=master)](https://travis-ci.org/kamilsk/form-api)
[![Code Coverage](https://scrutinizer-ci.com/g/kamilsk/form-api/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/kamilsk/form-api/?branch=master)
[![Code Quality](https://scrutinizer-ci.com/g/kamilsk/form-api/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/kamilsk/form-api/?branch=master)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Roadmap

- [x] v1: [MVP](https://github.com/kamilsk/form-api/projects/1)
  - [**May 31, 2018**](https://github.com/kamilsk/form-api/milestone/1)
  - Main concepts and working prototype.
- [ ] v2: [API v2 and CLI CRUD](https://github.com/kamilsk/form-api/projects/2)
  - [**August 31, 2018**](https://github.com/kamilsk/form-api/milestone/2)
  - Command line interface for create, read, update and delete operations above gRPC.
  - Template system and Edge Side Includes/Server Side Includes support.
- [ ] v3: [DSL for validation and CSI](https://github.com/kamilsk/form-api/projects/3)
  - [**September 30, 2018**](https://github.com/kamilsk/form-api/milestone/3)
  - Client-side integration.
  - Domain-specific language to define validation rules.
- [ ] v4: [GUI CRUD](https://github.com/kamilsk/form-api/projects/4)
  - [**October 31, 2018**](https://github.com/kamilsk/form-api/milestone/4)
  - Graphical user interface and admin panel to perform create, read, update and delete operations.
- [ ] Forma, SaaS
  - **December 31, 2018**
  - Ready to apply on Cloud.
  - Move to [OctoLab](https://github.com/octolab/) organization and rename project to `forma`.

## Motivation

- We need better integration with static site generators like [Hugo](https://gohugo.io/).
- We want cheaper products than [Formata](https://www.formata.io/) or [FormKeep](https://formkeep.com/).
- We have to full control over our users' data and protect it from third parties.

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
```

## Specification

### API

You can find API specification [here](env/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest)
HTTP client to work with the API - you can import data for it from the [file](env/insomnia.json).

### CLI

```bash
$ form-api --help
Forma

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

#### Bash and Zsh completions

You can find completion files [here](https://github.com/kamilsk/shared/tree/dotfiles/bash_completion.d) or
build your own using these commands

```bash
$ form-api completion -f bash > /path/to/bash_completion.d/form-api.sh
$ form-api completion -f zsh  > /path/to/zsh-completions/_form-api.zsh
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
$ docker pull kamilsk/form-api:1.x
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

## Update

This application is in a state of [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) and under active
development. [SemVer](https://semver.org/) is used for releases, and you can easily be updated within minor versions,
but major versions can be not [BC](https://en.wikipedia.org/wiki/Backward_compatibility)-safe.

## Notes

- [research](../../tree/research)
- tested on Go 1.9 and 1.10

---

[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kamilsk/form-api)
[![@kamilsk](https://img.shields.io/badge/author-%40kamilsk-blue.svg)](https://twitter.com/ikamilsk)
[![@octolab](https://img.shields.io/badge/sponsor-%40octolab-blue.svg)](https://twitter.com/octolab_inc)

made with ❤️ by [OctoLab](https://www.octolab.org/)
