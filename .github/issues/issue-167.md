---
code: null
id: MDU6SXNzdWUzNTA4OTU5NTI=
databaseId: 350895952
number: 167
url: https://github.com/octopot/forma/issues/167
title: "make a decision about data types in storage layer"
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
createdAt: 2018-08-15T17:28:43Z
updatedAt: 2018-08-22T20:02:45Z
lastEditedAt: null
closedAt: 2018-08-22T20:02:44Z
---

# make a decision about data types in storage layer

I have dissymmetry:
- InputWrite.Write use JSON to marshal verified data
- InputReader.Read* don't restore original value and return a byte array
