package main

import (
	"flag"
	"log"
	"time"

	"github.com/qmerce/fontster/pkg/font"
	"github.com/qmerce/fontster/pkg/http"
)

// Command line flags config.
var (
	address     = flag.String("address", "localhost:3000", "listen address")
	idleTimeout = flag.Duration("timeout", 620*time.Second, "IdleTimeout for the HTTP server")
	fontsSource = flag.String("fonts-source", "", "URL to where the fonts are hosted")
	fontsDir    = flag.String("fonts-dir", "./fonts", "path in the filesystem containing fonts")
	tlsCert     = flag.String("tls-cert", "../box/certs/apester.local.com.crt", "path to TLS certificate")
	tlsKey      = flag.String("tls-key", "../box/certs/apester.local.com.key", "path to TLS key")
)

func main() {
	flag.Parse()
	err := http.FontsServer(http.Options{
		Address:     *address,
		LocalDir:    *fontsDir,
		URL:         *fontsSource,
		Template:    font.CSSTemplate(),
		IdleTimeout: *idleTimeout,
		CertFile:    *tlsCert,
		KeyFile:     *tlsKey,
	})

	if err != nil {
		log.Fatal(err)
	}
}
