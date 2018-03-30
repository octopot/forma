FROM golang:1.10 AS build

ARG QUICK

WORKDIR /go/src/github.com/kamilsk/form-api

ADD . .

ENV GOOS        linux
ENV GOARCH      amd64
ENV CGO_ENABLED 0

RUN echo "os: ${GOOS}", "arch: ${GOARCH}", "cgo: ${CGO_ENABLED}" \
 && go install -ldflags '-s -w -X main.version=dev -X main.commit=dev -X main.date=dev' .

# ~~~

FROM alpine:latest AS service

LABEL maintainer="Kamil Samigullin <kamil@samigullin.info>"

COPY --from=build /go/bin/form-api /usr/local/bin/

ENV BIND            "0.0.0.0"
ENV PORT            "80"
ENV PROFILE_PORT    "8090"
ENV MONITORING_PORT "8091"

EXPOSE ${PORT}
EXPOSE ${PROFILE_PORT}
EXPOSE ${MONITORING_PORT}

ENTRYPOINT [ "form-api" ]
CMD        [ "run" ]
