---
id: 47
database_id: 281378605
node_id: MDU6SXNzdWUyODEzNzg2MDU=
status: open
title: "add target for stress test"
labels: ["type:task"]
url: https://github.com/octopot/forma/issues/47
created_at: 2017-12-12T13:10:20Z
updated_at: 2018-08-31T15:40:39Z
---

# add target for stress test

```
$ make stress-test

semaphore create 2
semaphore add -- curl -O http://user:pass@localhost:8080/debug/
semaphore add -- hey -cpu=2 localhost:8080/api/v1/{schema}
semaphore wait
```
