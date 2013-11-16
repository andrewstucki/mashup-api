package endpoint

import (
  "github.com/emicklei/go-restful"
  
  "github.com/mashup-cms/mashup-api/model"
  "github.com/mashup-cms/mashup-api/middleware"
  "github.com/mashup-cms/mashup-api/helpers"
  "github.com/mashup-cms/mashup-api/services"
  
  "strconv"
  // "./repos"
)

type RepoEndpoint struct {}
type RepoDataWrapper struct {
  Repo *RepoData `json:repo`
}
type RepoData struct {
  Name string `json:name`
  Owner string `json:owner`
  DefaultBranch string `json:description`
  Description string `json:description`
  Active bool `json:active`
}


func (endpoint RepoEndpoint) Register(container *restful.Container) {
  ws := new(restful.WebService)
  ws.
    Path("/repos").
    Consumes(restful.MIME_XML, restful.MIME_JSON).
    Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

  ws.Route(ws.GET("/").Filter(middleware.AuthCheck).To(endpoint.getRepos).
    // docs
    Doc("Get Repositories").
    Writes(model.Repos{})) // on the response

  ws.Route(ws.PUT("/{id}").Filter(middleware.AuthCheck).To(endpoint.putRepo).
    // docs
    Doc("Update Repository").
    Writes(model.Repos{})) // on the response

  ws.Route(ws.GET("/{account}/{name}").Filter(middleware.AuthCheck).To(endpoint.getRepo).
    // docs
    Doc("Get a single Repository").
    Writes(model.RepoService{})) // on the response

  ws.Route(ws.POST("/{account}/{name}/activate").Filter(middleware.AuthCheck).To(endpoint.toggleActivateRepo).
    // docs
    Doc("Set repository as active").
    Writes(model.RepoService{})) // on the response

  container.Add(ws)
}

func (endpoint RepoEndpoint) getRepos(request *restful.Request, response *restful.Response) {
  params := helpers.QueryParameters(request)
  authToken := request.Attribute("authToken").(*model.AccessToken)
  repos, err := services.FindRepos(params, authToken.UserId)
  helpers.ServiceResponse(response, repos, err)
}

func (endpoint RepoEndpoint) getRepo(request *restful.Request, response *restful.Response) {
  account := request.PathParameter("account")
  name := request.PathParameter("name")
  authToken := request.Attribute("authToken").(*model.AccessToken)
  repo, err := services.FindRepo(account, name, authToken.UserId)
  helpers.ServiceResponse(response, repo, err)
}

func (endpoint RepoEndpoint) putRepo(request *restful.Request, response *restful.Response) {
  repoWrapper := &RepoDataWrapper{}
  err := request.ReadEntity(repoWrapper)
  if err != nil {
    return
  }
  repo := repoWrapper.Repo
  repoIdString := request.PathParameter("id")
  repoId, err := strconv.ParseInt(repoIdString, 0, 0)
  if err != nil {
    return
  }
  authToken := request.Attribute("authToken").(*model.AccessToken)
  realRepo, err := services.UpdateRepo(int(repoId), repo.Active, authToken.UserId)
  helpers.ServiceResponse(response, realRepo, err)
}

func (endpoint RepoEndpoint) toggleActivateRepo(request *restful.Request, response *restful.Response) {
  //UPDATE dbo.Table1 SET col2 = (CASE col2 WHEN 1 THEN 0 ELSE 1 END);
  //should optimize, only 1 sql transaction
  account := request.PathParameter("account")
  name := request.PathParameter("name")
  authToken := request.Attribute("authToken").(*model.AccessToken)
  repo, err := services.FindRepo(account, name, authToken.UserId)
  if err == nil {
    err = repo.Activate()
  }
  helpers.ServiceResponse(response, repo, err)
}