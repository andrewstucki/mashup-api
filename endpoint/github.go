package endpoint

import (
	"github.com/emicklei/go-restful"

	"github.com/mashup-cms/mashup-api/helpers"
	"github.com/mashup-cms/mashup-api/middleware"
	"github.com/mashup-cms/mashup-api/model"
	"github.com/mashup-cms/mashup-api/services"

	"log"
)

type GithubEndpoint struct{}
type GithubToken struct {
	Token string `json:"githubToken"`
}

func (endpoint GithubEndpoint) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/githubAccounts").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.POST("/token").Filter(middleware.AuthCheck).To(endpoint.addToken).
		// docs
		Doc("Add a GitHub account").
		Reads(GithubToken{Token: "123"}).
		Writes(model.GithubAccount{})) // on the response

	ws.Route(ws.GET("/").Filter(middleware.AuthCheck).To(endpoint.getAccounts).
		// docs
		Doc("Get GitHub accounts").
		Writes(model.GithubAccounts{})) // on the response

	ws.Route(ws.GET("/{login}").Filter(middleware.AuthCheck).To(endpoint.getAccount).
		// docs
		Doc("Get a single GitHub account").
		Param(ws.PathParameter("login", "account name").DataType("string")).
		Writes(model.GithubAccountService{})) // on the response

	ws.Route(ws.POST("/{login}/sync").Filter(middleware.AuthCheck).To(endpoint.syncAccount).
		// docs
		Doc("Sync repositories from a GitHub account").
		Param(ws.PathParameter("login", "account name").DataType("string")).
		Writes(model.GithubAccountService{})) // on the response

	ws.Route(ws.GET("/{login}/admins").Filter(middleware.AuthCheck).To(endpoint.getAdmins).
		// docs
		Doc("Sync repositories from a GitHub account").
		Param(ws.PathParameter("login", "account name").DataType("string")).
		Writes(model.GithubAccountService{})) // on the response

	ws.Route(ws.POST("/{login}/admins").Filter(middleware.AuthCheck).To(endpoint.addAdmins).
		// docs
		Doc("Sync repositories from a GitHub account").
		Param(ws.PathParameter("login", "account name").DataType("string")).
		Writes(model.GithubAccountService{})) // on the response

	ws.Route(ws.DELETE("/{login}/admins").Filter(middleware.AuthCheck).To(endpoint.removeAdmins).
		// docs
		Doc("Sync repositories from a GitHub account").
		Param(ws.PathParameter("login", "account name").DataType("string")).
		Writes(model.GithubAccountService{})) // on the response

	container.Add(ws)
}

func (endpoint GithubEndpoint) getAdmins(request *restful.Request, response *restful.Response) {
	login := request.PathParameter("login")
	authToken := request.Attribute("authToken").(*model.AccessToken)
	accounts, err := services.FindGithubAdmins(login, authToken.UserId)
	helpers.ServiceResponse(response, accounts, err)
}

func (endpoint GithubEndpoint) addAdmins(request *restful.Request, response *restful.Response) {
	login := request.PathParameter("login")
	authToken := request.Attribute("authToken").(*model.AccessToken)
	users := &model.Users{}
	err := request.ReadEntity(users)
	if err == nil && authToken != nil && login != "" {
		err = services.AddGithubAdmins(login, users, authToken.UserId)
	} else {
		log.Printf(err.Error())
	}
	helpers.ServiceResponse(response, users, err)
}

func (endpoint GithubEndpoint) removeAdmins(request *restful.Request, response *restful.Response) {
	login := request.PathParameter("login")
	authToken := request.Attribute("authToken").(*model.AccessToken)
	users := &model.Users{}
	err := request.ReadEntity(users)
	if err == nil && authToken != nil && login != "" {
		err = services.RemoveGithubAdmins(login, users, authToken.UserId)
	} else {
		log.Printf(err.Error())
	}
	helpers.ServiceResponse(response, users, err)
}

func (endpoint GithubEndpoint) getAccount(request *restful.Request, response *restful.Response) {
	login := request.PathParameter("login")
	authToken := request.Attribute("authToken").(*model.AccessToken)
	account, err := services.FindGithubAccount(login, authToken.UserId)
	helpers.ServiceResponse(response, account, err)
}

func (endpoint GithubEndpoint) syncAccount(request *restful.Request, response *restful.Response) {
	login := request.PathParameter("login")
	authToken := request.Attribute("authToken").(*model.AccessToken)
	account, err := services.SyncGithubAccount(login, authToken.UserId)
	helpers.ServiceResponse(response, account, err)
}

func (endpoint GithubEndpoint) getAccounts(request *restful.Request, response *restful.Response) {
	params := helpers.QueryParameters(request)
	authToken := request.Attribute("authToken").(*model.AccessToken)
	accounts, err := services.FindGithubAccounts(params, authToken.UserId)
	helpers.ServiceResponse(response, accounts, err)
}

func (endpoint GithubEndpoint) addToken(request *restful.Request, response *restful.Response) {
	token := &GithubToken{}
	err := request.ReadEntity(token)
	account := &model.GithubAccount{}
	if err == nil {
		authToken := request.Attribute("authToken").(*model.AccessToken)
		account, err = services.AddGithubToken(token.Token, authToken.UserId)
	}
	helpers.ServiceResponse(response, account, err)
}
