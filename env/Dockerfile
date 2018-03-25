FROM golang:1.10 AS build

ARG QUICK

WORKDIR /go/src/github.com/kamilsk/form-api

ADD . .

ENV GOOS        linux
ENV GOARCH      amd64
ENV CGO_ENABLED 0

RUN echo "os: ${GOOS}", "arch: ${GOARCH}", "cgo: ${CGO_ENABLED}" \
 && if [ -z "${QUICK}" ]; \
    then \
      echo "downloading dep..."; \
      wget -q -O dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64; \
      echo "moving dep..."; \
      chmod +x dep && mv dep /go/bin; \
      echo "running dep..."; \
      dep ensure; \
      echo "done"; \
    fi \
 && go install -ldflags '-s -w -X main.version=dev -X main.commit=dev -X main.date=dev' .

# ~~~

FROM alpine:latest AS pkg

LABEL maintainer="Kamil Samigullin <kamil@samigullin.info>"

COPY --from=build /go/bin/form-api /usr/local/bin/

ENV BIND "0.0.0.0"
ENV PORT "80"

ENTRYPOINT [ "form-api" ]
CMD        [ "run" ]
