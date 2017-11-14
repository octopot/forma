.DEFAULT: watch


.PHONY: docker-build-fast dbf
docker-build-fast:
	docker build -t kamilsk/form-api:latest \
	             --force-rm --no-cache --pull --rm \
	             --build-arg QUICK=true \
	             .
dbf: docker-build-fast

.PHONY: docker-build
docker-build:
	docker build -t kamilsk/form-api:latest \
	             --force-rm --no-cache --pull --rm \
	             .


.PHONY: docker-start watch
docker-start:
	docker run --rm -d \
	           --name form-api-dev \
	           --publish 8080:8080 \
	           kamilsk/form-api:latest
watch: docker-start docker-logs

.PHONY: docker-logs
docker-logs:
	docker logs -f form-api-dev

.PHONY: docker-stop
docker-stop:
	docker stop form-api-dev
