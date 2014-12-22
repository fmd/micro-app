package main_test

import (
    . "github.com/fmd/micro-app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

    "net/http"
    "bytes"
    "fmt"
)

var url string = "http://localhost:3000"

var _ = Describe("MicroApp", func() {
    var (
        app *App
    )

    app = &App{}
    app.Init()
    go app.Run()

    Describe("Testing GET requests", func() {
        Context("using the root URL", func() {
            It("should return OK", func() {
                resp, err := http.Get(url)
                Expect(err).To(BeZero())
                Expect(resp.StatusCode).To(Equal(200))
            })
        })

        Context("using the /submit URL", func() {
            It("should return 404", func() {
                resp, err := http.Get(fmt.Sprintf("%s/submit", url))
                Expect(err).To(BeZero())
                Expect(resp.StatusCode).To(Equal(404))
            })
        })

        Context("using the /error URL", func() {
            It("should return 500", func() {
                resp, err := http.Get(fmt.Sprintf("%s/error",url))
                Expect(err).To(BeZero())
                Expect(resp.StatusCode).To(Equal(500))
            })
        })
    })

    Describe("Testing POST requests", func() {
        Context("using the root URL", func() {
            It("should return 404", func() {
                resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(`"data"`)))
                Expect(err).To(BeZero())
                Expect(resp.StatusCode).To(Equal(404))
            })
        })

        Context("using the /submit URL", func() {
            It("should return 401", func() {
                resp, err := http.Post(fmt.Sprintf("%s/submit", url), "application/json", bytes.NewBuffer([]byte(`"data"`)))
                Expect(err).To(BeZero())
                Expect(resp.StatusCode).To(Equal(401))
            })
        })

        Context("using the /error URL", func() {
            It("should return 404", func() {
                resp, err := http.Post(fmt.Sprintf("%s/error", url), "application/json", bytes.NewBuffer([]byte(`"data"`)))
                Expect(err).To(BeZero())
                Expect(resp.StatusCode).To(Equal(404))
            })
        })
    })
})
