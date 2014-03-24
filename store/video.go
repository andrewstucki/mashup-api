package store

import (
	"io"
	"regexp"
	"errors"
	"log"
	"bytes"
)

const (
	VIDEO_TYPES = "video/(mov|mp4|quicktime|x-m4v)"
)

var (
	defaultVideoConfig = config{
		minFileSize:        1,       // bytes
		maxFileSize:        50000000, // bytes
		acceptFileTypes:    regexp.MustCompile(VIDEO_TYPES),
	}
	videoRegex        = regexp.MustCompile(VIDEO_TYPES)
)

func createVideo(fileInfo *FileInfo, reader io.Reader) (err error) {
  var buf bytes.Buffer
  limitedReader := &io.LimitedReader{R: reader, N: int64(defaultVideoConfig.maxFileSize + 1)}
  
  _, err = io.Copy(&buf, limitedReader)
	size := buf.Len()
	if size < defaultVideoConfig.minFileSize {
		log.Println("File failed validation: too small.", size, defaultVideoConfig.minFileSize)
		fileInfo.Error = "minFileSize"
		err = errors.New("Too Small")
		return
	} else if size > defaultVideoConfig.maxFileSize {
		log.Println("File failed validation: too large.", size, defaultVideoConfig.maxFileSize)
		fileInfo.Error = "maxFileSize"
    err = errors.New("Too Big")
		return
	}
	fileInfo.Size = int64(size)
	log.Println("yay!!! video!!!")
	return
}

func isStoreableVideo(fileInfo *FileInfo) bool {
  return defaultVideoConfig.acceptFileTypes.MatchString(fileInfo.Type) && videoRegex.MatchString(fileInfo.Type) && (fileInfo.Backend == "vimeo")
}