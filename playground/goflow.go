package main

import (
  "fmt"
  "sync"
  "time"
)

type Flow struct {
  funcs map[string]*flowStruct
}

type flowFunc func(res map[string]interface{}) (interface{}, error)

type flowStruct struct {
  Deps []string
  Ctr  int
  Fn   flowFunc
  C    chan interface{}
  once sync.Once
}

func (fs *flowStruct) Done(r interface{}) {
  for i := 0; i < fs.Ctr; i++ {
    fs.C <- r
  }
}

func (fs *flowStruct) Close() {
  fs.once.Do(func() {
    close(fs.C)
  })
}

func (fs *flowStruct) Init() {
  fs.C = make(chan interface{}, fs.Ctr)
}

func New() *Flow {
  return &Flow{
    funcs: make(map[string]*flowStruct),
  }
}

func (flw *Flow) Add(name string, d []string, fn flowFunc) *Flow {
  flw.funcs[name] = &flowStruct{
    Deps: d,
    Fn:   fn,
    Ctr:  1,
  }
  return flw
}

func (flw *Flow) Do() (map[string]interface{}, error) {
  for name, fn := range flw.funcs {
    for _, dep := range fn.Deps {
      if dep == name {
        return nil, fmt.Errorf("Error: Function \"%s\" depends on it self", name)
      }

      if _, exists := flw.funcs[dep]; exists == false {
        return nil, fmt.Errorf("Error: Function \"%s\" not exits!", dep)
      }

      flw.funcs[dep].Ctr++
    }
  }
  return flw.do()
}

func (flw *Flow) do() (map[string]interface{}, error) {
  var err error
  res := make(map[string]interface{}, len(flw.funcs))

  for name, f := range flw.funcs {
    f.Init()
    go func(name string, fs *flowStruct) {
      defer func() { fs.Close() }()
      results := make(map[string]interface{}, len(fs.Deps))

      for _, dep := range fs.Deps {
        results[dep] = <-flw.funcs[dep].C
      }

      r, fnErr := fs.Fn(results)
      if fnErr != nil {
        for _, fn := range flw.funcs {
          fn.Close()
        }
        err = fnErr
        return
      }
      if err != nil {
        return
      }
      fs.Done(r)
    }(name, f)

  }

  for name, fs := range flw.funcs {
    res[name] = <-fs.C
  }

  return res, err
}

func main() {
  f1 := func(r map[string]interface{}) (interface{}, error) {
    fmt.Println("function1 started")
    time.Sleep(time.Millisecond * 1000)
    return 1, nil
  }

  f2 := func(r map[string]interface{}) (interface{}, error) {
    time.Sleep(time.Millisecond * 1000)
    fmt.Println("function2 started", r["f1"])
    return "some results", nil
  }

  f3 := func(r map[string]interface{}) (interface{}, error) {
    fmt.Println("function3 started", r["f1"])
    return nil, nil
  }

  f4 := func(r map[string]interface{}) (interface{}, error) {
    fmt.Println("function4 started", r)
    return nil, nil
  }

  res, err := New().
      Add("f1", nil, f1).
      Add("f2", []string{"f1"}, f2).
      Add("f3", []string{"f1"}, f3).
      Add("f4", []string{"f2", "f3"}, f4).
      Do()

  fmt.Println(res, err)
}
