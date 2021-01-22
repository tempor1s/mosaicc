package upload

// ImgurUploader is the structure that will allow us to upload images to imgur
type ImgurUploader struct{}

// Upload will upload the given contents to imgur
func (i *ImgurUploader) Upload(contents []byte) (string, error) {
	return "", nil
}
