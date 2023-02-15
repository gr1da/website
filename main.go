package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

const github = "https://github.com/gr1da"

type NavItem struct {
	Url  string
	Name string
}

var NavList = []NavItem{
	{"/", "Home"},
	{"/blog", "Blog"},
	{github, "Github"},
}

var mainTemplate *template.Template
var blogPageTemplate *template.Template

func setup() {
	mt, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf(err.Error())
	}
	mainTemplate = mt

	bpt, err := template.ParseFiles("blogpage_template.html")
	if err != nil {
		log.Fatalf(err.Error())
	}
	blogPageTemplate = bpt
}

func main() {
	setup()
	http.HandleFunc("/", homepage)
	http.HandleFunc("/blog", blog)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type BasicPage struct {
	Title       string
	Description string
	Navigation  []NavItem
	Content     template.HTML
}

type BlogPageEntry struct {
	Url   string
	Date  string
	Title string
}

type BlogPage struct {
	List []BlogPageEntry
}

func homepage(w http.ResponseWriter, r *http.Request) {
	page := BasicPage{
		"Daniel's Website",
		"Daniel's personal homepage",
		NavList,
		"<p>Work in progress, obviously.</p>",
	}
	mainTemplate.Execute(w, page)
}

func blog(w http.ResponseWriter, r *http.Request) {
	var contentBldr bytes.Buffer
	blogPageTemplate.Execute(&contentBldr,
		BlogPage{[]BlogPageEntry{{"#", "nodate", "nothing"}}})

	mainTemplate.Execute(w, BasicPage{
		"Daniel's Blog",
		"Daniel's personal blog",
		NavList,
		template.HTML(contentBldr.String()),
	})
}
