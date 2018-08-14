package query

// CreateTemplate TODO
type CreateTemplate struct {
	Title      string
	Definition string
}

// ReadTemplate TODO
type ReadTemplate struct {
	ID string
}

// UpdateTemplate TODO
type UpdateTemplate struct {
	ID         string
	Title      string
	Definition string
}

// DeleteTemplate TODO
type DeleteTemplate struct {
	ID          string
	Permanently bool
}
