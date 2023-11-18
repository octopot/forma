---
id: 181
database_id: 371095459
node_id: MDU6SXNzdWUzNzEwOTU0NTk=
status: open
title: "new cli command `separate`"
labels: ["type:feature"]
url: https://github.com/octopot/forma/issues/181
created_at: 2018-10-17T14:07:11Z
updated_at: 2018-10-17T14:07:11Z
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
