DO $$
DECLARE   demoUser "form_schema"."user"%TYPE := '00000000-0000-4000-8000-000000000000';
  DECLARE email    "form_data"."uuid"%TYPE := '41ca5e09-3ce2-4094-b108-3ecc257c6fa4';
  DECLARE githubEn "form_data"."uuid"%TYPE := '41059676-6cb6-42bc-9fb5-8df1699b6881';
  DECLARE githubRu "form_data"."uuid"%TYPE := '41d33b3d-26fe-435c-9402-c8e4a1f87cc1';
BEGIN
  DELETE FROM "form_data"
  WHERE "uuid" IN (email, githubEn, githubRu);

  DELETE FROM "form_schema"
  WHERE "uuid" IN (email, githubEn, githubRu);

  INSERT INTO "form_schema" ("uuid", "user", "schema")
  VALUES
    (email, demoUser, '
    <form lang="en" title="Email subscription" action="https://kamil.samigullin.info/" method="post"
          enctype="application/x-www-form-urlencoded">
        <input name="email" type="email" title="Email" maxlength="64" required="1"/>
    </form>
    ' :: XML),
    (githubEn, demoUser, '
    <form lang="en" title="GitHub demo page" action="https://kamilsk.github.io/form-api/" method="post"
          enctype="application/x-www-form-urlencoded">
        <input name="name" type="text" title="Name" placeholder="Name..." maxlength="25" required="1"/>
        <input name="feedback" type="text" title="Feedback" placeholder="Your feedback..." maxlength="255"
               required="1"/>
    </form>
    ' :: XML),
    (githubRu, demoUser, '
    <form lang="ru" title="GitHub демо" action="https://kamilsk.github.io/form-api/" method="post"
          enctype="application/x-www-form-urlencoded">
        <input name="name" type="text" title="Имя" placeholder="Имя..." maxlength="25" required="1"/>
        <input name="feedback" type="text" title="Комментарий" placeholder="Ваш комментарий..." maxlength="255"
               required="1"/>
    </form>
    ' :: XML);

  INSERT INTO "form_data" ("uuid", "data")
  VALUES
    (email, '{
      "email": [
        "test@my.email"
      ]
    }'),
    (githubEn, '{
      "name": [
        "C. Northcote Parkinson"
      ],
      "feedback": [
        "Work contracts to fit in the time we give it."
      ]
    }'),
    (githubRu, '{
      "name": [
        "Сирил Норткот Паркинсон"
      ],
      "feedback": [
        "Работа заполняет время, отпущенное на неё."
      ]
    }');
END;
$$;
