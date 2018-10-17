IMAGE_VERSION := 3.x
PACKAGE       := github.com/kamilsk/form-api


.PHONY: docker-build
docker-build:
	docker build -f env/Dockerfile \
	             -t kamilsk/form-api:$(IMAGE_VERSION) \
	             -t kamilsk/form-api:latest \
	             -t quay.io/kamilsk/form-api:$(IMAGE_VERSION) \
	             -t quay.io/kamilsk/form-api:latest \
	             --build-arg PACKAGE=$(PACKAGE) \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-push
docker-push:
	docker push kamilsk/form-api:$(IMAGE_VERSION)
	docker push kamilsk/form-api:latest
	docker push quay.io/kamilsk/form-api:$(IMAGE_VERSION)
	docker push quay.io/kamilsk/form-api:latest

.PHONY: docker-refresh
docker-refresh:
	docker images --all \
	| grep '^kamilsk\/form-api\s\+' \
	| awk '{print $$3}' \
	| xargs docker rmi -f &>/dev/null || true
	docker pull kamilsk/form-api:$(IMAGE_VERSION)



.PHONY: publish
publish: docker-build docker-push



.PHONY: docker-start
docker-start:
	docker run --rm -d \
	           --env-file env/.env.example \
	           --name form-api-dev \
	           --publish 8080:80 \
	           --publish 8090:8090 \
	           --publish 8091:8091 \
	           --publish 8092:8092 \
	           kamilsk/form-api:$(IMAGE_VERSION)

.PHONY: docker-logs
docker-logs:
	docker logs -f form-api-dev

.PHONY: docker-stop
docker-stop:
	docker stop form-api-dev
