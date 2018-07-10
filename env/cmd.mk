LDFLAGS     ?= -ldflags '-s -w -X main.version=dev -X main.commit=$(shell git rev-parse --short HEAD)'
BUILD_FILES ?= main.go


.PHONY: cmd-help
cmd-help:
	go run $(LDFLAGS) $(BUILD_FILES) help

.PHONY: cmd-help-migrate
cmd-help-migrate:
	go run $(LDFLAGS) $(BUILD_FILES) migrate --help

.PHONY: cmd-help-run
cmd-help-run:
	go run $(LDFLAGS) $(BUILD_FILES) run --help

.PHONY: cmd-version
cmd-version:
	go run $(LDFLAGS) $(BUILD_FILES) version

.PHONY: cmd-migrate-up
cmd-migrate-up:
	go run $(LDFLAGS) $(BUILD_FILES) migrate $(FLAGS) up 1

.PHONY: cmd-migrate-down
cmd-migrate-down:
	go run $(LDFLAGS) $(BUILD_FILES) migrate $(FLAGS) down 1


.PHONY: demo
demo: FLAGS = --with-demo
demo: cmd-migrate-up

.PHONY: dev-server
dev-server:
	go run $(LDFLAGS) $(BUILD_FILES) run --port=8080 --with-profiler --with-monitoring
