-- +migrate Up

CREATE TABLE "log" (
  "id"          BIGSERIAL PRIMARY KEY,
  "account_id"  UUID      NOT NULL,
  "schema_id"   UUID      NOT NULL,
  "input_id"    UUID      NOT NULL,
  "template_id" UUID      NOT NULL,
  "identifier"  UUID      NOT NULL,
  "code"        INTEGER   NOT NULL,
  "context"     JSONB     NOT NULL,
  "created_at"  TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TRIGGER "immutable_log"
  BEFORE UPDATE
  ON "log"
  FOR EACH ROW EXECUTE PROCEDURE ignore_update();



-- +migrate Down

DROP TRIGGER "immutable_log"
ON "log";

DROP TABLE "log";
