package main

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"log"
	"net/http"

	"ex01/handlers"
)

func main() {
	
	cert, err := tls.LoadX509KeyPair("certs/candy.tld/cert.pem", "certs/candy.tld/key.pem")
	if err != nil {
		log.Fatalf("failed to load server certificate and key: %v", err)
	}

	
	caCert, err := os.ReadFile("certs/minica.pem")
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	server := &http.Server{
		Addr:      ":3333",
		TLSConfig: tlsConfig,
		Handler:   http.HandlerFunc(handlers.BuyCandy),
	}

	log.Println("Server starting on port 3333 with TLS")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
