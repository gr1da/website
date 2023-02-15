package main

import (
	"html/template"
	"strconv"
	"time"
)

type (
	NavItem struct {
		Url  string
		Name string
	}

	BaseContent struct {
		Title       string
		Description string
		NavItems    []NavItem
		MainContent template.HTML
		Year        string
	}
)

const (
	github = "https://github.com/gr1da"
)

var (
	navItems = []NavItem{
		{"/", "Home"},
		{"/blog", "Blog"},
		{github, "GitHub"},
	}

	homepageMain = `
		<p>Work in progress, obviously.</p>
	`
	homepageContent = BaseContent{
		Title:       "Daniel's Homepage",
		Description: "Daniel's Personal Website (Homepage)",
		NavItems:    navItems,
		MainContent: template.HTML(homepageMain),
		Year:        strconv.Itoa(time.Now().Year()),
	}
)
