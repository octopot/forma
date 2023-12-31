---
code: null
id: MDU6SXNzdWUzNDMzMjgxMjQ=
databaseId: 343328124
number: 143
url: https://github.com/octopot/forma/issues/143
title: "provide token by flag and/or environment variable"
issueType: null
labels:
  - "type:task"
assignees:
  - kamilsk
milestone: "[[milestone-2]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-07-21T13:54:58Z
updatedAt: 2018-07-25T19:02:11Z
lastEditedAt: null
closedAt: 2018-07-25T19:02:11Z
---

# provide token by flag and/or environment variable

- `form-api ctl create -f some.yaml --token="my secret"`
- `export FORMA_TOKEN=""; form-api ctl create -f some.yaml`
