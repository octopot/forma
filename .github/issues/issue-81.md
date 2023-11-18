---
id: 81
database_id: 304682595
node_id: MDU6SXNzdWUzMDQ2ODI1OTU=
status: closed
title: "add useful method for Schema, Input"
labels: ["type:feature"]
url: https://github.com/octopot/forma/issues/81
created_at: 2018-03-13T09:17:49Z
updated_at: 2018-08-17T16:01:04Z
---

# add useful method for Schema, Input

```
{{ with .Schema.Input("email") }}
  {{ template "input" . }}
{{ end }}
```
