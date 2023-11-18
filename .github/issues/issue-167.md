---
id: 167
database_id: 350895952
node_id: MDU6SXNzdWUzNTA4OTU5NTI=
status: closed
title: "make a decision about data types in storage layer"
labels: ["type:refactoring"]
url: https://github.com/octopot/forma/issues/167
created_at: 2018-08-15T17:28:43Z
updated_at: 2018-08-22T20:02:45Z
---

# make a decision about data types in storage layer

I have dissymmetry:
- InputWrite.Write use JSON to marshal verified data
- InputReader.Read* don't restore original value and return a byte array
