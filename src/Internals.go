/*
	GOLANG UI Skeleton

	Copyright (C) 2017 Carsten Seeger

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.

	@author Carsten Seeger
	@copyright Copyright (C) 2017 Carsten Seeger
	@license http://www.gnu.org/licenses/gpl-3.0 GNU General Public License 3
	@link https://github.com/cseeger-epages/rest-api-go-skeleton
*/

package main

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// concat tpl files to create the template string
func GetTemplateData(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	var template string

	for _, v := range Conf.Templates.BaseTemplatesBc {
		tmplData, err := ioutil.ReadFile(Conf.Templates.BaseDir + "/" + v)
		if err != nil {
			return "", err
		}
		template = template + string(tmplData)
	}

	template = template + string(data)

	for _, v := range Conf.Templates.BaseTemplatesAc {
		tmplData, err := ioutil.ReadFile(Conf.Templates.BaseDir + "/" + v)
		if err != nil {
			return "", err
		}
		template = template + string(tmplData)
	}

	return template, nil
}

// set content from file
func SetTemplateContent() Template {
	content := Template{
		TmplConf.General.Title,
		TmplConf.General.Branding,
		TmplConf.Navs,
		TmplConf.General.Footer,
	}
	return content
}

// fix missing .html suffix :)
func CheckHtmlSuffix(path *string) {
	var html string = ".html"
	for k, _ := range ContentTypes {
		if strings.HasSuffix(*path, k) {
			html = ""
			break
		}
	}
	*path += html
}

func SetHeaders(w http.ResponseWriter, path string) string {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	var contentType string = "text/plain"

	for k, v := range ContentTypes {
		if strings.HasSuffix(path, k) {
			contentType = v
			break
		}
	}
	w.Header().Add("Content-Type", contentType)
	return contentType
}

/*
  write 404 error page back
*/
func ErrorHtml(w http.ResponseWriter) {
	w.WriteHeader(404)
	path := Conf.Templates.ErrorPage
	f, _ := os.Open(path)
	bufferedReader := bufio.NewReader(f)
	bufferedReader.WriteTo(w)
}
