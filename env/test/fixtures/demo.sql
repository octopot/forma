DO $$
DECLARE   demoAccount "account"."id"%TYPE := '10000000-2000-4000-8000-160000000001';
  DECLARE demoUser    "user"."id"%TYPE := '10000000-2000-4000-8000-160000000002';
  DECLARE demoToken   "token"."id"%TYPE := '10000000-2000-4000-8000-160000000003';
  DECLARE form        "schema"."id"%TYPE := '10000000-2000-4000-8000-160000000004';
  DECLARE html        "template"."id"%TYPE := '10000000-2000-4000-8000-160000000005';
BEGIN
  TRUNCATE TABLE "event" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "input" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "template" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "schema" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "token" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "user" RESTART IDENTITY RESTRICT;
  TRUNCATE TABLE "account" RESTART IDENTITY RESTRICT;

  INSERT INTO "account" ("id", "name") VALUES (demoAccount, 'Demo Account');

  INSERT INTO "user" ("id", "account_id", "name") VALUES (demoUser, demoAccount, 'Demo User');

  INSERT INTO "token" ("id", "user_id", "expired_at") VALUES (demoToken, demoUser, NULL);

  INSERT INTO "schema" ("id", "account_id", "title", "definition")
  VALUES (form, demoAccount, 'Email Subscription', '
    <form lang="en" action="https://kamil.samigullin.info/">
        <input name="email" type="email" title="Email" maxlength="64" required="1"/>
    </form>');

  INSERT INTO "template" ("id", "account_id", "title", "definition")
  VALUES (html, demoAccount, 'Email Subscription template', '{{- define "forma.body" -}}
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
{{- end -}}');
END;
$$;
