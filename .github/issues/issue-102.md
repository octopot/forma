---
code: null
id: MDU6SXNzdWUzMjQ2OTAzMjI=
databaseId: 324690322
number: 102
url: https://github.com/octopot/forma/issues/102
title: "use # instead of ? for redirect"
issueType: null
labels:
  - "type:bug"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-05-20T07:38:09Z
updatedAt: 2018-05-21T08:36:52Z
lastEditedAt: null
closedAt: 2018-05-21T08:36:52Z
---

# use # instead of ? for redirect

I send just one message, but url has previos result: `?629c995b-3dc1-4484-8e6c-f74007645bcd=success&b3a7610d-a982-4ede-b702-f556592f8204=success`

<!-- 2018-05-20T07:39Z https://github.com/octopot/forma/issues/102#issuecomment-390463718
also fix script in /docs
-->

<!-- 2018-05-21T08:34Z https://github.com/octopot/forma/issues/102#issuecomment-390589911
`window.JSON.parse(window.atob('#'))` -> `{id: "41ca5e09-3ce2-4094-b108-3ecc257c6fa4", result: "success"}`
-->
