package server

import(
  "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
  
  "github.com/mashup-cms/mashup-api/connections"
  "github.com/mashup-cms/mashup-api/helpers"
  "github.com/mashup-cms/mashup-api/endpoint"
  "github.com/mashup-cms/mashup-api/commands"
  
  "log"
  "time"
  "net/http"
)

var RunCmd = &commands.Command{
	Name:    "server:run",
	Usage:   "",
	Summary: "Run the server",
	Help:    `create extended help here...`,
	Run:     runServer,
}

func runServer(cmd *commands.Command, args ...string) {
	connections.SetupRedis()
	connections.SetupPostgres()

	helpers.RandomSource.Seed(time.Now().UTC().UnixNano()) //for generating secure tokens

	wsContainer := restful.NewContainer()

	auth := endpoint.AuthEndpoint{}
	auth.Register(wsContainer)
	github := endpoint.GithubEndpoint{}
	github.Register(wsContainer)
	repos := endpoint.RepoEndpoint{}
	repos.Register(wsContainer)
	users := endpoint.UserEndpoint{}
	users.Register(wsContainer)

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs and enter http://localhost:8080/apidocs.json in the api input field.
	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:3457",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "../../swagger-ui/dist",
	}
	swagger.RegisterSwaggerService(config, wsContainer)

	wsContainer.Filter(helpers.EnableCORS)

	log.Printf("start listening on localhost:3457")
	server := &http.Server{Addr: ":3457", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
