.PHONY: mocks
mocks:
	find . -name "mock_*.go" | grep -v ./vendor | xargs rm || true
	go generate -run="mockgen" ./...

.PHONY: tools
tools:
	if ! command -v go-bindata > /dev/null; then \
	    go get -d github.com/a-urth/go-bindata/go-bindata; \
	    cd $(GOPATH)/src/github.com/a-urth/go-bindata && git checkout df38da1; \
	    go install github.com/a-urth/go-bindata/go-bindata; \
	fi
	if ! command -v mockgen > /dev/null; then \
	    go get github.com/golang/mock/mockgen; \
	fi

.PHONY: generate
generate: tools mocks

.PHONY: static
static: tools
	go-bindata -o pkg/static/bindata.go -pkg static -ignore "\.go$$" -ignore "fixtures" -prefix pkg/ pkg/static/...
