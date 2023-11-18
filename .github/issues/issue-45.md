---
id: 45
database_id: 280782010
node_id: MDU6SXNzdWUyODA3ODIwMTA=
status: closed
title: "code quality problems"
labels: ["type:refactoring"]
url: https://github.com/octopot/forma/issues/45
created_at: 2017-12-10T04:41:40Z
updated_at: 2017-12-13T04:34:43Z
---

# code quality problems

```
cmd/root.go:14:5:warning: exported var RootCmd should have comment or be unexported (golint)
dao/layer.go:15:36:warning: exported func Must returns unexported type *dao.layer, which can be annoying to use (golint)
dao/layer.go:24:36:warning: exported func New returns unexported type *dao.layer, which can be annoying to use (golint)
dao/layer_test.go:96:12:warning: ineffectual assignment to err (ineffassign)
dao/postgres/postgres.go:36:25:warning: uuid can be fmt.Stringer (interfacer)
errors/error.go:11:2:warning: exported const ServerError should have comment (or a comment on this block) or be unexported (golint)
errors/error.go:17:64:warning: exported func NotFound returns unexported type *errors.withCode, which can be annoying to use (golint)
errors/error.go:22:66:warning: exported func Validation returns unexported type *errors.withCode, which can be annoying to use (golint)
errors/error.go:27:64:warning: exported func Database returns unexported type *errors.withCode, which can be annoying to use (golint)
errors/error.go:32:69:warning: exported func Serialization returns unexported type *errors.withCode, which can be annoying to use (golint)
errors/error.go:6:2:warning: exported const ResourceNotFound should have comment (or a comment on this block) or be unexported (golint)
server/error.go:67:1:warning: comment on exported method Error.InvalidReferer should be of the form "InvalidReferer ..." (golint)
server/error.go:77:1:warning: exported method Error.FromGetV1 should have comment or be unexported (golint)
server/middleware.go:13:14:warning: next can be http.Handler (interfacer)
server/middleware.go:35:34:warning: error return value not checked (err.NotProvidedUUID().MarshalTo(rw)) (errcheck)
server/middleware.go:39:30:warning: error return value not checked (err.InvalidUUID().MarshalTo(rw)) (errcheck)
server/server.go:19:34:warning: exported func New returns unexported type *server.server, which can be annoying to use (golint)
server/server.go:37:46:warning: error return value not checked (httpErr.FromGetV1(response.Error).MarshalTo(rw)) (errcheck)
server/server.go:40:12:warning: error return value not checked (enc.Encode(response.Schema)) (errcheck)
server/server.go:50:41:warning: error return value not checked (httpErr.InvalidFormData(err).MarshalTo(rw)) (errcheck)
server/server.go:56:32:warning: error return value not checked (httpErr.NoReferer().MarshalTo(rw)) (errcheck)
server/server.go:62:40:warning: error return value not checked (httpErr.InvalidReferer(err).MarshalTo(rw)) (errcheck)
service/service.go:6:25:warning: exported func New returns unexported type *service.formAPI, which can be annoying to use (golint)
static/static.go:101:6:warning: func staticTemplateErrorHtml should be staticTemplateErrorHTML (golint)
static/static.go:72:5:warning: don't use underscores in Go names; var _staticMigrations1_initialSql should be _staticMigrations1InitialSQL (golint)
static/static.go:74:6:warning: don't use underscores in Go names; func staticMigrations1_initialSqlBytes should be staticMigrations1InitialSQLBytes (golint)
static/static.go:81:6:warning: don't use underscores in Go names; func staticMigrations1_initialSql should be staticMigrations1InitialSQL (golint)
static/static.go:92:5:warning: var _staticTemplateErrorHtml should be _staticTemplateErrorHTML (golint)
static/static.go:94:6:warning: func staticTemplateErrorHtmlBytes should be staticTemplateErrorHTMLBytes (golint)
```

need to create separate tasks for all packages
