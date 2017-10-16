package main

import (
	"html/template"
	"net/http"
	"strings"
)

func IndexRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Path+"index", 301)
}

func Index(w http.ResponseWriter, r *http.Request) {

	// set correct path and suffix
	path := r.URL.Path
	if path != "/" {
		CheckHtmlSuffix(&path)
	} else {
		path = "index.html"
	}

	folder := Conf.Templates.TemplatesDir
	if !strings.HasSuffix(path, ".html") {
		folder = Conf.Templates.StaticDir
	}
	path = folder + path

	// exclude tmplbase
	if strings.Contains(path, Conf.Templates.BaseDir) {
		ErrorHtml(w)
		return
	}

	data, err := GetTemplateData(path)
	if err == nil {
		contentType := SetHeaders(w, string(path))
		if contentType == "text/html" {
			content := SetTemplateContent(r)
			tmpl, err := template.New("tmpl").Parse(data)
			if err != nil {
				Error("error while creating template", err)
				ErrorHtml(w)
				return
			}
			tmpl.Execute(w, content)
		}
	} else {
		Error("template parse error", err)
		ErrorHtml(w)
	}
}
