package main

import (
	"bufio"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func IndexRedirect(w http.ResponseWriter, r *http.Request) {
	DebugMsg("Index Redirect not implemented yet")
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	path := Conf.Templates.StaticDir + r.URL.Path

	f, err := os.Open(path)
	if err != nil {
		ErrorHtml(w)
	}
	SetHeaders(w, string(path))
	bufferedReader := bufio.NewReader(f)
	bufferedReader.WriteTo(w)
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
			content := SetTemplateContent()
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
