---
code: null
id: MDU6SXNzdWUyNzY1MTUwMDE=
databaseId: 276515001
number: 26
url: https://github.com/octopot/forma/issues/26
title: "support accept header"
issueType: null
labels:
  - "type:refactoring"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2017-11-24T04:58:41Z
updatedAt: 2017-12-12T16:14:57Z
lastEditedAt: null
closedAt: 2017-12-12T16:14:57Z
---

# support accept header

parse:
```
Accept: text/html

Accept: image/*

Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8
```

```
/* TODO for GetV1
supported := map[string]string{
	"json": "application/json",
	//"toml": "application/toml",
	//"xml":  "application/xml",
	//"yaml": "application/yaml",
}
format := strings.ToLower(req.URL.Query().Get("format"))
if format == "" {
	format = "json"
}
ct, ok := supported[format]
if !ok {
	http.Error(rw, "Unsupported output format.", http.StatusUnsupportedMediaType)
	return
}
rw.Header().Set("Content-Type", ct)
fmt.Fprintf(rw, "get form schema in format %q with Content-Type %q", format, ct)
*/
```

<!-- 2017-11-24T16:53Z https://github.com/octopot/forma/issues/26#issuecomment-346868269
https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept
-->
