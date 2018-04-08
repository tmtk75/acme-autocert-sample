package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	d := flag.String("d", "acme-test.tmtk.net", "domain")
	flag.Parse()

	domains := []string{
		*d,
	}
	mgr := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domains...),
		Cache:      autocert.DirCache("certs"), // to store certs
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world\n"))
		log.Printf("%v", r)
	})

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: mgr.GetCertificate,
		},
	}

	go http.ListenAndServe(":http", mgr.HTTPHandler(nil))
	log.Printf("start listening at :http")

	log.Fatal(server.ListenAndServeTLS("", "")) // Key and cert provided by Let's Encrypt
}
