package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"
)

func CacheControlHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var age time.Duration
			ext := filepath.Ext(r.URL.String())

			switch ext {
			case ".rss", ".atom":
				age = time.Hour / time.Second
			case ".css", ".js":
				age = (time.Hour * 24 * 365) / time.Second
			case ".jpg", ".jpeg", ".gif", ".png", ".ico", ".cur", ".gz", ".svg", ".svgz", ".mp4", ".ogg", ".ogv", ".webm", ".htc":
				age = (time.Hour * 24 * 30) / time.Second
			default:
				age = 0
			}

			if age > 0 {
				w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, public, must-revalidate, proxy-revalidate", age))
			}
		}

		h.ServeHTTP(w, r)
	})
}
