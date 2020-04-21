-- +migrate Up

CREATE TABLE "event" (
  "id"          BIGSERIAL PRIMARY KEY,
  "account_id"  UUID          NOT NULL,
  "schema_id"   UUID          NOT NULL,
  "input_id"    UUID          NOT NULL,
  "template_id" UUID          NULL     DEFAULT NULL,
  "identifier"  UUID          NULL     DEFAULT NULL,
  "context"     JSONB         NOT NULL,
  "code"        INTEGER       NOT NULL,
  "url"         VARCHAR(1024) NOT NULL,
  "created_at"  TIMESTAMP     NOT NULL DEFAULT now()
);

CREATE TRIGGER "immutable_event"
  BEFORE UPDATE
  ON "event"
  FOR EACH ROW EXECUTE PROCEDURE ignore_update();



-- +migrate Down

DROP TRIGGER "immutable_event"
ON "event";

DROP TABLE "event";
