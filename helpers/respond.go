package helpers

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"log"
)

func ServiceResponse(response *restful.Response, object interface{}, err error) {
  log.Printf("response %s", err)
	if err == nil {
		response.WriteEntity(object)
		return
	}
	response.WriteErrorString(http.StatusNotFound, err.Error())
}
