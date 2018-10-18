package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	repository "github.com/kamilsk/form-api/pkg/storage/types"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/static"
	"github.com/kamilsk/form-api/pkg/transfer/api/v1"
)

// New returns a new instance of the Forma service.
// It can raise the panic if base URL is invalid or HTML templates are not available.
func New(cnf config.ServiceConfig, storage Storage, tracker Tracker) *Forma {
	u, err := url.Parse(cnf.BaseURL)
	if err != nil {
		panic(err)
	}
	return &Forma{config: cnf, storage: storage, tracker: tracker,
		baseURL: u, templates: struct {
			errorTpl    *template.Template
			redirectTpl *template.Template
		}{
			errorTpl:    template.Must(template.New("error").Parse(must(cnf.TemplateDir, "error.html"))),
			redirectTpl: template.Must(template.New("redirect").Parse(must(cnf.TemplateDir, "redirect.html"))),
		}}
}

// Forma is the primary application service.
type Forma struct {
	config  config.ServiceConfig
	storage Storage
	tracker Tracker

	baseURL   *url.URL
	templates struct {
		errorTpl    *template.Template
		redirectTpl *template.Template
	}
}

// HandleGetV1 handles an input request.
// Deprecated: TODO issue#version3.0 use SchemaEditor and gRPC gateway instead
func (service *Forma) HandleGetV1(ctx context.Context, req v1.GetRequest) (resp v1.GetResponse) {
	resp.Schema, resp.Error = service.storage.Schema(ctx, req.ID)
	if resp.Error != nil {
		return
	}
	enrich(service.baseURL, &resp.Schema)
	return
}

// HandleInput handles an input request.
func (service *Forma) HandleInput(ctx context.Context, req v1.PostRequest) (resp v1.PostResponse) {
	schema, readErr := service.storage.Schema(ctx, req.ID)
	if readErr != nil {
		resp.Error = readErr
		return
	}

	enrich(service.baseURL, &schema)
	resp.URL = req.InputData.Redirect(req.Context.Referer(), schema.Action)

	verified, validateErr := schema.Apply(req.InputData)
	if validateErr != nil {
		service.templates.errorTpl.Execute(req.Output, static.ErrorPageContext{
			Schema:   schema,
			Error:    validateErr,
			Delay:    30 * time.Duration(len(validateErr.InputWithErrors())) * time.Second,
			Redirect: resp.URL,
		})
		resp.Error = errors.Validation(errors.InvalidFormDataMessage, validateErr,
			"trying to apply data to the schema %q", req.ID)
		return
	}

	var input *repository.Input
	input, resp.Error = service.storage.StoreInput(ctx, req.ID, verified)

	if resp.Error == nil && !req.Context.Option().NoLog && !ignore(req.Context) {
		// if option.Anonymously {}
		event := domain.InputEvent{
			SchemaID:   req.ID,
			InputID:    input.ID,
			TemplateID: req.InputData.Template(),
			Identifier: req.Context.Identifier(),
			Context:    req.Context,
			Code:       http.StatusFound, // TODO issue#design
			URL:        resp.URL,
		}
		resp.Error = service.tracker.LogInput(ctx, event) // TODO issue#109
	}

	if u, err := url.Parse(resp.URL); err == nil {
		type feedback struct {
			ID       domain.ID `json:"input"`
			SchemaID domain.ID `json:"id"`
			Result   string    `json:"result"`
		}
		if resp.Error != nil {
			u.Fragment = base64.StdEncoding.EncodeToString(func() []byte {
				raw, _ := json.Marshal(feedback{SchemaID: req.ID, Result: "failure"})
				return raw
			}())
		} else {
			u.Fragment = base64.StdEncoding.EncodeToString(func() []byte {
				raw, _ := json.Marshal(feedback{ID: input.ID, SchemaID: req.ID, Result: "success"})
				return raw
			}())
		}
		resp.URL = u.String()
	}

	return resp
}

// configure form action
func enrich(base *url.URL, schema *domain.Schema) {
	schema.Action = extend(*base, "api/v1", schema.ID)
	schema.Method = http.MethodPost
	schema.EncodingType = "application/x-www-form-urlencoded"
}

func extend(base url.URL, paths ...string) string {
	if len(paths) == 0 {
		return base.String()
	}
	base.Path = path.Join(append([]string{base.Path}, paths...)...)
	return base.String()
}

func ignore(req domain.InputContext) bool {
	// do not log curl request
	return strings.HasPrefix(req.UserAgent(), "curl/")
}

func must(base, tpl string) string {
	b, err := static.LoadTemplate(base, tpl)
	if err != nil {
		panic(tpl)
	}
	return string(b)
}
