package domain

import (
	"net/url"
	"time"
)

const (
	// EmailType specifies `<input type="email">`.
	EmailType = "email"
	// HiddenType specifies `<input type="hidden">`
	HiddenType = "hidden"
	// TextType specifies `<input type="text">`.
	TextType = "text"

	// TODO issue#174
	captchaType   = "captcha"
	reCAPTCHAType = "reCAPTCHA"
)

const (
	redirectKey = "_redirect"
	resourceKey = "_resource"
	templateKey = "_template"
	timeoutKey  = "_timeout"
)

// Input represents an element of an HTML form.
//go:generate easyjson -all
type Input struct {
	ID          string `json:"id,omitempty"          yaml:"id,omitempty"          xml:"id,attr,omitempty"`
	Name        string `json:"name"                  yaml:"name"                  xml:"name,attr"`
	Type        string `json:"type"                  yaml:"type"                  xml:"type,attr"`
	Title       string `json:"title,omitempty"       yaml:"title,omitempty"       xml:"title,attr,omitempty"`
	Placeholder string `json:"placeholder,omitempty" yaml:"placeholder,omitempty" xml:"placeholder,attr,omitempty"`
	Value       string `json:"value,omitempty"       yaml:"value,omitempty"       xml:"value,attr,omitempty"`
	MinLength   int    `json:"minlength,omitempty"   yaml:"minlength,omitempty"   xml:"minlength,attr,omitempty"`
	MaxLength   int    `json:"maxlength,omitempty"   yaml:"maxlength,omitempty"   xml:"maxlength,attr,omitempty"`
	Required    bool   `json:"required,omitempty"    yaml:"required,omitempty"    xml:"required,attr,omitempty"`
	Strict      bool   `json:"strict,omitempty"      yaml:"strict,omitempty"      xml:"strict,attr,omitempty"`
}

// InputData TODO issue#173
type InputData url.Values

// Redirect TODO issue#173
func (d InputData) Redirect(fallback ...string) string {
	value := url.Values(d).Get(redirectKey)
	if value == "" {
		for _, value = range fallback {
			if value != "" {
				break
			}
		}
	}
	return value
}

// Resource TODO issue#173
func (d InputData) Resource() ID {
	return ID(url.Values(d).Get(resourceKey))
}

// Template TODO issue#173
func (d InputData) Template() *ID {
	var id = ID(url.Values(d).Get(templateKey))
	if id.IsValid() {
		return &id
	}
	return nil
}

// Timeout TODO issue#173
func (d InputData) Timeout() time.Duration {
	timeout, err := time.ParseDuration(url.Values(d).Get(timeoutKey))
	if err != nil {
		return 30 * time.Second
	}
	return timeout
}
