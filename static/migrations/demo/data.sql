-- +migrate Up

-- +migrate StatementBegin
BEGIN;
  DO $$
    DECLARE demoUser UUID := uuid_generate_v4();
    DECLARE email    UUID := '41ca5e09-3ce2-4094-b108-3ecc257c6fa4';
    DECLARE github   UUID := '41059676-6cb6-42bc-9fb5-8df1699b6881';
    BEGIN
      INSERT INTO "form_schema" ("uuid","user","schema") VALUES
        (email,demoUser,'
<form title="Email subscription" method="post" enctype="application/x-www-form-urlencoded">
    <input name="email" type="email" title="Email" maxlength="64" required="1"/>
    <input name="_redirect" type="hidden" value="https://kamil.samigullin.info/"/>
</form>
'),
        (github,demoUser,'
<form title="GitHub demo page" method="post" enctype="application/x-www-form-urlencoded">
    <input name="name" type="text" title="Name" placeholder="Name..." maxlength="25" required="1"/>
    <input name="feedback" type="text" title="Feedback" placeholder="Your feedback..." maxlength="255" required="1"/>
    <input name="_redirect" type="hidden" value="http://kamilsk.github.io/form-api/"/>
</form>
');
      INSERT INTO "form_data" ("uuid","data") VALUES
        (email,'{"email":["test@my.email"]}'),
        (github,'{"name":["C. Northcote Parkinson"],"feedback":["Work contracts to fit in the time we give it."]}');
    END;
  $$;
COMMIT;
-- +migrate StatementEnd

-- +migrate Down

-- +migrate StatementBegin
BEGIN;
  DO $$
    DECLARE email  UUID := '41ca5e09-3ce2-4094-b108-3ecc257c6fa4';
    DECLARE github UUID := '41059676-6cb6-42bc-9fb5-8df1699b6881';
    BEGIN
      DELETE FROM "form_data" WHERE "uuid" IN (email, github);
      DELETE FROM "form_schema" WHERE "uuid" IN (email, github);
    END;
  $$;
COMMIT;
-- +migrate StatementEnd
