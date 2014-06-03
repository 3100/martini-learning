package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

    "github.com/go-martini/martini"
    "github.com/yosssi/rendergold"
    "net/http"
    "net/http/httptest"
    "testing"
)

var (
  response *httptest.ResponseRecorder
)

func Request(method string, route string, handler martini.Handler) {
    m := martini.Classic()
    m.Use(rendergold.Renderer())
    m.Get(route, handler)
    request, _ := http.NewRequest(method, route, nil)
    response = httptest.NewRecorder()
    m.ServeHTTP(response, request)
}

func TestMartiniLearning(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MartiniLearning Suite")
}
