package workers

import (
  "github.com/mashup-cms/mashup-api/websocket"
  
  "github.com/mashup-cms/mashup-api/services/github"
)

type task struct {
  id string
  requestorId string
  action string
  args interface{}
}

func (t *task) run() {
  var message *websocket.Message
  switch t.action {
    case "SynchronizeUser":
      err := github.SyncGithubAccount("internal", t.args)
      if err != nil {
        message = &websocket.Message{t.id, "error"}        
      } else {
        message = &websocket.Message{t.id, "done"}
      }
    default:
      //do nothing
  }
  websocket.Hub.SendMessage(t.requestorId, message)
}

func NewTask(taskId, requestor, action string, args interface{}) {
  tasks <- &task{taskId, requestor, action, args}
}