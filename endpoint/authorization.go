package endpoint

import (
	"github.com/emicklei/go-restful"

	"github.com/mashup-cms/mashup-api/helpers"
	"github.com/mashup-cms/mashup-api/middleware"
	"github.com/mashup-cms/mashup-api/model"
	"github.com/mashup-cms/mashup-api/services"
)

type AuthEndpoint struct{}
type Credentials struct {
	Username string
	Password string
}

func (r AuthEndpoint) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/auth").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.POST("/login").To(r.login).
		// docs
		Doc("login").
		Param(ws.BodyParameter("username", "identifier of the user").DataType("string")).
		Param(ws.BodyParameter("password", "password of the user").DataType("string")).
		Writes(model.AccessToken{})) // on the response

	ws.Route(ws.GET("/user").Filter(middleware.AuthCheck).To(r.user).
		// docs
		Doc("login").
		Param(ws.BodyParameter("username", "identifier of the user").DataType("string")).
		Param(ws.BodyParameter("password", "password of the user").DataType("string")).
		Writes(model.AccessToken{})) // on the response

	container.Add(ws)
}

func (r AuthEndpoint) login(request *restful.Request, response *restful.Response) {
	credentials := &Credentials{}
	err := request.ReadEntity(credentials)
	matched := false
	var token *model.AccessToken
	if err == nil {
		user, err := services.FindUserByName(credentials.Username, -1)
		if err == nil {
			matched = user.CheckPassword(credentials.Password)
			if matched {
				token = model.GetReusableAccessToken(user.Id, 1)
				if token == nil {
					token = model.GenerateAccessToken(user.Id, 1)
				}
			}
		}
	}
	helpers.ServiceResponse(response, token, err)
}

func (r AuthEndpoint) user(request *restful.Request, response *restful.Response) {
	authToken := request.Attribute("authToken").(*model.AccessToken)
	account, err := services.FindUserById(authToken.UserId)
	helpers.ServiceResponse(response, account, err)
}
