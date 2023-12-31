---
code: null
id: MDU6SXNzdWUyNzczNDgwODU=
databaseId: 277348085
number: 32
url: https://github.com/octopot/forma/issues/32
title: "schema.Validate"
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
createdAt: 2017-11-28T11:20:20Z
updatedAt: 2017-12-05T08:31:20Z
lastEditedAt: null
closedAt: 2017-12-05T08:31:20Z
---

# schema.Validate

Continue of issue https://github.com/kamilsk/form-api/issues/29

```
var data url.Values = request.PostForm()
filtered := schema.Apply(data)
validated := schema.Validate(filtered)
```
