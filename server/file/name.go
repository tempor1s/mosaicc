package file

import (
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateName generates a name of the specified length.
func GenerateName(strLen int) string {
	b := make([]byte, strLen)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// GetFileBasename returns the base file name of a given filename.
// Eg. the file name without the extension.
func GetFileBasename(filename string) (basename string) {
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return filename
	}

	return filename[:index]
}
