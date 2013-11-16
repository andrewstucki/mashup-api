package helpers

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

type service interface {
	GetService() interface{}
}

func ServiceResponse(response *restful.Response, object interface{}, err error) {
	if err == nil {
		if serviceObj, ok := object.(service); ok {
			response.WriteEntity(serviceObj.GetService())
		} else {
			response.WriteEntity(object)
		}
		return
	}
	response.WriteErrorString(http.StatusNotFound, err.Error())
}
