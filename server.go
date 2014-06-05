package main

import (
  "./controller"
  "github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()
  controller.AddRoute(m)
  m.Run()
}
