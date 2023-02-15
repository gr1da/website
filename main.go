package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type GlobalCtx struct {
	BaseTemplate        *template.Template
	GetCompiledHomepage func() *template.HTML
}

func NewGlobalCtx(basePath string) *GlobalCtx {
	tmpl, err := template.ParseFiles(basePath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var _compiledHomepage template.HTML = ""
	var _homepageNotBuilt bool = true

	return &GlobalCtx{
		BaseTemplate: tmpl,

		GetCompiledHomepage: func() *template.HTML {
			if _homepageNotBuilt {
				var builder strings.Builder
				tmpl.Execute(&builder, homepageContent)
				_compiledHomepage = template.HTML(builder.String())
				_homepageNotBuilt = false
			}
			return &_compiledHomepage
		},
	}
}

func homepage(ctx *GlobalCtx) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(*ctx.GetCompiledHomepage()))
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	ctx := NewGlobalCtx("templates/main_template.gohtml")
	http.HandleFunc("/", homepage(ctx))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
