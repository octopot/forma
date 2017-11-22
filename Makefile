include env/docker.mk
include env/docker-compose.mk

.PHONY: deps
deps:
	dep ensure -v

.PHONY: run
run: BIND = 127.0.0.1
run: PORT = 8080
run:
	( \
	  export BIND=$(BIND) PORT=$(PORT); \
	  go run -ldflags '-s -w -X main.version=dev -X main.commit=unknown -X main.date=unknown' main.go $(COMMAND); \
	)

.PHONY: help
help: COMMAND = help
help: run

.PHONY: migrate
migrate: COMMAND = migrate
migrate: run

.PHONY: service
service: COMMAND = run --with-profiler
service: run
