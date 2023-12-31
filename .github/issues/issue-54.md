---
code: null
id: MDU6SXNzdWUyODUyMjUzOTk=
databaseId: 285225399
number: 54
url: https://github.com/octopot/forma/issues/54
title: "do not use http.ListenAndServe, use http.Server instead"
issueType: null
labels:
  - "type:critical"
  - "type:refactoring"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2017-12-30T21:41:33Z
updatedAt: 2018-01-23T12:41:15Z
lastEditedAt: 2017-12-30T21:43:49Z
closedAt: 2018-01-23T12:41:15Z
---

# do not use http.ListenAndServe, use http.Server instead

https://github.com/kamilsk/form-api/blob/master/cmd/run.go#L23

configure it, use wrapper for logging instead of https://github.com/kamilsk/form-api/blob/master/cmd/run.go#L22
