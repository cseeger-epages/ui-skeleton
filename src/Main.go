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
	"flag"
	"net/http"
	"os"
)

var (
	Conf     config
	confFile *string
	TmplConf templateConfig
)

func init() {
	confFile = flag.String("c", "conf/ui.conf", "path to config file")
	flag.Parse()

	err := ParseConfig(*confFile, &Conf)
	if err != nil {
		Error("config parse error", err)
		os.Exit(1)
	}
	err = ParseConfig(Conf.Templates.TemplateConf, &TmplConf)
	if err != nil {
		Error("config parse error in templateconf", err)
		os.Exit(1)
	}

	InitLogger()
	Info("Basic Authentication", map[string]interface{}{"enabled": Conf.General.BasicAuth})
	Info("HTTP Strict Transport Security", map[string]interface{}{"enabled": Conf.Tls.Hsts})
}

func main() {
	router := NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(Conf.Templates.StaticDir + "/")))

	s, l, err := CreateServerAndListener(router, Conf.General.Listen, Conf.General.Port)
	if err != nil {
		Error("can not create server", err)
		os.Exit(1)
	}

	Info("starting server", map[string]interface{}{"ip": Conf.General.Listen, "port": Conf.General.Port})
	err = s.ServeTLS(l, Conf.Certs.Public, Conf.Certs.Private)
	Error("can't start server", err)
}
