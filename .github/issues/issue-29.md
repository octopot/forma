---
code: null
id: MDU6SXNzdWUyNzY1MzEzMTM=
databaseId: 276531313
number: 29
url: https://github.com/octopot/forma/issues/29
title: "schema.Apply"
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
createdAt: 2017-11-24T07:08:56Z
updatedAt: 2017-11-28T11:20:29Z
lastEditedAt: 2017-11-24T07:10:00Z
closedAt: 2017-11-28T11:20:29Z
---

# schema.Apply

```
var data url.Values = request.PostForm()
filtered := schema.Apply(data)
validated := schema.Validate(filtered)
```
