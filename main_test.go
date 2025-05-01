// File: main_test.go
package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoogleSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Google Search Suite")
}
