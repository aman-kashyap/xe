package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
	"time"
)

func tlsConfig(host string) error {
	now := time.Now()
	tpl := x509.Certificate{
		SerialNumber:          new(big.Int).SetInt64(0),
		Subject:               pkix.Name{CommonName: host},
		NotBefore:             now.Add(-24 * time.Hour).UTC(),
		NotAfter:              now.AddDate(1, 0, 0).UTC(),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		MaxPathLen:            1,
		IsCA:                  true,
		SubjectKeyId:          []byte{1, 2, 3, 4},
	}
	priv, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return err
	}
	der, err := x509.CreateCertificate(rand.Reader, &tpl, &tpl, &priv.PublicKey, priv)
	if err != nil {
		return err
	}
	crt, err := x509.ParseCertificate(der)
	if err != nil {
		return err
	}
	opts := x509.VerifyOptions{DNSName: host, Roots: x509.NewCertPool()}
	opts.Roots.AddCert(crt)
	_, err = crt.Verify(opts)
	return err
}

func main() {
	err := tlsConfig("localhost")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ok")
}
