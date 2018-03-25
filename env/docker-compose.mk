.PHONY: docker-compose
docker-compose:
	cp -n env/.example.env env/.env || true
	cp env/.env .env # because https://docs.docker.com/compose/env-file/
	docker-compose -f env/docker-compose.yml $(COMMAND)



.PHONY: up
up: COMMAND = up -d
up: docker-compose

.PHONY: fresh-up
fresh-up: COMMAND = up --build --force-recreate -d
fresh-up: docker-compose

.PHONY: down
down: COMMAND = down
down: docker-compose

.PHONY: clean-down
clean-down: COMMAND = down --volumes --rmi local
clean-down: docker-compose

.PHONY: status
status: COMMAND = ps
status: docker-compose



.PHONY: start-db
start-db: COMMAND = start db
start-db: docker-compose

.PHONY: stop-db
stop-db: COMMAND = stop db
stop-db: docker-compose

.PHONY: logs-db
logs-db: COMMAND = logs -f db
logs-db: docker-compose

.PHONY: psql
psql: COMMAND = exec db /bin/sh -c 'su - postgres -c psql'
psql: docker-compose

.PHONY: backup
backup: COMMAND = exec db /bin/sh -c 'su - postgres -c "pg_dump --if-exists --clean $${POSTGRES_DB}"' > ./env/backup.sql
backup: docker-compose

.PHONY: restore
restore:
	cat ./env/backup.sql | docker exec -i $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1) \
	  /bin/sh -c 'cat > /tmp/backup.sql && su - postgres -c "psql --single-transaction $${POSTGRES_DB} < /tmp/backup.sql"'



.PHONY: start-migration
start-migration: COMMAND = start migration
start-migration: docker-compose

.PHONY: stop-migration
stop-migration: COMMAND = stop migration
stop-migration: docker-compose

.PHONY: logs-migration
logs-migration: COMMAND = logs -f migration
logs-migration: docker-compose



.PHONY: start-server
start-server: COMMAND = start server
start-server: docker-compose

.PHONY: stop-server
stop-server: COMMAND = stop server
stop-server: docker-compose

.PHONY: logs-server
logs-server: COMMAND = logs -f server
logs-server: docker-compose



.PHONY: start-service
start-service: COMMAND = start service
start-service: docker-compose

.PHONY: stop-service
stop-service: COMMAND = stop service
stop-service: docker-compose

.PHONY: logs-service
logs-service: COMMAND = logs -f service
logs-service: docker-compose



.PHONY: demo
demo: COMMAND = exec service form-api migrate up --with-demo
demo: docker-compose



.PHONY: rm-volumes
rm-volumes: down
	docker volume ls | tail +2 | awk '{print $$2}' | grep ^env_ | xargs docker volume rm
