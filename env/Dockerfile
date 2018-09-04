FROM golang:latest AS build

ARG PACKAGE

WORKDIR /go/src/${PACKAGE}

ADD . .

ENV GOOS        linux
ENV GOARCH      amd64
ENV CGO_ENABLED 0

RUN echo "os: ${GOOS}", "arch: ${GOARCH}", "cgo: ${CGO_ENABLED}" \
 && go build -o /go/bin/service -ldflags '-s -w -X main.version=dev -X main.commit=dev -X main.date=dev' .

# ~~~

FROM alpine:latest AS service

LABEL maintainer="Kamil Samigullin <kamil@samigullin.info>"

COPY --from=build /go/bin/service /usr/local/bin/

ENV BIND            "0.0.0.0"
ENV HTTP_PORT       "80"
ENV PROFILING_PORT  "8090"
ENV MONITORING_PORT "8091"
ENV GRPC_PORT       "8092"

EXPOSE ${HTTP_PORT}
EXPOSE ${PROFILING_PORT}
EXPOSE ${MONITORING_PORT}
EXPOSE ${GRPC_PORT}

ENTRYPOINT [ "service" ]
CMD        [ "run" ]
