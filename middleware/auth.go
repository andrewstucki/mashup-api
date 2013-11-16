package middleware

import (
  "github.com/emicklei/go-restful"
  "strings"
  "github.com/mashup-cms/mashup-api/model"
)

func AuthCheck(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
  mashupKey := req.Request.Header.Get("X-Mashup-Key")
  if mashupKey != "" {
    token := model.FindAccessToken(mashupKey)
    if token != nil {
      req.SetAttribute("authToken", token)
      chain.ProcessFilter(req, resp)
      return
    }
  }
  encoded := req.Request.Header.Get("Authorization")
  tokens := strings.SplitN(encoded, " ", 2)
  if len(tokens) == 2 {
    method, payload := tokens[0], tokens[1]
    if method == "token" {
      token := model.FindAccessToken(payload)
      if token != nil {
        req.SetAttribute("authToken", token)
        chain.ProcessFilter(req, resp)
        return
      }
    }
  }
  resp.WriteErrorString(401, "401: Not Authorized")
  return
}