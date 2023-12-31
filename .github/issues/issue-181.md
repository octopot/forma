---
code: null
id: MDU6SXNzdWUzNzEwOTU0NTk=
databaseId: 371095459
number: 181
url: https://github.com/octopot/forma/issues/181
title: "new cli command `separate`"
issueType: null
labels:
  - "type:feature"
assignees:
  - kamilsk
milestone: "[[milestone-4]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2018-10-17T14:07:11Z
updatedAt: 2018-10-17T14:07:11Z
lastEditedAt: null
closedAt: null
---

# new cli command `separate`

```bash
$ form-api separate -f path/to/form/template.html
schema:
<form>...

template:
<div>{{ template forma }}</div>
```

what it must do:
- parse a html template
- extract a schema from it
- replace markup by template variables
