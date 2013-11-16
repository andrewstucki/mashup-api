package add

import(
  "github.com/mashup-cms/mashup-api/connections"
  "github.com/mashup-cms/mashup-api/model"
	"github.com/mashup-cms/mashup-api/commands"
  
  "log"
)

var AdminCmd = &commands.Command{
	Name:    "add:admin",
	Usage:   "",
	Summary: "Creates admin user",
	Help:    `create extended help here...`,
	Run:     createAdmin,
}

func createAdmin(cmd *commands.Command, args ...string) {
	connections.SetupPostgres()
	_, err := model.CreateUser("admin", "test123")
	if err != nil {
		log.Printf(err.Error())
	}
}