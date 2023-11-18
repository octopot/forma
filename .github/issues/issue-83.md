---
id: 83
database_id: 305874005
node_id: MDU6SXNzdWUzMDU4NzQwMDU=
status: closed
title: "external configuration for router"
labels: ["type:refactoring"]
url: https://github.com/octopot/forma/issues/83
created_at: 2018-03-16T10:22:57Z
updated_at: 2018-03-30T14:09:59Z
---

# external configuration for router

add a possibility to configure route external

as example
```go
	if withProfiler {
		r.Route("/debug/pprof", func(r chi.Router) {
			r.Get("/", pprof.Index)
			r.Get("/cmdline", pprof.Cmdline)
			r.Get("/profile", pprof.Profile)
			r.Get("/symbol", pprof.Symbol)
			r.Get("/trace", pprof.Trace)
		})
	}
```
this is a not a part of chi router, also as `expvar`.
