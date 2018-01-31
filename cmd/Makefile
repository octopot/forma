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
migrate-up: FLAGS   =
migrate-up: COMMAND = migrate $(FLAGS) up 1
migrate-up: run

.PHONY: migrate-down
migrate-down: FLAGS   =
migrate-down: COMMAND = migrate $(FLAGS) down 1
migrate-down: run

.PHONY: server
server: COMMAND = run --with-profiler
server: run

.PHONY: version
version: COMMAND = version
version: run
