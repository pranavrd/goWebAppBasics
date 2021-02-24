package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fn := p.Title + ".txt"
	return ioutil.WriteFile(fn, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fn := title + ".txt"
	body, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "t1", Body: []byte("Sample text.")}
	p1.save()
	p2, _ := loadPage("t1")
	fmt.Println(string(p2.Body))
}
