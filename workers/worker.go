package workers

import (
  "runtime"
)

func maxParallelism() int {
    maxProcs := runtime.GOMAXPROCS(0)
    numCPU := runtime.NumCPU()
    if maxProcs < numCPU {
        return maxProcs
    }
    return numCPU
}

var tasks chan *task

func StartWorkers() {
  tasks = make(chan *task, 256)
  for i := 0; i < maxParallelism(); i++ {
    go handleTask()
  }
}

func handleTask() {
  for task := range tasks {
    task.run()
  }
}
