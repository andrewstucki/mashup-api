package store

import (
	"io"
	"errors"
	"log"
	"bytes"
)

var (
	defaultFileConfig = config{
		minFileSize:        1,       // bytes
		maxFileSize:        50000000, // bytes
	}
)

func createFile(fileInfo *FileInfo, reader io.Reader) (err error) {
  var buf bytes.Buffer
  limitedReader := &io.LimitedReader{R: reader, N: int64(defaultFileConfig.maxFileSize + 1)}
  
  _, err = io.Copy(&buf, limitedReader)
	size := buf.Len()
	if size < defaultFileConfig.minFileSize {
		log.Println("File failed validation: too small.", size, defaultFileConfig.minFileSize)
		fileInfo.Error = "minFileSize"
		err = errors.New("Too Small")
		return
	} else if size > defaultFileConfig.maxFileSize {
		log.Println("File failed validation: too large.", size, defaultFileConfig.maxFileSize)
		fileInfo.Error = "maxFileSize"
    err = errors.New("Too Big")
		return
	}
	fileInfo.Size = int64(size)
	log.Println("yay!!! file!!!")
	return
}

func isStoreableFile(fileInfo *FileInfo) bool {
  return (fileInfo.Backend == "s3" && true)
}