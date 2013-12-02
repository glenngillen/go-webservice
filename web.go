package main

// Import formatting and IO libraries
import (
	"fmt"
	"io/ioutil"
        "net/http"
)

// Define a structure to hold our page
type Page struct {
	Title string
	Body  []byte // IO libs expect a byte slice rather than a string
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
        p, _ := loadPage(title)
            fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
            }


func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
        http.HandleFunc("/view/", viewHandler)
            http.ListenAndServe(":8080", nil)
}
