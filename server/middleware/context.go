package middleware

type (
	// EncoderKey is used as a context key to store a form schema encoder.
	EncoderKey struct{}
	// SchemaKey is used as a context key to store a form schema UUID.
	SchemaKey struct{}
	// TemplateKey is used as a context key to store a template UUID.
	TemplateKey struct{}
)
