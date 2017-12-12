-- +migrate Up

INSERT INTO "form_schema" ("uuid","user","schema")
VALUES ('41ca5e09-3ce2-4094-b108-3ecc257c6fa4',uuid_generate_v4(),'
<form title="Email subscription" method="post" enctype="application/x-www-form-urlencoded">
    <input name="email" type="email" title="Email" maxlength="64" required="1"/>
    <input name="_redirect" type="hidden" value="https://kamil.samigullin.info/"/>
</form>
');

INSERT INTO "form_data" ("uuid","data") VALUES ('41ca5e09-3ce2-4094-b108-3ecc257c6fa4','{"email":"kamil@samigullin.info"}');

-- +migrate Down

DELETE FROM "form_data" WHERE "uuid" = '41ca5e09-3ce2-4094-b108-3ecc257c6fa4';

DELETE FROM "form_schema" WHERE "uuid" = '41ca5e09-3ce2-4094-b108-3ecc257c6fa4';
