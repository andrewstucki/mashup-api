package helpers

import (
	"github.com/emicklei/go-restful"
)

func QueryParameters(request *restful.Request) map[string][]string {
	return request.Request.URL.Query()
}
