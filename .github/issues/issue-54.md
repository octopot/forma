---
id: 54
database_id: 285225399
node_id: MDU6SXNzdWUyODUyMjUzOTk=
status: closed
title: "do not use http.ListenAndServe, use http.Server instead"
labels: ["type:critical","type:refactoring"]
url: https://github.com/octopot/forma/issues/54
created_at: 2017-12-30T21:41:33Z
updated_at: 2018-01-23T12:41:15Z
---

# do not use http.ListenAndServe, use http.Server instead

https://github.com/kamilsk/form-api/blob/master/cmd/run.go#L23

configure it, use wrapper for logging instead of https://github.com/kamilsk/form-api/blob/master/cmd/run.go#L22
