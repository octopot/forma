-- +migrate Up

CREATE TABLE "schema" (
  "id"         UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "account_id" UUID         NOT NULL,
  "title"      VARCHAR(256) NOT NULL,
  "definition" TEXT         NOT NULL,
  "created_at" TIMESTAMP    NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP    NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP    NULL                 DEFAULT NULL,
  UNIQUE ("account_id", "title")
);

CREATE TABLE "template" (
  "id"         UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "account_id" UUID         NOT NULL,
  "title"      VARCHAR(256) NOT NULL,
  "definition" TEXT         NOT NULL,
  "created_at" TIMESTAMP    NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP    NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP    NULL                 DEFAULT NULL,
  UNIQUE ("account_id", "title")
);

CREATE TABLE "input" (
  "id"         UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "schema_id"  UUID      NOT NULL,
  "data"       JSONB     NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now()
);

CREATE TRIGGER "schema_updated"
  BEFORE UPDATE
  ON "schema"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER "template_updated"
  BEFORE UPDATE
  ON "template"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER "immutable_input"
  BEFORE UPDATE
  ON "input"
  FOR EACH ROW EXECUTE PROCEDURE ignore_update();



-- +migrate Down

DROP TRIGGER "immutable_input"
ON "input";

DROP TRIGGER "template_updated"
ON "template";

DROP TRIGGER "schema_updated"
ON "schema";

DROP TABLE "input";

DROP TABLE "template";

DROP TABLE "schema";
