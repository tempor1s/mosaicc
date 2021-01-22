package upload

// Uploader is an interface that different types can implement that will allow an image to be uploaded
type Uploader interface {
	Upload([]byte) (string, error)
}
