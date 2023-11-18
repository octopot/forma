---
id: 79
database_id: 303102835
node_id: MDU6SXNzdWUzMDMxMDI4MzU=
status: closed
title: "resource leak"
labels: ["type:critical"]
url: https://github.com/octopot/forma/issues/79
created_at: 2018-03-07T13:42:27Z
updated_at: 2018-03-25T19:46:54Z
---

# resource leak

- `rows.Close`
- `db.SetMaxOpenConns`
