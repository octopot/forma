#!/usr/bin/env bash

export FORMA_TOKEN=10000000-2000-4000-8000-160000000003

form-api ctl create -f env/client/grpc/schema.create.yml
form-api ctl create -f env/client/grpc/template.create.yml

form-api ctl read -f env/client/grpc/schema.read.yml
form-api ctl read -f env/client/grpc/template.read.yml

form-api ctl update -f env/client/grpc/schema.update.yml
form-api ctl update -f env/client/grpc/template.update.yml

form-api ctl delete -f env/client/grpc/schema.delete.yml
form-api ctl delete -f env/client/grpc/template.delete.yml
