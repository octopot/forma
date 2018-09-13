OPEN_BROWSER       =
SUPPORTED_VERSIONS = 1.9 1.10 1.11 latest


include env/makefiles/env.mk
include env/makefiles/docker.mk
include env/makefiles/local.mk
include env/cmd.mk
include env/docker.mk
include env/docker-compose.mk
include env/tools.mk


.PHONY: code-quality-check
code-quality-check: ARGS = \
	--exclude=".*_test\.go:.*error return value not checked.*\(errcheck\)$$" \
	--exclude="duplicate of.*_test.go.*\(dupl\)$$" \
	--exclude="static/bindata.go" \
	--exclude="mock_.*.go" \
	--vendor --deadline=5m ./... | sort
code-quality-check: docker-tool-gometalinter

.PHONY: code-quality-report
code-quality-report:
	time make code-quality-check | tail +7 | tee report.out


.PHONY: dev
dev: up stop-server stop-service clear status demo dev-server
