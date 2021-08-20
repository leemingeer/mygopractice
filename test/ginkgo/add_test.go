package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//. "qualified.io/challenge"
	"testing"
)

// as long as it begins with Test since it’s just a regular Go test function
func TestCart(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}



var _ = Describe("Add", func() {
	It("should add integers", func() {
		Expect(Add(1, 1)).To(Equal(2))
	})

	It("should add minus integers", func() {
		Expect(Add(-1, -1)).To(Equal(-2))
	})
})

var _ = Describe("Addtens", func() {
	It("should add integers", func() {
		Expect(Add(10, 20)).To(Equal(30))
	})
})
