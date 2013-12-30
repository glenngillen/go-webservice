package api

// Import formatting and IO libraries
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
)

// Load pages too
func loadPages() (pages []*Page, err error) {
	err = nil
	files, _ := filepath.Glob("/Users/glenngillen/Development/go-webservice/*.txt")
	re := regexp.MustCompile("([^/]+).txt$")
	title := re.FindStringSubmatch(files[0])[1]
	page, err := loadPage(title)
	if err != nil {
		return nil, err
	}
	pages = append(pages, page)
	return pages, nil
}

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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
                SaveHandler(w, r)
	case "GET":
		pages, err := loadPages()
		if err != nil {
			http.NotFound(w, r)
			return
		}
		data, _ := json.Marshal(pages)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(data)
	}

}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.NotFound(w, r)
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

func SaveHandler(w http.ResponseWriter, r *http.Request) {
        var page Page
        jsonBody := make([]byte, r.ContentLength)
        _, err := r.Body.Read(jsonBody)
        json.Unmarshal(jsonBody, &page)
        err1 := page.Save()
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+page.Title, http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/pages", IndexHandler)
	http.HandleFunc("/pages/", makeHandler(ViewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", SaveHandler)
	http.ListenAndServe(":8080", nil)
}
