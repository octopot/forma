.PHONY: docker-build-fast dbf
docker-build-fast:
	docker build -f env/Dockerfile \
	             -t kamilsk/form-api:latest \
	             --force-rm --no-cache --pull --rm \
	             --build-arg QUICK=true \
	             .
dbf: docker-build-fast

.PHONY: docker-build
docker-build:
	docker build -f env/Dockerfile \
	             -t kamilsk/form-api:latest \
	             --force-rm --no-cache --pull --rm \
	             .



.PHONY: docker-push
docker-push:
	docker push kamilsk/form-api:latest

.PHONY: docker-refresh
docker-refresh:
	docker images --all \
	| grep '^kamilsk\/form-api\s\+' \
	| awk '{print $$3}' \
	| xargs docker rmi -f &>/dev/null || true
	docker pull kamilsk/form-api:latest



.PHONY: docker-start
docker-start:
	docker run --rm -d \
	           --name form-api-dev \
	           --publish 8080:8080 \
	           kamilsk/form-api:latest

.PHONY: docker-logs
docker-logs:
	docker logs -f form-api-dev

.PHONY: docker-stop
docker-stop:
	docker stop form-api-dev
