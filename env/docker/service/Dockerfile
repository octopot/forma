FROM golang:latest AS build

WORKDIR /src

ADD . .

ENV GOOS        linux
ENV GOARCH      amd64
ENV CGO_ENABLED 0
ENV GO111MODULE on

RUN go env \
 && export _commit="-X main.commit=$(git rev-parse --short HEAD || echo 'none')" \
 && export _date="-X main.date=$(date -u +%FT%X%Z || echo 'unknown')" \
 && export _version="-X main.version=$(git describe --tags 2>&- || echo 'dev' | cut -d - -f 1)" \
 && go build -o /go/bin/service \
             -ldflags "-s -w ${_commit} ${_date} ${_version}" -mod vendor .


FROM alpine:latest AS service

LABEL maintainer="Kamil Samigullin <kamil@samigullin.info>"

RUN adduser -D -H -u 1000 service
USER service

COPY --from=build /go/bin/service /usr/local/bin/

ENV BIND=0.0.0.0 \
    HTTP_PORT=8080 \
    PROFILING_PORT=8090 \
    MONITORING_PORT=8091 \
    GRPC_PORT=8092 \
    GRPC_GATEWAY_PORT=8093

EXPOSE ${HTTP_PORT} \
       ${PROFILING_PORT} \
       ${MONITORING_PORT} \
       ${GRPC_PORT} \
       ${GRPC_GATEWAY_PORT}

ENTRYPOINT [ "service" ]
CMD        [ "help", "run" ]
