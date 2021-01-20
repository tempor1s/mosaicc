package file

import (
	"mime/multipart"
	"path/filepath"
)

// File holds a file.
type File struct {
	File      multipart.File        `json:"-"`
	Header    *multipart.FileHeader `json:"-"`
	Basename  string                `json:"basename,omitempty"` // Without extension
	Fullname  string                `json:"fullname,omitempty"` // With extension
	Extension string                `json:"extension,omitempty"`
	MIMEType  string                `json:"mime_type,omitempty"`
	Size      int                   `json:"size,omitempty"`
}

// NewFile will create a new file from a multipart.FileHeader.
func NewFile(file multipart.File, fileHeader *multipart.FileHeader) *File {
	return &File{
		File:      file,
		Header:    fileHeader,
		Basename:  GetFileBasename(fileHeader.Filename),
		Fullname:  fileHeader.Filename,
		Extension: filepath.Ext(fileHeader.Filename),
		MIMEType:  fileHeader.Header["Content-Type"][0],
		Size:      int(fileHeader.Size),
	}
}

// GenerateName will generate a new name with a given length.
func (f *File) GenerateName(len int) *File {
	return f.GiveName(GenerateName(len))
}

// GiveName will give the File a new name, and update the basename and fullname.
func (f *File) GiveName(name string) *File {
	f.Basename = name
	f.Fullname = name + f.Extension
	return f
}

// Close will properly close the file.
func (f *File) Close() error {
	return f.File.Close()
}
