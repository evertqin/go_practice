package main

// TODO: continue: https://gobyexample.com/timeouts

import (
  "fmt"
  "time"
)

func ordinary_channel (ch chan string) {
  time.Sleep(1000)
  ch <- "sample"
}

func buffered_channel(ch chan string) {
  ch <- "buffered_channel"
  ch <- "buffered_channel1"
}

func pongs(pings <-chan string, pongs chan<- string) {
  pongs <- <-pings
}

func worker(done chan bool) {
  fmt.Println("working")
  time.Sleep(time.Second)
  fmt.Println("done")
  done <- true

}

func main() {
  ch0 := make(chan string)
  go ordinary_channel(ch0)

  fmt.Println(<-ch0)

  ch1 := make(chan string, 2)
  go buffered_channel(ch1)

  fmt.Println(<-ch1)
  fmt.Println(<-ch1)

  ch1 <- "hmm"
  go pongs(ch0, ch1)
  fmt.Println(<-ch1)

  done:= make(chan bool, 1)
  go worker(done)
  <-done

  c1 := make(chan string)
  c2 := make(chan string)


  go func() {
    time.Sleep(time.Second * 1)
    c1 <- "one"
  }()
  go func() {
    time.Sleep(time.Second * 2)
    c2 <- "two"
  }()


  for i := 0; i < 2; i++ {
    select {
    case msg1 := <-c1:
      fmt.Println("received", msg1)
    case msg2 := <-c2:
      fmt.Println("received", msg2)
    }
  }
}
