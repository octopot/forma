# execx
[![CI](https://github.com/izumin5210/execx/workflows/CI/badge.svg)](https://github.com/izumin5210/execx/actions?workflow=CI)
[![GoDoc](https://godoc.org/github.com/izumin5210/execx?status.svg)](https://godoc.org/github.com/izumin5210/execx)
[![License](https://img.shields.io/github/license/izumin5210/execx)](./LICENSE)


Make `os/exec` testable and graceful

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

cmd := execx.CommandContext(ctx, "sh", "-c", "sleep 5; echo done")
out, err := cmd.Output()

st := err.(*execx.ExitStatus)

fmt.Println(out, err, st.Signaled, st.Killed)

// Output: [] context deadline exceeded true false
```

## Reference

- [github.com/Songmu/timeout](https://godoc.org/github.com/Songmu/timeout)
- [k8s.io/utils/exec](https://godoc.org/k8s.io/utils/exec)
