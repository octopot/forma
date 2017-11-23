-- +migrate Up

CREATE EXTENSION "uuid-ossp";

-- +migrate StatementBegin
CREATE FUNCTION update_timestamp() RETURNS TRIGGER AS $update_timestamp$
BEGIN
  NEW.updated_at := current_timestamp;
  RETURN NEW;
-- [42883] ERROR: could not identify an equality operator for type xml
--   IF NEW.* IS DISTINCT FROM OLD.* THEN
--     NEW.updated_at := current_timestamp;
--     RETURN NEW;
--   ELSE
--     RETURN OLD;
--   END IF;
END;
$update_timestamp$ LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TYPE STATUS AS ENUM ('enabled','disabled');

CREATE TABLE "form_schema" (
  "uuid" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user" UUID NOT NULL,
  "schema" XML NOT NULL,
  "status" STATUS NOT NULL DEFAULT 'enabled',
  "created_at" TIMESTAMP NOT NULL DEFAULT now(),
  "updated_at" TIMESTAMP
);

CREATE TRIGGER "form_data_updated" BEFORE UPDATE ON "form_schema" FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TABLE "form_data" (
  "id" SERIAL PRIMARY KEY,
  "uuid" UUID NOT NULL,
  "data" JSONB NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down

DROP TABLE "form_data";

DROP TRIGGER "form_data_updated" ON "form_schema";

DROP TABLE "form_schema";

DROP TYPE STATUS;

DROP FUNCTION update_timestamp();

DROP EXTENSION "uuid-ossp";
