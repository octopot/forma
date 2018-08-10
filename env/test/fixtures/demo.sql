DO $$
DECLARE   demoAccount  "account"."id"%TYPE := '10000000-2000-4000-8000-160000000001';
  DECLARE demoUser     "user"."id"%TYPE := '10000000-2000-4000-8000-160000000002';
  DECLARE demoToken    "token"."id"%TYPE := '10000000-2000-4000-8000-160000000003';
  DECLARE subscribe    "schema"."id"%TYPE := '10000000-2000-4000-8000-160000000004';
  DECLARE feedbackEn   "schema"."id"%TYPE := '10000000-2000-4000-8000-160000000005';
  DECLARE feedbackRu   "schema"."id"%TYPE := '10000000-2000-4000-8000-160000000006';
  DECLARE subscribeTpl "template"."id"%TYPE := '10000000-2000-4000-8000-160000000007';
  DECLARE feedbackTpl  "template"."id"%TYPE := '10000000-2000-4000-8000-160000000008';
BEGIN
  DELETE FROM "log"
  WHERE "account_id" = demoAccount
        OR "schema_id" IN (subscribe, feedbackEn, feedbackRu)
        OR "template_id" IN (subscribeTpl, feedbackTpl);

  DELETE FROM "input"
  WHERE "schema_id" IN (subscribe, feedbackEn, feedbackRu);

  DELETE FROM "template"
  WHERE "account_id" = demoAccount OR "id" IN (subscribeTpl, feedbackTpl);

  DELETE FROM "schema"
  WHERE "account_id" = demoAccount OR "id" IN (subscribe, feedbackEn, feedbackRu);

  DELETE FROM "token"
  WHERE "id" = demoToken OR "user_id" = demoUser;

  DELETE FROM "user"
  WHERE "id" = demoUser OR "account_id" = demoAccount;

  DELETE FROM "account"
  WHERE "id" = demoAccount;

  INSERT INTO "account" ("id", "name")
  VALUES (demoAccount, 'Demo account');

  INSERT INTO "user" ("id", "account_id", "name")
  VALUES (demoUser, demoAccount, 'Demo user');

  INSERT INTO "token" ("id", "user_id", "expired_at")
  VALUES (demoToken, demoUser, NULL);

  INSERT INTO "schema" ("id", "account_id", "language", "title", "definition")
  VALUES
    (subscribe, demoAccount, 'en', 'Email subscription', '
    <form action="https://kamil.samigullin.info/" enctype="application/x-www-form-urlencoded" method="post">
        <input name="email" type="email" title="Email" maxlength="64" required="1"/>
    </form>'),
    (feedbackEn, demoAccount, 'en', 'GitHub demo page', '
    <form action="https://kamilsk.github.io/form-api/" enctype="application/x-www-form-urlencoded" method="post">
        <input name="name" type="text" title="Name" placeholder="Name..." maxlength="25" required="1"/>
        <input name="feedback" type="text" title="Feedback" placeholder="Your feedback..." maxlength="255"
               required="1"/>
    </form>'),
    (feedbackRu, demoAccount, 'ru', 'GitHub демо', '
    <form action="https://kamilsk.github.io/form-api/" enctype="application/x-www-form-urlencoded" method="post">
        <input name="name" type="text" title="Имя" placeholder="Имя..." maxlength="25" required="1"/>
        <input name="feedback" type="text" title="Комментарий" placeholder="Ваш комментарий..." maxlength="255"
               required="1"/>
    </form>');

  INSERT INTO "template" ("id", "account_id", "title", "definition")
  VALUES
    (subscribeTpl, demoAccount, 'Subscribe template', '{{- define "forma.body" -}}
    <div class="row">
        {{- with .Schema.Input "email" -}}
            <div class="col-md-8">
                <span class="bmd-form-group">
                    <div class="input-group">
                        <div class="input-group-prepend">
                            <span class="input-group-text">
                                <i class="material-icons">mail</i>
                            </span>
                        </div>
                        {{- template "forma.input" . -}}
                    </div>
                </span>
            </div>
        {{- end -}}
    </div>
{{- end -}}
{{- define "forma.submit" -}}
    {{- template "forma.powered_by" . -}}
    <div class="col-md-4">
        <button class="btn btn-primary btn-block" type="submit">Subscribe</button>
    </div>
{{- end -}}'),
    (feedbackTpl, demoAccount, 'Feedback template', '{{- define "forma.body" -}}
    {{- with .Schema.Input "name" -}}
        <div class="form-group row">
            <label for="{{ .ID }}"
                   class="col-sm-3 col-form-label">{{ .Name }}</label>
            <div class="col-sm-9">
                {{- template "forma.input" . -}}
            </div>
        </div>
    {{- end -}}
    {{- with .Schema.Input "feedback" -}}
        <div class="form-group row">
            <label for="{{ .ID }}"
                   class="col-sm-3 col-form-label">{{ .Name }}</label>
            <div class="col-sm-9">
                {{- template "forma.input" . -}}
            </div>
        </div>
    {{- end -}}
{{- end -}}
{{- define "forma.submit" -}}
    <input name="_redirect" type="hidden" value="https://kamilsk.github.io/form-api/">
    <input name="_timeout" type="hidden" value="60">
    <input class="btn btn-dark" type="submit">
{{- end -}}');

  INSERT INTO "input" ("schema_id", "data")
  VALUES
    (subscribe, '{
      "email": ["test@my.email"]
    }' :: JSONB),
    (feedbackEn, '{
      "name": ["C. Northcote Parkinson"],
      "feedback": ["Work contracts to fit in the time we give it."]
    }' :: JSONB),
    (feedbackRu, '{
      "name": ["Сирил Норткот Паркинсон"],
      "feedback": ["Работа заполняет время, отпущенное на неё."]
    }' :: JSONB);
END;
$$;
