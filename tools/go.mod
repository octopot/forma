module go.octolab.org/ecosystem/forma/tools

go 1.12

require (
	github.com/golang/mock v1.4.3
	github.com/golangci/golangci-lint v1.24.0
	github.com/kamilsk/egg v0.0.14
	github.com/mailru/easyjson v0.7.1
	golang.org/x/tools v0.3.3
)

replace github.com/izumin5210/gex => github.com/kamilsk/gex v0.6.0-e4

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.3
