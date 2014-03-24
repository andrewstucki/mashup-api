package store

import (
	"io"
	"regexp"
	"errors"
	// Register image handling libraries by importing them.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"bytes"
)

const (
	IMAGE_TYPES = "image/(gif|p?jpeg|(x-)?png)"
)

var (
	defaultImageConfig = config{
		minFileSize:        1,       // bytes
		maxFileSize:        5000000, // bytes
		acceptFileTypes:    regexp.MustCompile(IMAGE_TYPES),
	}
	imageRegex        = regexp.MustCompile(IMAGE_TYPES)
)

func createImage(fileInfo *FileInfo, reader io.Reader) (err error) {
  var buf bytes.Buffer
  limitedReader := &io.LimitedReader{R: reader, N: int64(defaultImageConfig.maxFileSize + 1)}
  
  _, err = io.Copy(&buf, limitedReader)
	size := buf.Len()
	if size < defaultImageConfig.minFileSize {
		log.Println("File failed validation: too small.", size, defaultImageConfig.minFileSize)
		fileInfo.Error = "minFileSize"
		err = errors.New("Too Small")
		return
	} else if size > defaultImageConfig.maxFileSize {
		log.Println("File failed validation: too large.", size, defaultImageConfig.maxFileSize)
		fileInfo.Error = "maxFileSize"
    err = errors.New("Too Big")
		return
	}
	fileInfo.Size = int64(size)
	log.Println("yay!!!")
	return
}

func isStoreableImage(fileInfo *FileInfo) bool {
  return defaultImageConfig.acceptFileTypes.MatchString(fileInfo.Type) && imageRegex.MatchString(fileInfo.Type) && (fileInfo.Backend == "flickr")
}