package store

import (
  "regexp"
  "errors"
  "io"
  "log"
)

type FileInfo struct {
	Key            string `json:"-"`
	ExpirationTime int64  `json:"-"`
	Url            string `json:"url,omitempty"`
	ThumbnailUrl   string `json:"thumbnail_url,omitempty"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Backend        string `json:"-"`
	Size           int64  `json:"size"`
	Error          string `json:"error,omitempty"`
	DeleteUrl      string `json:"delete_url,omitempty"`
	DeleteType     string `json:"delete_type,omitempty"`
}

type config struct {
	minFileSize        int // bytes
	maxFileSize        int // bytes
	acceptFileTypes    *regexp.Regexp
}

func Create(fileInfo *FileInfo, reader io.Reader) error {
  log.Printf("%s", fileInfo)
  if isStoreableImage(fileInfo) {
    return createImage(fileInfo, reader)
	} else if isStoreableVideo(fileInfo) {
	  return createVideo(fileInfo, reader)
	} else if isStoreableFile(fileInfo) {
	  return createFile(fileInfo, reader)
	} else {
	  fileInfo.Error = "acceptFileTypes"
		return errors.New("Bad File Type")
	}
}

func Get(key string) (*FileInfo, io.Reader, error) {
  return nil, nil, nil
}

func Delete(key string) error {
  return nil
}