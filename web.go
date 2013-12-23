package main

// Import formatting and IO libraries
import (
        "encoding/json"
	"io/ioutil"
	"net/http"
        "regexp"
)

// Define a structure to hold our page
type Page struct {
	Title string `json:"title"`
	Body  []byte `json:"body,omitempty"`
}

type PageJSON struct {
        Title string `json:"title"`
        Body  string  `json:"body,omitempty"`
}

func (p *Page) MarshalJSON() ([]byte, error) {
  pj := &PageJSON{Title: p.Title, Body: string(p.Body[:])}
  data, error := json.Marshal(pj)
  return data, error
}

// Add a save method to our Page struct so we can persist our data
// This method's signature reads: "This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error."
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load pages too
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var validPath = regexp.MustCompile("^/pages/([a-zA-Z0-9]+)$")

func renderJSON(w http.ResponseWriter, tmpl string, p *Page) {
        data, _ := json.Marshal(p)
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.Write(data)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
                if m == nil {
                      http.NotFound(w, r)
                      return
                }
                fn(w, r, m[2])
	}
}
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderJSON(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderJSON(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/pages/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":8080", nil)
}
