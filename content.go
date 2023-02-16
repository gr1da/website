package main

import (
	"html/template"
	"strconv"
	"strings"
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

	homepageStatic = `
		<p>Work in progress, obviously.</p>
	`
	homepageContent = BaseContent{
		Title:       "Daniel's Homepage",
		Description: "Daniel's Personal Website (Homepage)",
		NavItems:    navItems,
		MainContent: template.HTML(homepageStatic),
		Year:        strconv.Itoa(time.Now().Year()),
	}
)

func getWeatherContent() string {
	//weather, err := NewRealWeather()
	//if err != nil {}

	//temp := weather.get().Hourly.Temperature_2m[len(weather.get().Hourly.Temperature_2m)-1]

	var bldr strings.Builder
	bldr.WriteString(`<h2>What it's like outside my window`)
	//bldr.WriteString(strconv.FormatFloat(temp, 'f', 1, 64))

	return bldr.String()
}
