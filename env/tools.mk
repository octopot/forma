.PHONY: tools
tools:
	if ! command -v mockgen > /dev/null; then \
	    go get github.com/golang/mock/mockgen; \
	fi
	# quick fix of https://github.com/kamilsk/form-api/issues/70
	# https://github.com/jteeuwen/go-bindata/compare/master...a-urth:master
	if ! command -v go-bindata > /dev/null; then \
	    go get -d github.com/a-urth/go-bindata/go-bindata; \
	    cd $(GOPATH)/src/github.com/a-urth/go-bindata && git checkout df38da1; \
	    go install github.com/a-urth/go-bindata/go-bindata; \
	fi

.PHONY: generate
generate: tools
	find . -name mock_*.go | grep -v ./vendor | xargs rm || true
	go generate ./...

.PHONY: static
static: tools
	go-bindata -o static/bindata.go -pkg static -ignore "^.+\.go$$" -ignore "static/fixtures" static/...
