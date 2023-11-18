---
id: 175
database_id: 359378600
node_id: MDU6SXNzdWUzNTkzNzg2MDA=
status: closed
title: "right listen"
labels: ["type:refactoring","status:declined"]
url: https://github.com/octopot/forma/issues/175
created_at: 2018-09-12T08:54:39Z
updated_at: 2018-10-17T13:38:41Z
---

# right listen

```
func Listen(host string, port uint16) (net.Listener, error) {
  addr, addrErr := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
  if addrErr != nil {
    return nil, fmt.Errorf("Listen: %s", addrErr)
  }

  listener, listenError := net.ListenTCP("tcp", addr)
  if listenError != nil {
    return nil, fmt.Errorf("Listen: %s", listenError)
  }

  return listener, nil
}
```
