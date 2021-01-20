package file

import (
	"mime/multipart"
	"path/filepath"

	"github.com/tempor1s/mosaic/models"
	"go.mongodb.org/mongo-driver/mongo"
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
func (f *File) GenerateName(db *mongo.Database, len int) *File {
	// we will attempt to generate the name until the name we generated is not in the database

	store := models.NewStore(db)

	// first, generate an initial name
	name := GenerateName(len)
	_, exists, _ := store.GetImageByID(name)

	// while the name still exists (not a new unique name yet)
	for exists == true {
		// generate a new name and run the check again
		name = GenerateName(len)
		_, exists, _ = store.GetImageByID(name)
	}

	// we found a unique name, so give that to the file
	f.GiveName(name)

	return f
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
