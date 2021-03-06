package api_test

import (
	. "../go-webservice"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

        "io/ioutil"
	"net/http"
	"net/http/httptest"
        "os"
        "path/filepath"
        "strings"
)

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var _ = Describe("Web", func() {
	var (
		page *Page
	)

	BeforeEach(func() {
		page = &Page{Title: "foo", Body: []byte("Example body")}
		page.Save()
	})

        AfterEach(func() {
	  files, _ := filepath.Glob("/Users/glenngillen/Development/go-webservice/*.txt")
          for _, file := range files {
            os.Remove(file)
          }
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

	Describe("Creating a page", func() {
		It("is created from JSON", func() {
			jsonObject := `{"title":"bar","body":"New page body"}`
			request, _ := http.NewRequest("POST", "/pages", strings.NewReader(jsonObject))
			response := httptest.NewRecorder()
			SaveHandler(response, request)
                        page, _:= loadPage("bar")
                        Expect(response.Code).To(Equal(303))
                        Expect(string(page.Body[:])).To(Equal("New page body"))
		})
	})

	Describe("Updating a page", func() {
		It("is updated from JSON", func() {
			jsonObject := `{"title":"foo","body":"My new page body"}`
			request, _ := http.NewRequest("PUT", "/pages/foo", strings.NewReader(jsonObject))
			response := httptest.NewRecorder()
			UpdateHandler(response, request)
                        page, _:= loadPage("foo")
                        Expect(response.Code).To(Equal(204))
                        Expect(string(page.Body[:])).To(Equal("My new page body"))
		})
	})
})
