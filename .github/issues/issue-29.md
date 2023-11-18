---
id: 29
database_id: 276531313
node_id: MDU6SXNzdWUyNzY1MzEzMTM=
status: closed
title: "schema.Apply"
labels: ["type:task"]
url: https://github.com/octopot/forma/issues/29
created_at: 2017-11-24T07:08:56Z
updated_at: 2017-11-28T11:20:29Z
---

# schema.Apply

```
var data url.Values = request.PostForm()
filtered := schema.Apply(data)
validated := schema.Validate(filtered)
```
