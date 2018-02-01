-- +migrate Up

-- +migrate StatementBegin
DO $$
  DECLARE demoUser UUID := uuid_generate_v4();
  DECLARE email    UUID := '41ca5e09-3ce2-4094-b108-3ecc257c6fa4';
  DECLARE githubEn UUID := '41059676-6cb6-42bc-9fb5-8df1699b6881';
  DECLARE githubRu UUID := '41d33b3d-26fe-435c-9402-c8e4a1f87cc1';
  BEGIN
    INSERT INTO "form_schema" ("uuid","user","schema") VALUES
      (email,demoUser,'
<form lang="en" title="Email subscription" action="https://kamil.samigullin.info/" method="post" enctype="application/x-www-form-urlencoded">
  <input name="email" type="email" title="Email" maxlength="64" required="1"/>
</form>
'),
      (githubEn,demoUser,'
<form lang="en" title="GitHub demo page" action="https://kamilsk.github.io/form-api/" method="post" enctype="application/x-www-form-urlencoded">
  <input name="name" type="text" title="Name" placeholder="Name..." maxlength="25" required="1"/>
  <input name="feedback" type="text" title="Feedback" placeholder="Your feedback..." maxlength="255" required="1"/>
</form>
'),
      (githubRu,demoUser,'
<form lang="ru" title="GitHub демо" action="https://kamilsk.github.io/form-api/" method="post" enctype="application/x-www-form-urlencoded">
  <input name="name" type="text" title="Имя" placeholder="Имя..." maxlength="25" required="1"/>
  <input name="feedback" type="text" title="Комментарий" placeholder="Ваш комментарий..." maxlength="255" required="1"/>
</form>
');
    INSERT INTO "form_data" ("uuid","data") VALUES
      (email,'{"email":["test@my.email"]}'),
      (githubEn,'{"name":["C. Northcote Parkinson"],"feedback":["Work contracts to fit in the time we give it."]}'),
      (githubRu,'{"name":["Сирил Норткот Паркинсон"],"feedback":["Работа заполняет время, отпущенное на неё."]}');
  END;
$$;
-- +migrate StatementEnd

-- +migrate Down

-- +migrate StatementBegin
DO $$
  DECLARE email    UUID := '41ca5e09-3ce2-4094-b108-3ecc257c6fa4';
  DECLARE githubEn UUID := '41059676-6cb6-42bc-9fb5-8df1699b6881';
  DECLARE githubRu UUID := '41d33b3d-26fe-435c-9402-c8e4a1f87cc1';
  BEGIN
    DELETE FROM "form_data" WHERE "uuid" IN (email,githubEn,githubRu);
    DELETE FROM "form_schema" WHERE "uuid" IN (email,githubEn,githubRu);
  END;
$$;
-- +migrate StatementEnd
