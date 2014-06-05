package controller

import (
  "github.com/go-martini/martini"
  "github.com/yosssi/rendergold"
)

func AddRoute(m *martini.ClassicMartini) {
  // reads "templates" directory by default
  m.Use(rendergold.Renderer())

  m.Get("/", top)
}

func top(r rendergold.Render) {
  r.HTML(200, "top", nil)
}
