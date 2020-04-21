-- +migrate Up

CREATE TABLE "account" (
  "id"         UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name"       VARCHAR(128) NOT NULL,
  "created_at" TIMESTAMP    NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP    NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP    NULL                 DEFAULT NULL
);

CREATE TABLE "user" (
  "id"         UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "account_id" UUID         NOT NULL,
  "name"       VARCHAR(128) NOT NULL,
  "created_at" TIMESTAMP    NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP    NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP    NULL                 DEFAULT NULL
);

CREATE TABLE "token" (
  "id"         UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id"    UUID      NOT NULL,
  "expired_at" TIMESTAMP NULL                 DEFAULT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now()
);

CREATE TRIGGER "account_updated"
  BEFORE UPDATE
  ON "account"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER "user_updated"
  BEFORE UPDATE
  ON "user"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER "immutable_token"
  BEFORE UPDATE
  ON "token"
  FOR EACH ROW EXECUTE PROCEDURE ignore_update();



-- +migrate Down

DROP TRIGGER "immutable_token"
ON "token";

DROP TRIGGER "user_updated"
ON "user";

DROP TRIGGER "account_updated"
ON "account";

DROP TABLE "token";

DROP TABLE "user";

DROP TABLE "account";
