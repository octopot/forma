---
code: null
id: MDU6SXNzdWUzMDQ2ODI1OTU=
databaseId: 304682595
number: 81
url: https://github.com/octopot/forma/issues/81
title: "add useful method for Schema, Input"
issueType: null
labels:
  - "type:feature"
assignees:
  - kamilsk
milestone: "[[milestone-2]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-03-13T09:17:49Z
updatedAt: 2018-08-17T16:01:04Z
lastEditedAt: null
closedAt: 2018-08-17T16:01:03Z
---

# add useful method for Schema, Input

```
{{ with .Schema.Input("email") }}
  {{ template "input" . }}
{{ end }}
```
