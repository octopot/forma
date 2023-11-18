---
id: 26
database_id: 276515001
node_id: MDU6SXNzdWUyNzY1MTUwMDE=
status: closed
title: "support accept header"
labels: ["type:refactoring"]
url: https://github.com/octopot/forma/issues/26
created_at: 2017-11-24T04:58:41Z
updated_at: 2017-12-12T16:14:57Z
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
