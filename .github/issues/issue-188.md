---
code: null
id: MDU6SXNzdWUzOTIwNzYyMjg=
databaseId: 392076228
number: 188
url: https://github.com/octopot/forma/issues/188
title: "wrap migration with transaction"
issueType: null
labels:
  - "type:refactoring"
assignees:
  - kamilsk
milestone: "[[milestone-3]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-12-18T09:54:43Z
updatedAt: 2018-12-20T17:28:17Z
lastEditedAt: null
closedAt: 2018-12-20T17:28:17Z
---

# wrap migration with transaction

```
begin;
...
commit;
```

<!-- 2018-12-20T17:28Z https://github.com/octopot/forma/issues/188#issuecomment-449074986
won't do
-->
