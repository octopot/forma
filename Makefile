OPEN_BROWSER       =
SUPPORTED_VERSIONS = 1.9 latest

include makes/env.mk
include makes/local.mk
include makes/docker.mk
include cmd/Makefile
include env/docker.mk
include env/docker-compose.mk

.PHONY: code-quality-check
code-quality-check: ARGS = \
	--exclude=".*_test\.go:.*error return value not checked.*\(errcheck\)$$" \
	--exclude="duplicate of.*_test.go.*\(dupl\)$$" \
	--exclude="static/bindata.go" \
	--vendor --deadline=5m ./... | sort
code-quality-check: docker-tool-gometalinter

.PHONY: code-quality-report
code-quality-report:
	time make code-quality-check | tail +7 | tee report.out | pbcopy


.PHONY: tools
tools:
	if ! command -v dep > /dev/null; then \
	    go get github.com/golang/dep/cmd/dep; \
	fi
	if ! command -v mockgen > /dev/null; then \
	    go get github.com/golang/mock/mockgen; \
	fi
	if ! command -v go-bindata > /dev/null; then \
	    # quick fix of https://github.com/kamilsk/form-api/issues/70
	    # https://github.com/jteeuwen/go-bindata/compare/master...a-urth:master
	    go get -d github.com/a-urth/go-bindata/go-bindata; \
	    cd $GOPATH/src/github.com/a-urth/go-bindata && git checkout df38da164efcd92b3da59d8199c91ab90d2556bc; \
	    go install github.com/a-urth/go-bindata/go-bindata; \
	fi

.PHONY: deps
deps: tools
	dep ensure -v

.PHONY: generate
generate: tools
	find . -name mock_*.go | grep -v ./vendor | xargs rm || true
	go generate ./...

.PHONY: static
static: tools
	go-bindata -o static/bindata.go -pkg static -ignore "^.+\.go$$" -ignore "static/fixtures" static/...


.PHONY: pull-github-tpl
pull-github-tpl:
	rm -rf .github
	git clone git@github.com:kamilsk/shared.git .github
	( \
	  cd .github && \
	  git checkout github-tpl-go-v1 && \
	  git branch -d master && \
	  echo '- ' $$(cat README.md | head -n1 | awk '{print $$3}') 'at revision' $$(git rev-parse HEAD) \
	)
	rm -rf .github/.git .github/README.md

.PHONY: pull-makes
pull-makes:
	rm -rf makes
	git clone git@github.com:kamilsk/shared.git makes
	( \
	  cd makes && \
	  git checkout makefile-go-v1 && \
	  git branch -d master && \
	  echo '- ' $$(cat README.md | head -n1 | awk '{print $$3}') 'at revision' $$(git rev-parse HEAD) \
	)
	rm -rf makes/.git
