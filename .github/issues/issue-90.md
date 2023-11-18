---
id: 90
database_id: 310234671
node_id: MDU6SXNzdWUzMTAyMzQ2NzE=
status: closed
title: "bug with migrations"
labels: ["type:bug"]
url: https://github.com/octopot/forma/issues/90
created_at: 2018-03-31T11:40:51Z
updated_at: 2018-05-05T08:49:40Z
---

# bug with migrations

```bash
$ make cmd-migrate-up
go run -ldflags '-s -w -X main.version=dev -X main.commit=b28c609' main.go build.go migrate  up 1
2018/03/31 14:39:12 Applied 1 migration(s)!
$ make cmd-migrate-up
go run -ldflags '-s -w -X main.version=dev -X main.commit=b28c609' main.go build.go migrate  up 1
2018/03/31 14:39:15 Applied 0 migration(s)!
$ make cmd-migrate-down
go run -ldflags '-s -w -X main.version=dev -X main.commit=b28c609' main.go build.go migrate  down 1
2018/03/31 14:39:24 Applied 1 migration(s)!
$ make cmd-migrate-down
go run -ldflags '-s -w -X main.version=dev -X main.commit=b28c609' main.go build.go migrate  down 1
2018/03/31 14:39:28 Applied 2 migration(s)!
$ make cmd-migrate-down
go run -ldflags '-s -w -X main.version=dev -X main.commit=b28c609' main.go build.go migrate  down 1
2018/03/31 14:39:30 Applied 2 migration(s)!
```
