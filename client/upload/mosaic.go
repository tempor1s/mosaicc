package upload

// MosaicUploader is the structure that will allow us to upload images to mosaic
type MosaicUploader struct{}

// Upload will upload the given contents to mosaic
func (m *MosaicUploader) Upload(contents []byte) (string, error) {
	return "", nil
}
