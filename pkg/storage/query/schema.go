package query

// CreateSchema TODO
type CreateSchema struct {
	Language   string
	Title      string
	Definition string
}

// ReadSchema TODO
type ReadSchema struct {
	ID string
}

// UpdateSchema TODO
type UpdateSchema struct {
	ID         string
	Language   string
	Title      string
	Definition string
}

// DeleteSchema TODO
type DeleteSchema struct {
	ID          string
	Permanently bool
}
