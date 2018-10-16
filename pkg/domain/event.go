package domain

// Option contains rules for request processing.
type Option struct {
	// Anonymously: use zero-identifier instead of origin.
	Anonymously bool
	// Debug: return debug information to the client.
	Debug bool
	// NoLog: do not log link navigation.
	NoLog bool
}

// InputEvent TODO issue#173
//go:generate easyjson -all
type InputEvent struct {
	SchemaID   ID           `json:"schema_id"`
	InputID    ID           `json:"input_id"`
	TemplateID *ID          `json:"template_id,omitempty"`
	Identifier *ID          `json:"identifier,omitempty"`
	Context    InputContext `json:"context"`
	Code       int          `json:"code"`
	URL        string       `json:"url"`
}

// Redirect TODO issue#173
func (event InputEvent) Redirect() string {
	if event.URL != "" {
		return event.URL
	}
	return event.Context.Referer()
}

// InputContext contains context information about an input event.
//go:generate easyjson -all
type InputContext struct {
	Cookies map[string]string   `json:"cookies,omitempty"`
	Headers map[string][]string `json:"headers,omitempty"`
	Queries map[string][]string `json:"queries,omitempty"`
}
