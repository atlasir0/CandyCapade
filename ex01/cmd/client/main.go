package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type CandyRequest struct {
	CandyType string `json:"candy_type"`
	Count     int    `json:"count"`
	Money     int    `json:"money"`
}

func main() {
	candyType := flag.String("k", "", "Two-letter abbreviation for the candy type")
	count := flag.Int("c", 0, "Count of candy to buy")
	money := flag.Int("m", 0, "Amount of money given to the machine")
	insecure := flag.Bool("insecure", false, "Ignore certificate validation")
	addr := flag.String("addr", "https://127.0.0.1:3333", "Address of the candy server")

	certFile := flag.String("cert", "certs/client/cert.pem", "Client certificate file")
	keyFile := flag.String("key", "certs/client/key.pem", "Client key file")
	caFile := flag.String("ca", "certs/minica.pem", "CA certificate file")

	flag.Parse()

	if *candyType == "" || *count == 0 || *money == 0 {
		log.Fatalf("All flags -k, -c and -m must be provided and non-zero")
	}

	reqBody := &CandyRequest{
		CandyType: *candyType,
		Count:     *count,
		Money:     *money,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}


	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("Failed to load client certificate and key: %v", err)
	}

	caCert, err := os.ReadFile(*caFile)
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: *insecure,
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", *addr+"/buy_candy", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println(string(body))
}
