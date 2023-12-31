---
code: null
id: MDU6SXNzdWUyODEzNzg2MDU=
databaseId: 281378605
number: 47
url: https://github.com/octopot/forma/issues/47
title: "add target for stress test"
issueType: null
labels:
  - "type:task"
assignees:
  - kamilsk
milestone: "[[milestone-4]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2017-12-12T13:10:20Z
updatedAt: 2018-08-31T15:40:39Z
lastEditedAt: null
closedAt: null
---

# add target for stress test

```
$ make stress-test

semaphore create 2
semaphore add -- curl -O http://user:pass@localhost:8080/debug/
semaphore add -- hey -cpu=2 localhost:8080/api/v1/{schema}
semaphore wait
```
