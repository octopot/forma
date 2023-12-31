---
code: null
id: MDU6SXNzdWUzNTA4OTY5NTA=
databaseId: 350896950
number: 168
url: https://github.com/octopot/forma/issues/168
title: "review domain layer"
issueType: null
labels:
  - "type:refactoring"
assignees:
  - kamilsk
milestone: "[[milestone-2]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-08-15T17:32:00Z
updatedAt: 2018-09-02T17:09:23Z
lastEditedAt: 2018-08-15T17:32:22Z
closedAt: 2018-09-02T17:09:23Z
---

# review domain layer

- UUID -> ID
- Schema knows nothing about XML
- InputData knows nothing about JSON

use them at storage layer directly
