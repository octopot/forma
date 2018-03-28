package middleware

type (
	// EncoderKey is used as a context key to store a Schema encoder.
	EncoderKey struct{}
	// SchemaKey is used as a context key to store a Schema ID.
	SchemaKey struct{}
	// TemplateKey is used as a context key to store a Template ID.
	TemplateKey struct{}
)
