package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMicroApp(t *testing.T) {
	RegisterFailHandler(Fail)
    RunSpecs(t, "MicroApp Suite")
}
