COMPOSE ?= docker-compose -f env/docker-compose.base.yml -f env/docker-compose.dev.yml -p forma


.PHONY: __env__
__env__:
	cp -n env/.env{.example,} || true # for containers
	cp -n env/.env .env       || true # for docker compose, https://docs.docker.com/compose/env-file/

.PHONY: rm-env
rm-env:
	find . -name .env | xargs rm -f || true


.PHONY: config
config:
	$(COMPOSE) config

.PHONY: up
up: __env__
	$(COMPOSE) up -d
	$(COMPOSE) rm -f

.PHONY: fresh-up
fresh-up: __env__
	$(COMPOSE) up --build --force-recreate -d
	$(COMPOSE) rm -f

.PHONY: down
down: __env__
	$(COMPOSE) down

.PHONY: clean-down
clean-down: __env__
	$(COMPOSE) down --volumes --rmi local

.PHONY: clear
clear: __env__
	$(COMPOSE) rm -f

.PHONY: status
status: __env__
	$(COMPOSE) ps


.PHONY: up-db
up-db: __env__
	$(COMPOSE) up -d db

.PHONY: start-db
start-db: __env__
	$(COMPOSE) start db

.PHONY: stop-db
stop-db: __env__
	$(COMPOSE) stop db

.PHONY: log-db
log-db: __env__
	$(COMPOSE) logs -f db


.PHONY: up-migration
up-migration: __env__
	$(COMPOSE) up --build -d migration

.PHONY: start-migration
start-migration: __env__
	$(COMPOSE) start migration

.PHONY: stop-migration
stop-migration: __env__
	$(COMPOSE) stop migration

.PHONY: log-migration
log-migration: __env__
	$(COMPOSE) logs -f migration


.PHONY: up-service
up-service: __env__
	$(COMPOSE) up --build -d service

.PHONY: start-service
start-service: __env__
	$(COMPOSE) start service

.PHONY: stop-service
stop-service: __env__
	$(COMPOSE) stop service

.PHONY: log-service
log-service: __env__
	$(COMPOSE) logs -f service


.PHONY: up-server
up-server: __env__
	$(COMPOSE) up -d server

.PHONY: start-server
start-server: __env__
	$(COMPOSE) start server

.PHONY: stop-server
stop-server: __env__
	$(COMPOSE) stop server

.PHONY: log-server
log-server: __env__
	$(COMPOSE) logs -f server
#|                    --- Database-specific commands
#|
.PHONY: psql
psql: __env__      #| Connect to the database.
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c psql')
#|
.PHONY: backup
backup: __env__    #| Backup the database.
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c "pg_dump --format=custom --file=/tmp/db.dump $${POSTGRES_DB}"')
	@(docker cp $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1):/tmp/db.dump ./env/docker/db/)
	@($(COMPOSE) exec db rm /tmp/db.dump)
	@(ls -l ./env/docker/db/db.dump)
#|
.PHONY: restore
restore: __env__   #| Restore the database.
	@(docker cp ./env/docker/db/reset.sql $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1):/tmp/)
	@(docker cp ./env/docker/db/db.dump $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1):/tmp/)
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c "psql $${POSTGRES_DB} < /tmp/reset.sql"')
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c "pg_restore -Fc -d $${POSTGRES_DB} /tmp/db.dump"')
	@($(COMPOSE) exec db rm /tmp/reset.sql /tmp/db.dump)
#|
.PHONY: truncate
truncate: __env__  #| Truncate the database tables.
	@(docker cp ./env/docker/db/truncate.sql $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1):/tmp/)
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c "psql $${POSTGRES_DB} < /tmp/truncate.sql"')
	@($(COMPOSE) exec db rm /tmp/truncate.sql)
#|
#|                    --- Server-specific commands
#|
.PHONY: reload
reload: __env__    #| Reload nginx configuration.
	@($(COMPOSE) exec server nginx -s reload)
