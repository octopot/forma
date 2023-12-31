---
code: null
id: MDU6SXNzdWUzMDMxMDI4MzU=
databaseId: 303102835
number: 79
url: https://github.com/octopot/forma/issues/79
title: "resource leak"
issueType: null
labels:
  - "type:critical"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-03-07T13:42:27Z
updatedAt: 2018-03-25T19:46:54Z
lastEditedAt: null
closedAt: 2018-03-25T19:46:54Z
---

# resource leak

- `rows.Close`
- `db.SetMaxOpenConns`
