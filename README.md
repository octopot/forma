> # 🗃 Forma
>
> Data Collector as a Service &mdash; your personal server for HTML forms.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

## Roadmap

- [x] v1: [MVP][project_v1]
  - [**May 31, 2018**][project_v1_dl]
  - Main concepts and working prototype.
- [x] v2: [Accounts and CLI CRUD][project_v2]
  - [**August 31, 2018**][project_v2_dl]
  - Command line interface for create, read, update and delete operations above gRPC.
- [ ] v3: [Templating and RESTful API][project_v3]
  - [**Someday, 20xx**][project_v3_dl]
  - Template system and Edge Side Includes/Server Side Includes support.
  - Integrate gRPC gateway.
  - Improve gRPC layer.
- [ ] v4: [DSL for validation and CSI][project_v4]
  - [**Sometime, 20xx**][project_v4_dl]
  - Domain-specific language to define validation rules.
  - Client-side integration.
  - Graphical user interface and admin panel to perform create, read, update and delete operations.

## Motivation

- We need better integration with static sites built with [Hugo](https://gohugo.io/).
- We want cheaper products than [Formata](https://www.formata.io/) or [FormKeep](https://formkeep.com/).
- We have to full control over our users' data and protect it from third parties.

## Quick start

Requirements:

- Docker 18.06.0-ce or above
- Docker Compose 1.22.0 or above
- Go 1.9.2 or above
- GNU Make 3.81 or above

```bash
$ make up demo status

     Name                    Command               State                          Ports
---------------------------------------------------------------------------------------------------------------
forma_db_1        docker-entrypoint.sh postgres    Up      0.0.0.0:5432->5432/tcp
forma_server_1    /bin/sh -c echo $BASIC_USE ...   Up      0.0.0.0:443->443/tcp, 0.0.0.0:80->80/tcp
forma_service_1   service run --with-profili ...   Up      0.0.0.0:8080->80/tcp, 0.0.0.0:8090->8090/tcp,
                                                           0.0.0.0:8091->8091/tcp, 0.0.0.0:8092->8092/tcp

$ open http://127.0.0.1.xip.io/api/v1/10000000-2000-4000-8000-160000000004

$ make help
```

<details>
<summary><strong>GET curl /api/v1/UUID</strong></summary>

```bash
$ curl http://127.0.0.1.xip.io/api/v1/10000000-2000-4000-8000-160000000004
# <form id="10000000-2000-4000-8000-160000000004" lang="en" title="Email Subscription"
#       action="http://localhost/api/v1/10000000-2000-4000-8000-160000000004" method="POST"
#       enctype="application/x-www-form-urlencoded">
#       <input id="10000000-2000-4000-8000-160000000004_email" name="email" type="email" title="Email"
#              maxlength="64" required="true"></input>
#       <input type="submit">
# </form>
```
</details>

<details>
<summary><strong>POST /api/v1/UUID</strong></summary>

```bash
$ curl -v -H "Content-Type: application/x-www-form-urlencoded" \
       --data-urlencode "email=test@my.email" \
       http://127.0.0.1.xip.io/api/v1/10000000-2000-4000-8000-160000000004
# > POST /api/v1/10000000-2000-4000-8000-160000000004 HTTP/1.1
# > Host: 127.0.0.1.xip.io
# > User-Agent: curl/7.54.0
# > Accept: */*
# > Content-Type: application/x-www-form-urlencoded
# > Content-Length: 21
# >
# < HTTP/1.1 302 Found
# < Location: http://localhost/api/v1/10000000-2000-4000-8000-160000000004#eyJpbnB1dCI6ImJmM2MyYWIwLWVkYjQtNDFiZi1iNDlkLWY3ZjNiMmI5ZDViMiIsImlkIjoiMTAwMDAwMDAtMjAwMC00MDAwLTgwMDAtMTYwMDAwMDAwMDA0IiwicmVzdWx0Ijoic3VjY2VzcyJ9
# < Date: Sat, 05 May 2018 09:34:47 GMT
# < Content-Length: 0
# <
```
</details>

## Specification

### API

You can find API specification [here](env/client/rest.http). Also, we recommend using [Insomnia](https://insomnia.rest/)
HTTP client to work with the API - you can import data for it from the [file](env/client/insomnia.json).
Or you can choose [Postman](https://www.getpostman.com/) - its import data is [here](env/client/postman.json) and
[here](env/client/postman.env.json).

### CLI

You can use CLI not only to start the HTTP server but also to execute
[CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) operations.

<details>
<summary><strong>Service command-line interface</strong></summary>

```bash
$ make install

$ form-api help
Forma

Usage:
  form-api [command]

Available Commands:
  completion  Print Bash or Zsh completion
  ctl         Forma Service Control
  help        Help about any command
  migrate     Apply database migration
  run         Start HTTP server
  version     Show application version

Flags:
  -h, --help   help for form-api

Use "form-api [command] --help" for more information about a command.
```
</details>

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
$ export REQ_VER=2.0.0  # all available versions are on https://github.com/kamilsk/form-api/releases/
$ export REQ_OS=Linux   # macOS and Windows are also available
$ export REQ_ARCH=64bit # 32bit is also available
# wget -q -O forma.tar.gz
$ curl -sL -o forma.tar.gz \
       https://github.com/kamilsk/form-api/releases/download/"${REQ_VER}/form-api_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf forma.tar.gz -C "${GOPATH}"/bin/ && rm forma.tar.gz
```

### Docker Hub

```bash
$ docker pull kamilsk/form-api:2.x
# or use mirror
$ docker pull quay.io/kamilsk/form-api:2.x
```

### From source code

```bash
$ egg github.com/kamilsk/form-api@^2.0.0 -- make test install
# or use mirror
$ egg bitbucket.org/kamilsk/form-api@^2.0.0 -- make test install
```

> [egg](https://github.com/kamilsk/egg)<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/octopot/forma
[build.icon]:       https://travis-ci.com/octopot/forma.svg?branch=master
[design.page]:      https://www.notion.so/octolab/Forma-713aa8203eaf474e8f4ae639b930d36f?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://octopot.github.io/forma/
[template.page]:    https://github.com/octomation/go-service
[template.icon]:    https://img.shields.io/badge/template-go--service-blue

[egg]:              https://github.com/kamilsk/egg

[project_v1]:       https://github.com/octopot/forma/projects/1
[project_v1_dl]:    https://github.com/octopot/forma/milestone/1
[project_v2]:       https://github.com/octopot/forma/projects/2
[project_v2_dl]:    https://github.com/octopot/forma/milestone/2
[project_v3]:       https://github.com/octopot/forma/projects/3
[project_v3_dl]:    https://github.com/octopot/forma/milestone/3
[project_v4]:       https://github.com/octopot/forma/projects/4
[project_v4_dl]:    https://github.com/octopot/forma/milestone/4
