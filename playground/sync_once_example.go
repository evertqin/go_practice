package main

import (
  "sync"
  "fmt"
)

func main() {
  var once sync.Once
  onceBody := func() {
    fmt.Print("Only Once")
  }

  done := make(chan bool)
  for i := 0; i < 10; i++ {
    go func() {
      once.Do(onceBody)
      done <- true
    }()
  }

  for i := 0; i < 10; i ++ {
    <-done
  }
}
