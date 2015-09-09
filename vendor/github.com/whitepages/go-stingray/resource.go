package stingray

// Resourcer is the interface implemented by objects that represent
// Stingray configuration resources.
type Resourcer interface {
	Name() string
	setName(string)
	endpoint() string
	contentType() string
	decode([]byte) error
	String() string
}

type resource struct {
	name        string
	contentType string
}

func (r *resource) Name() string {
	return r.name
}

func (r *resource) setName(name string) {
	r.name = name
}
