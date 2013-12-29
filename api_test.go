package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
        . "../go-webservice"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Web", func() {
	var (
		page *Page
	)

	BeforeEach(func() {
		page = &Page{Title: "foo", Body: []byte("Example body")}
                page.Save()
	})

	Describe("Fetching a page", func() {
		Context("by ID", func() {
			It("has a JSON representation", func() {
				request, _ := http.NewRequest("GET", "/pages/foo", nil)
				response := httptest.NewRecorder()
				ViewHandler(response, request, "foo")
                                expectedJSON := `{"title":"foo","body":"Example body"}`
				Expect(response.Body.String()).To(Equal(expectedJSON))
			})

                        It("returns a 404 if doesn't exist", func() {
				request, _ := http.NewRequest("GET", "/pages/bar", nil)
				response := httptest.NewRecorder()
				ViewHandler(response, request, "bar")
				Expect(response.Code).To(Equal(404))
                        })
		})
	})

        Describe("Fetching all pages", func() {
                It("returns a JSON array", func() {
			request, _ := http.NewRequest("GET", "/pages", nil)
			response := httptest.NewRecorder()
			IndexHandler(response, request)
                        expectedJSON := `[{"title":"foo","body":"Example body"}]`
			Expect(response.Body.String()).To(Equal(expectedJSON))
                })
        })
})
