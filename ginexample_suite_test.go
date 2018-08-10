package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinexample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginexample Suite")
}
