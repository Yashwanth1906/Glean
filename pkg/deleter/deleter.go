package deleter


import (
  "fmt"
  "os"
  "sync"
)


func DeleteTargets(paths []string, workers int) {
  var wg sync.WaitGroup
  jobs := make(chan string, len(paths))

  for i := 0; i < workers;i++ {
    wg.Add(1)
    go func(id int) {
      defer wg.Done()
      for path := range jobs {
        err := os.RemoveAll(path)
        if err != nil {
          fmt.Printf("❌ Worker %d: Failed to delete %s: %v\n", id, path, err)
        } else {
          fmt.Printf("🧹 Worker %d: Deleted %s\n", id, path)
        }
      }
    }(i+1)
  }

  for _, path := range paths {
    jobs <- path
  }
  close(jobs)
  wg.Wait()
}
