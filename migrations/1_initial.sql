-- +migrate Up

CREATE EXTENSION "uuid-ossp";

CREATE TABLE "form_schema" (
  "uuid" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user" UUID NOT NULL,
  "schema" XML NOT NULL
);

CREATE TABLE "form_data" (
  "id" SERIAL PRIMARY KEY,
  "uuid" UUID NOT NULL,
  "data" JSONB NOT NULL
);


-- +migrate Down

DROP TABLE "form_data";

DROP TABLE "form_schema";
