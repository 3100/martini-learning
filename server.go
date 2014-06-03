package main

import (
  "github.com/go-martini/martini"
  "github.com/yosssi/rendergold"
)

func main() {
  m := martini.Classic()
  // reads "templates" directory by default
  m.Use(rendergold.Renderer())

  //m.Get("/", func(r rendergold.Render) {
    // parses "templates/top.gold"
  //  r.HTML(200, "top", nil)
  //})
  m.Get("/", top)

  m.Run()
}

func top(r rendergold.Render) {
  r.HTML(200, "top", nil)
}
