package main

import (
	"crypto/tls"
	"crypto/x509"
)

// helper function to create a cert template with a serial number and other required fields
func CertTemplate() (*x509.Certificate, error) {
    // generate a random serial number (a real cert authority would have some logic behind this)
    serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
    serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
    if err != nil {
        return nil, errors.New("failed to generate serial number: " + err.Error())
    }

    tmpl := x509.Certificate{
        SerialNumber:          serialNumber,
        Subject:               pkix.Name{Organization: []string{"Yhat, Inc."}},
        SignatureAlgorithm:    x509.SHA256WithRSA,
        NotBefore:             time.Now(),
        NotAfter:              time.Now().Add(time.Hour), // valid for an hour
        BasicConstraintsValid: true,
    }
    return &tmpl, nil
}

// create a key-pair for the server

servKey, err := rsa.GenerateKey(rand.Reader, 2048)
if err != nil {
	log.Fatalf("generating random key: %v", err)
}

// create a template for the server
servCertTmpl, err := CertTemplate()
if err != nil {
	log.Fatalf("creating cert template: %v", err)
}

servCertTmpl.KeyUsage = x509.KeyUsageDigitalSignature
servCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
servCertTmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}


// create a certificate which wraps the server's public key, sign it with the root private key
_, servCertPEM, err := CreateCert(servCertTmpl, rootCert, &servKey.PublicKey, rootKey)
if err != nil {
	log.Fatalf("error creating cert: %v", err)
}

// provide the private key and the cert
servKeyPEM := pem.EncodeToMemory(&pem.Block{
	Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(servKey),
})
servTLSCert, err := tls.X509KeyPair(servCertPEM, servKeyPEM)
if err != nil {
	log.Fatalf("invalid key pair: %v", err)
}
// create another test server and use the certificate
s = httptest.NewUnstartedServer(http.HandlerFunc(ok))
s.TLS = &tls.Config{
	Certificates: []tls.Certificate{servTLSCert},
}

// create a pool of trusted certs
certPool := x509.NewCertPool()
certPool.AppendCertsFromPEM(rootCertPEM)

// configure a client to use trust those certificates
client := &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: certPool},
	},
}

s.StartTLS()
resp, err := client.Get(s.URL)
s.Close()
if err != nil {
	log.Fatalf("could not make GET request: %v", err)
}
dump, err := httputil.DumpResponse(resp, true)
if err != nil {
	log.Fatalf("could not dump response: %v", err)
}
fmt.Printf("%s\n", dump)