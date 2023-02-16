package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type GlobalCtx struct {
	baseTemplate        *template.Template
	getCompiledHomepage func() *template.HTML
	requestBuildHome    func()
}

func NewGlobalCtx(basePath string) *GlobalCtx {
	_baseTemplate, err := template.ParseFiles(basePath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var _compiledHomepage template.HTML = ""
	var _homepageNotBuilt bool = true

	_getCompiledHomepage := func() *template.HTML {
		if _homepageNotBuilt {
			homepageContent.MainContent += template.HTML(getWeatherContent())

			var builder strings.Builder
			_baseTemplate.Execute(&builder, homepageContent)
			_compiledHomepage = template.HTML(builder.String())
			_homepageNotBuilt = false
		}
		return &_compiledHomepage
	}

	_requestBuildHome := func() {
		_homepageNotBuilt = false
	}

	return &GlobalCtx{
		baseTemplate:        _baseTemplate,
		getCompiledHomepage: _getCompiledHomepage,
		requestBuildHome:    _requestBuildHome,
	}
}

func homepage(ctx *GlobalCtx) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(*ctx.getCompiledHomepage()))
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	ctx := NewGlobalCtx("templates/main_template.gohtml")
	http.HandleFunc("/", homepage(ctx))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
