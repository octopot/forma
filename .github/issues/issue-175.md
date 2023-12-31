---
code: null
id: MDU6SXNzdWUzNTkzNzg2MDA=
databaseId: 359378600
number: 175
url: https://github.com/octopot/forma/issues/175
title: "right listen"
issueType: null
labels:
  - "type:refactoring"
  - "status:declined"
assignees:
  - kamilsk
milestone: "[[milestone-3]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-09-12T08:54:39Z
updatedAt: 2018-10-17T13:38:41Z
lastEditedAt: null
closedAt: 2018-10-16T08:14:12Z
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

<!-- 2018-10-16T08:14Z https://github.com/octopot/forma/issues/175#issuecomment-430143896
won't do
-->
