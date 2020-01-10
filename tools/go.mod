module go.octolab.org/ecosystem/forma/tools

go 1.11

require (
	github.com/golang/mock v1.3.1
	github.com/golangci/golangci-lint v1.22.2
	github.com/kamilsk/egg v0.0.8
	golang.org/x/tools v0.2.2
)

replace github.com/izumin5210/gex => github.com/kamilsk/gex v0.6.0-e3

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.1
