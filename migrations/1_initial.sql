-- +migrate Up

CREATE TABLE "form_schema" (
  "uuid" UUID NOT NULL PRIMARY KEY,
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
