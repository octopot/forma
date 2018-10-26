# TODO issue#environment
# - protoc
# - go-bindata replace
# TODO add . "github.com/kamilsk/form-api/pkg/transport/grpc/protobuf" to the import

.PHONY: tools
tools:
	if ! command -v easyjson > /dev/null; then \
	    go get github.com/mailru/easyjson/...; \
	fi
	if ! command -v go-bindata > /dev/null; then \
	    go get -d github.com/a-urth/go-bindata/go-bindata; \
	    cd $(GOPATH)/src/github.com/a-urth/go-bindata && git checkout df38da1; \
	    go install github.com/a-urth/go-bindata/go-bindata; \
	fi
	if ! command -v mockgen > /dev/null; then \
	    go get github.com/golang/mock/mockgen; \
	fi

.PHONY: json
json:
	find . -name "*_easyjson.go" | grep -v /vendor/ | xargs rm || true
	go generate -run="easyjson" ./...

.PHONY: mocks
mocks:
	find . -name "mock_*_test.go" | grep -v /vendor/ | xargs rm || true
	go generate -run="mockgen" ./...

.PHONY: protobuf
protobuf:
	@(protoc -Ienv/api \
	         -Ivendor/github.com/grpc-ecosystem/grpc-gateway \
	         -Ivendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	         --go_out=plugins=grpc,logtostderr=true:pkg/transport/grpc/protobuf \
	         --grpc-gateway_out=logtostderr=true,import_path=gateway:pkg/transport/grpc/gateway \
	         --swagger_out=logtostderr=true,allow_merge=true,merge_file_name=forma:env/api \
	         common.proto storage.proto event.proto)
	@(mv env/api/forma.swagger.json env/api/swagger.json)

.PHONY: static
static:
	go-bindata -o pkg/static/bindata.go -pkg static -ignore "\.go$$" -ignore "fixtures" -prefix pkg/ pkg/static/...

.PHONY: generate
generate: tools json mocks protobuf static
