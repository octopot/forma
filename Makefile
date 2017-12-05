include env/docker.mk
include env/docker-compose.mk

.PHONY: tools
tools:
	if ! command -v dep > /dev/null; then \
	    go get github.com/golang/dep/cmd/dep; \
	fi
	if ! command -v mockgen > /dev/null; then \
	    go get github.com/golang/mock/gomock; \
	    go get github.com/golang/mock/mockgen; \
	fi
	if ! command -v go-bindata > /dev/null; then \
	    go get github.com/jteeuwen/go-bindata/go-bindata; \
	fi

.PHONY: deps
deps: tools
	dep ensure -v

.PHONY: generate
generate: tools
	find . -name mock_*.go | grep -v ./vendor | xargs rm
	go generate ./...

.PHONY: static
static: tools
	go-bindata -o static/static.go -pkg static -ignore "^.+\.go$$" static/...

.PHONY: test
test: generate
	go test ./...

.PHONY: test-detailed
test-detailed:
	go test -cover -v ./...

.PHONY: test-with-race
test-with-race:
	go test -race ./...

.PHONY: test-with-coverage-formatted
test-formatted: generate
	go test -cover ./... | column -t | sort -r



.PHONY: run
run: BIND = 127.0.0.1
run: PORT = 8080
run:
	( \
	  export BIND=$(BIND) PORT=$(PORT); \
	  go run -ldflags '-s -w -X main.version=dev -X main.commit=unknown -X main.date=unknown' main.go build.go $(COMMAND); \
	)

.PHONY: help
help: COMMAND = help
help: run

.PHONY: help-migrate
help-migrate: COMMAND = migrate --help
help-migrate: run

.PHONY: help-run
help-run: COMMAND = run --help
help-run: run

.PHONY: migrate
migrate: COMMAND = migrate
migrate: run

.PHONY: migrate-up
migrate-up: migrate

.PHONY: migrate-down
migrate-down: COMMAND = migrate down
migrate-down: run

.PHONY: server
server: COMMAND = run --with-profiler
server: run

.PHONY: version
version: COMMAND = version
version: run



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
