package handlers

// Handlers is the structure that all handlers hang off of. This is so they are able to access dependencies like the DB
type Handlers struct {
}

// New will return a new handlers structure
func New() *Handlers {
	return &Handlers{}
}
