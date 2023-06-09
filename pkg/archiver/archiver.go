package archiver

import (
	"github.com/mholt/archiver/v3"
)

func Compress(sources []string, destination string) error {
	return archiver.Archive(sources, destination)
}

func Decompress(source string, destination string) error {
	return archiver.Unarchive(source, destination)
}
