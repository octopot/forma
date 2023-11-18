---
id: 32
database_id: 277348085
node_id: MDU6SXNzdWUyNzczNDgwODU=
status: closed
title: "schema.Validate"
labels: ["type:task"]
url: https://github.com/octopot/forma/issues/32
created_at: 2017-11-28T11:20:20Z
updated_at: 2017-12-05T08:31:20Z
---

# schema.Validate

Continue of issue https://github.com/kamilsk/form-api/issues/29

```
var data url.Values = request.PostForm()
filtered := schema.Apply(data)
validated := schema.Validate(filtered)
```
