---
code: null
id: MDU6SXNzdWUyODUzODA5NDM=
databaseId: 285380943
number: 56
url: https://github.com/octopot/forma/issues/56
title: "add encoding.Indent to tests"
issueType: null
labels:
  - "type:task"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-01-02T08:59:25Z
updatedAt: 2018-01-24T20:23:59Z
lastEditedAt: 2018-01-02T09:00:29Z
closedAt: 2018-01-24T20:23:58Z
---

# add encoding.Indent to tests

For better viewing

Here, for example: https://github.com/kamilsk/form-api/blob/master/data/form/serialization_test.go#L194

`json.MarshalIndent(..., "", "  ")`

<!-- 2018-01-02T09:00Z https://github.com/octopot/forma/issues/56#issuecomment-354724059
Also, support it by `MarshalHTML`
-->

<!-- 2018-01-02T09:02Z https://github.com/octopot/forma/issues/56#issuecomment-354724296
And add usage of `Compact()`
-->
