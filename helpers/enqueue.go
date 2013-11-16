package helpers

import (
  "github.com/mashup-cms/mashup-api/globals"
  "encoding/json"
)

type EnqueueData struct {
  Class string `json:"class"`
  Args  interface{} `json:"args"`
}

func Enqueue(queue, klass string, args interface{}) error {
  connection := globals.RedisPool.Get()
  defer connection.Close()
  bytes, err := json.Marshal(EnqueueData{klass, args})
  if err != nil {
    return err
  }

  _, err = connection.Do("sadd", "queues", queue)
  if err != nil {
    return err
  }
  queue = "queue:" + queue
  _, err = connection.Do("rpush", queue, bytes)
  return err
}