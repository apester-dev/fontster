package font

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// HandleCSS returns http.Handler for the fonts API.
func HandleCSS(tmpl *template.Template, url string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		familyParam := r.URL.Query().Get("family")
		if familyParam == "" {
			fmt.Fprintf(w, "must provide family e.g /css?family=Lato")
			return
		}
		fonts := Parse(familyParam, url)
		if pusher, ok := w.(http.Pusher); ok {
			for _, f := range fonts {
				if err := pusher.Push(f.Source(), nil); err != nil {
					log.Printf("could not push: %v", err)
				}
			}
		}
		w.Header().Set("Content-Type", "text/css")
		tmpl.Execute(w, fonts)
	})
}
