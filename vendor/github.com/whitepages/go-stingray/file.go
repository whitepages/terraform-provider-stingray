package stingray

// fileResource represents a file resource.
type fileResource struct {
	resource
	Content []byte
}

func (f *fileResource) String() string {
	return string(f.Content)
}

func (f *fileResource) decode(data []byte) error {
	f.Content = data

	return nil
}

func (f *fileResource) contentType() string {
	return "application/octet-stream"
}
