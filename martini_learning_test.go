package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MartiniLearning", func() {
  Context("/", func() {
    It("return 200 status", func() {
      Request("GET", "/",  top)
      Expect(response.Code).To(Equal(200))
      Expect(response.Body).To(ContainSubstring("Hello"))
    })
  })
})
