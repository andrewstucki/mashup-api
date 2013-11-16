package endpoint

import (
  "github.com/emicklei/go-restful"
  
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/middleware"
  "github.com/mashup-cms/mashup-api/helpers"
  "github.com/mashup-cms/mashup-api/services"
)

type UserEndpoint struct {}

func (endpoint UserEndpoint) Register(container *restful.Container) {
  ws := new(restful.WebService)
  ws.
    Path("/users").
    Consumes(restful.MIME_XML, restful.MIME_JSON).
    Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

  ws.Route(ws.GET("/").Filter(middleware.AuthCheck).To(endpoint.getUsers).
    // docs
    Doc("Get all users").
    Writes(model.User{})) // on the response

  ws.Route(ws.POST("/").Filter(middleware.AuthCheck).To(endpoint.createUser).
    // docs
    Doc("Create a user").
    Writes(model.User{})) // on the response

  ws.Route(ws.GET("/{id}").Filter(middleware.AuthCheck).To(endpoint.getUser).
    // docs
    Doc("Get a user").
    Writes(model.User{})) // on the response

  //something for currentUser?

  container.Add(ws)
}

func (endpoint UserEndpoint) getUsers(request *restful.Request, response *restful.Response) {
  params := helpers.QueryParameters(request)
  authToken := request.Attribute("authToken").(*model.AccessToken)
  users, err := services.FindUsers(params, authToken.UserId)
  helpers.ServiceResponse(response, users, err)
}

func (endpoint UserEndpoint) getUser(request *restful.Request, response *restful.Response) {
  id := request.PathParameter("id")
  authToken := request.Attribute("authToken").(*model.AccessToken)
  account, err := services.FindUser(id, authToken.UserId)
  helpers.ServiceResponse(response, account, err)
}

func (endpoint UserEndpoint) createUser(request *restful.Request, response *restful.Response) {
}