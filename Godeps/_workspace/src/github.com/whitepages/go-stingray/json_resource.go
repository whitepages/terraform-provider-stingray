package stingray

// jsonResource represents a JSON resource.
type jsonResource struct {
	resource
}

func (f *jsonResource) contentType() string {
	return "application/json"
}
