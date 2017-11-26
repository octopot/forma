.PHONY: docker-compose
docker-compose:
	docker-compose -f env/docker-compose.yml $(COMMAND)



.PHONY: up
up:
	docker-compose -f env/docker-compose.yml up -d
	docker-compose -f env/docker-compose.yml exec service form-api migrate

.PHONY: down
down: COMMAND = down
down: docker-compose

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



.PHONY: start-service
start-service: COMMAND = start service
start-service: docker-compose

.PHONY: stop-service
stop-service: COMMAND = stop service
stop-service: docker-compose

.PHONY: logs-service
logs-service: COMMAND = logs -f service
logs-service: docker-compose



.PHONY: start-server
start-server: COMMAND = start server
start-server: docker-compose

.PHONY: stop-server
stop-server: COMMAND = stop server
stop-server: docker-compose

.PHONY: logs-server
logs-server: COMMAND = logs -f server
logs-server: docker-compose
