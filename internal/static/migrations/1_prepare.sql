-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +migrate StatementBegin
CREATE FUNCTION update_timestamp()
  RETURNS TRIGGER AS $update_timestamp$
BEGIN
  IF NEW.* IS DISTINCT FROM OLD.*
  THEN
    NEW.updated_at := current_timestamp;
    RETURN NEW;
  ELSE
    RETURN OLD;
  END IF;
END;
$update_timestamp$
LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE FUNCTION ignore_update()
  RETURNS TRIGGER AS $ignore_update$
BEGIN
  RETURN OLD;
END;
$ignore_update$
LANGUAGE plpgsql;
-- +migrate StatementEnd



-- +migrate Down

DROP FUNCTION ignore_update();

DROP FUNCTION update_timestamp();
