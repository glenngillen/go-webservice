package api

import (
	"encoding/json"
	"io/ioutil"
)

type Page struct {
	Title string `json:"title"`
	Body  []byte `json:"body,omitempty"`
}

func (p *Page) MarshalJSON() ([]byte, error) {
	pj := &PageJSON{Title: p.Title, Body: string(p.Body[:])}
	data, error := json.Marshal(pj)
	return data, error
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}


type PageJSON struct {
	Title string `json:"title"`
	Body  string `json:"body,omitempty"`
}
