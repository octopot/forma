.PHONY: run
run: LDFLAGS = -s -w -X main.version=dev -X main.commit=$(shell git rev-parse --short HEAD)
run:
	go run -ldflags '$(LDFLAGS)' main.go build.go $(COMMAND)


.PHONY: cmd-help
cmd-help: COMMAND = help
cmd-help: run

.PHONY: cmd-help-migrate
cmd-help-migrate: COMMAND = migrate --help
cmd-help-migrate: run

.PHONY: cmd-help-run
cmd-help-run: COMMAND = run --help
cmd-help-run: run

.PHONY: cmd-version
cmd-version: COMMAND = version
cmd-version: run


.PHONY: migrate-up
migrate-up: FLAGS   =
migrate-up: COMMAND = migrate $(FLAGS) up 1
migrate-up: run

.PHONY: migrate-down
migrate-down: FLAGS   =
migrate-down: COMMAND = migrate $(FLAGS) down 1
migrate-down: run


.PHONY: server
server: COMMAND = run --port=8080 --with-profiler --with-monitoring
server: run
