package main_test

import (
	. "github.com/fmd/micro-app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MicroApp", func() {
    Describe("Testing GET requests", func() {
        Context("using the root URL", func() {
            It("should return OK", func() {
                //noop
            })
        })

        Context("using the /submit URL", func() {
            It("should return 404", func() {
                //noop
            })
        })

        Context("using the /error URL", func() {
            It("should return 500", func() {
                //noop
            })
        })
    })

    Describe("Testing POST requests", func() {
        Context("using the root URL", func() {
            It("should return 404", func() {
                //noop
            })
        })

        Context("using the /submit URL", func() {
            It("should return 401", func() {
                //noop
            })
        })

        Context("using the /error URL", func() {
            It("should return 404", func() {
                //noop
            })
        })
    })
})
