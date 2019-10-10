package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/qmerce/fontster/pkg/font"
)

var (
	address     = flag.String("address", "localhost:3000", "listen address")
	idleTimeout = flag.Duration("timeout", 620*time.Second, "IdleTimeout for the HTTP server")
	fontsSource = flag.String("fonts-source", "", "URL to where the fonts are hosted")
	fontsDir    = flag.String("fonts-dir", "./fonts", "path in the filesystem containing fonts")
	tlsCert     = flag.String("tls-cert", "", "path to TLS certificate")
	tlsKey      = flag.String("tls-key", "", "path to TLS key")
)

// Options to configure fontster.
type Options struct {
	Address     string
	LocalDir    string
	URL         string
	Template    *template.Template
	IdleTimeout time.Duration
	CertFile    string
	KeyFile     string
}

func main() {
	flag.Parse()

	opts := Options{
		Address:     *address,
		LocalDir:    *fontsDir,
		URL:         *fontsSource,
		Template:    font.CSSTemplate(),
		IdleTimeout: *idleTimeout,
		CertFile:    *tlsCert,
		KeyFile:     *tlsKey,
	}

	// Serve fonts from local filesystem.
	if opts.URL == "" {
		opts.URL = "/fonts"
		http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir(opts.LocalDir))))
	}
	http.Handle("/css", font.HandleCSS(opts.Template, opts.URL))

	srv := &http.Server{
		Addr:        opts.Address,
		IdleTimeout: opts.IdleTimeout,
	}
	idleConns := drain(srv)

	var err error
	log.Printf("Started server listen-address=%q", opts.Address)
	if opts.CertFile != "" && opts.KeyFile != "" {
		err = srv.ListenAndServeTLS(opts.CertFile, opts.KeyFile)
	} else {
		err = srv.ListenAndServe()
	}

	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
	<-idleConns
}

func drain(srv *http.Server) <-chan struct{} {
	idleConns := make(chan struct{})
	go func() {
		q := make(chan os.Signal, 1)
		signal.Notify(q, syscall.SIGTERM, os.Interrupt)
		<-q
		log.Println("Got shutdown signal")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("Could not shutdown gracefully: %v", err)
		}
		close(idleConns)
	}()
	return idleConns
}
