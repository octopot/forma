---
code: null
id: MDU6SXNzdWUzMTAyMzQ2NzE=
databaseId: 310234671
number: 90
url: https://github.com/octopot/forma/issues/90
title: "bug with migrations"
issueType: null
labels:
  - "type:bug"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-03-31T11:40:51Z
updatedAt: 2018-05-05T08:49:40Z
lastEditedAt: null
closedAt: 2018-05-05T08:19:23Z
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

<!-- 2018-03-31T11:44Z https://github.com/octopot/forma/issues/90#issuecomment-377687199
```
    id    |          applied_at
----------+-------------------------------
 data.sql | 2018-03-31 10:46:31.891721+00
```
-->

<!-- 2018-03-31T11:45Z https://github.com/octopot/forma/issues/90#issuecomment-377687281
the solution is to add `cmd-migrate-down: FLAGS = --with-demo`, but it is not clear.

```bash
$ make cmd-migrate-down
go run -ldflags '-s -w -X main.version=dev -X main.commit=b28c609' main.go build.go migrate --with-demo down 1
2018/03/31 14:44:49 Applied 2 migration(s)!
$ make cmd-migrate-down
go run -ldflags '-s -w -X main.version=dev -X main.commit=b28c609' main.go build.go migrate --with-demo down 1
2018/03/31 14:44:51 Applied 0 migration(s)!
```
-->

<!-- 2018-03-31T11:46Z https://github.com/octopot/forma/issues/90#issuecomment-377687354
and after that

```bash
$ make cmd-migrate-up
go run -ldflags '-s -w -X main.version=dev -X main.commit=b28c609' main.go build.go migrate  up 1
2018/03/31 14:46:10 pq: function "update_timestamp" already exists with same argument types handling 1_initial.sql
exit status 1
make: *** [cmd-migrate-up] Error 1
```

the reason is

```bash
ERROR:  cannot drop function update_timestamp() because other objects depend on it
DETAIL:  trigger form_data_updated on table form_schema depends on function update_timestamp()
HINT:  Use DROP ... CASCADE to drop the dependent objects too.
```

so `cmd-migrate-down` doesn't work properly
-->

<!-- 2018-03-31T11:50Z https://github.com/octopot/forma/issues/90#issuecomment-377687529
steps to reproduce

1. up with demo
1. drop without demo
1. drop with demo
1. up without demo
-->

<!-- 2018-04-01T18:54Z https://github.com/octopot/forma/issues/90#issuecomment-377808396
also I need to fix next logic:
```
$ make up-db
$ make backup-db
$ make up-migration
$ make restore-db

ERROR:  cannot drop extension plpgsql because other objects depend on it
DETAIL:  function public.update_timestamp() depends on language plpgsql
trigger form_data_updated on table public.form_schema depends on function public.update_timestamp()
HINT:  Use DROP ... CASCADE to drop the dependent objects too.
ERROR:  cannot drop schema public because other objects depend on it
DETAIL:  table public.migration depends on schema public
extension uuid-ossp depends on schema public
function public.update_timestamp() depends on schema public
type public.status depends on schema public
table public.form_schema depends on schema public
table public.form_data depends on schema public
HINT:  Use DROP ... CASCADE to drop the dependent objects too.
ERROR:  schema "public" already exists
```

but
```
$ make up-db up-migration backup
$ make cmd-migrate-down
$ make restore
```
works well (but I need to restore "clean" state)
-->

<!-- 2018-04-24T13:16Z https://github.com/octopot/forma/issues/90#issuecomment-383926292
`--with-demo` must not add `data.sql` into `migration` table. I have to change the mechanism of demo installation.
-->
