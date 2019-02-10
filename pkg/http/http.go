package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"

	"github.com/qmerce/fontster/pkg/font"
)

type Options struct {
	Address     string
	LocalDir    string
	URL         string
	Template    *template.Template
	IdleTimeout time.Duration
	CertFile    string
	KeyFile     string
}

func FontsServer(opts Options) error {
	if opts.URL == "" {
		opts.URL = "/fonts"
	}

	// routes
	http.Handle("/css", fontsHandler(opts.Template, opts.URL))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir(opts.LocalDir))))

	// custom server to control shutdown and timeouts
	srv := &http.Server{Addr: opts.Address, IdleTimeout: opts.IdleTimeout}

	// handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		q := make(chan os.Signal)
		signal.Notify(q, syscall.SIGTERM)
		<-q

		// turn off keepalives before shutdown
		srv.SetKeepAlivesEnabled(false)
		log.Println("shutting down...")
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("could not shutdown gracefully: %v", err)
		}
		cancel()
	}()

	log.Printf("started server on %s", opts.Address)

	var err error
	if opts.CertFile != "" && opts.KeyFile != "" {
		err = srv.ListenAndServeTLS(opts.CertFile, opts.KeyFile)
	} else {
		err = srv.ListenAndServe()

	}

	if err != http.ErrServerClosed {
		return err
	}

	<-ctx.Done()
	return nil
}

func fontsHandler(tmpl *template.Template, url string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		familyParam := r.URL.Query().Get("family")
		if familyParam == "" {
			fmt.Fprintf(w, "must provide family e.g /css?family=Lato")
			return
		}

		fonts := font.Parse(familyParam, url)
		if pusher, ok := w.(http.Pusher); ok {
			for _, f := range fonts {
				err := pusher.Push(fmt.Sprintf("%s/%s/%s", url, f.Family, f.Filename), nil)
				if err != nil {
					log.Printf("could not push: %v", err)
				}
			}
		}

		w.Header().Set("Content-Type", "text/css")
		tmpl.Execute(w, fonts)
	})
}
