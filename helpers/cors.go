package helpers

import (
	"github.com/emicklei/go-restful"
)

func EnableCORS(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	if origin := req.Request.Header.Get("Origin"); origin != "" {
		resp.AddHeader("Access-Control-Allow-Origin", origin)
		resp.AddHeader("Access-Control-Allow-Credentials", "true")
		resp.AddHeader("Access-Control-Expose-Headers", "Content-Type, Cache-Control, Expires, Etag, Last-Modified")

		if req.Request.Method == "OPTIONS" {
			resp.AddHeader("Access-Control-Allow-Methods", "HEAD, GET, POST, PATCH, PUT, DELETE")
			resp.AddHeader("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, If-None-Match, If-Modified-Since")
			return
		}

		chain.ProcessFilter(req, resp)

	}
}
