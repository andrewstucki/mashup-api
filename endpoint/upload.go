package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"errors"
	// Register image handling libraries by importing them.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	
	"github.com/mashup-cms/mashup-api/store"
	"github.com/mashup-cms/mashup-api/model"
)

var FileNotFoundError = errors.New("File Not Found")

// UploadHandler is an http Handler for uploading files and serving thumbnails.
type UploadHandler struct {}

// http500 Raises an HTTP 500 Internal Server Error if err is non-nil
func http500(w http.ResponseWriter, err error) {
	if err != nil {
		msg := "500 Internal Server Error: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		log.Panic(err)
	}
}

func escape(s string) string {
	return strings.Replace(url.QueryEscape(s), "+", "%20", -1)
}

func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {	
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Expose-Headers", "Content-Type, Cache-Control, Expires, Etag, Last-Modified, X-Mashup-Key, X-Backend, X-Repo-Id")

		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, PATCH, PUT, DELETE")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, If-None-Match, If-Modified-Since, X-Mashup-Key, X-Backend, X-Repo-Id")
			return
		}

  	mashupKey := r.Header.Get("X-Mashup-Key")
    if mashupKey != "" {
      token := model.FindAccessToken(mashupKey)
      if token != nil {
  
        params, err := url.ParseQuery(r.URL.RawQuery)
      	http500(w, err)
    	
      	switch r.Method {
      	case "GET":
      		h.get(w, r)
      	case "POST":
      		if len(params["_method"]) > 0 && params["_method"][0] == "DELETE" {
      			h.delete(w, r)
      		} else {
      			h.post(w, r)
      		}
      	case "DELETE":
      		h.delete(w, r)
      	default:
      		http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
      	}
      } else {
        http.Error(w, "401: Not Authorized", http.StatusUnauthorized)
      }
    } else {
      http.Error(w, "401: Not Authorized", http.StatusUnauthorized)
    }
  }
}

func (h *UploadHandler) get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// Noop - not sure why JFU plugin sometimes calls this with no args
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] == "" {
		log.Printf("Invalid URL: '%v'", r.URL)
		http.Error(w, "404 Invalid URL", http.StatusNotFound)
		return
	}
	key := parts[1]
	log.Println("Get", key)
	fi, file, err := store.Get(parts[1])
	if err == FileNotFoundError {
		log.Println("Not Found", key)
		http.NotFound(w, r)
		return
	}
	http500(w, err)
	log.Println("Found", key, fi.Type)
	w.Header().Add(
		"Cache-Control",
		fmt.Sprintf("public,max-age=%d", fi.ExpirationTime),
	)
/*  if imageRegex.MatchString(fi.Type) {
    w.Header().Add("X-Content-Type-Options", "nosniff")
  } else {
*/		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add(
			"Content-Disposition:",
			fmt.Sprintf("attachment; filename=%s;", parts[2]),
		)
/*  }*/
	io.Copy(w, file)
	return
}

// uploadFile handles the upload of a single file from a multipart form.  
func (h *UploadHandler) uploadFile(w http.ResponseWriter, p *multipart.Part, backend, repoId string) (fi *store.FileInfo) {
	log.Printf("%s", p.Header)
	fi = &store.FileInfo{
		Name: p.FileName(),
		Type: p.Header.Get("Content-Type"),
		Backend: backend,
	}
	err := store.Create(fi, p)
  http500(w, err)
  return	
}

func getFormValue(p *multipart.Part) string {
	var b bytes.Buffer
	io.CopyN(&b, p, int64(1<<20)) // Copy max: 1 MiB
	return b.String()
}

func (h *UploadHandler) post(w http.ResponseWriter, r *http.Request) {
	fileInfos := make([]*store.FileInfo, 0)
	mr, err := r.MultipartReader()
	http500(w, err)
	r.Form, err = url.ParseQuery(r.URL.RawQuery)
	http500(w, err)
	for {
		// var part *multipart.Part
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		http500(w, err)
		if name := part.FormName(); name != "" {
			if part.FileName() != "" {
				fileInfos = append(fileInfos, h.uploadFile(w, part, r.Header.Get("X-Backend"), r.Header.Get("X-Repo-Id")))
			} else {
				r.Form[name] = append(r.Form[name], getFormValue(part))
			}
		}
	}

	js, err := json.Marshal(fileInfos)
	http500(w, err)
	if redirect := r.FormValue("redirect"); redirect != "" {
		http.Redirect(w, r, fmt.Sprintf(
			redirect,
			escape(string(js)),
		), http.StatusFound)
		return
	}
	jsonType := "application/json"
	if strings.Index(r.Header.Get("Accept"), jsonType) != -1 {
		w.Header().Set("Content-Type", jsonType)
	}
	fmt.Fprintln(w, string(js))
}

func (h *UploadHandler) delete(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] == "" {
		log.Println("Invalid URL:", r.URL)
		http.Error(w, "404 Invalid URL", http.StatusNotFound)
		return
	}
	key := parts[1]
	log.Println("Delete", key)
	err := store.Delete(key)
	http500(w, err)
	return
}