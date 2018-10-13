_commit     = -X main.commit=$(shell git rev-parse --short HEAD)
_date       = -X main.date=$(shell date -u +%FT%X%Z)
_version    = -X main.version=dev

LDFLAGS     = -ldflags '-s -w $(_commit) $(_date) $(_version)'
BUILD_FILES = main.go


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

.PHONY: cmd-migrate
cmd-migrate:
	go run $(LDFLAGS) $(BUILD_FILES) migrate $(FLAGS) up

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
	go run $(LDFLAGS) $(BUILD_FILES) run -H 127.0.0.1:8080 --with-profiling --with-monitoring
