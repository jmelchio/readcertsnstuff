package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	keyBytes, err := ioutil.ReadFile("x509.key")
	check(err)

	certBytes, err := ioutil.ReadFile("x509.cert")
	check(err)

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	check(err)

	fmt.Println(cert.PrivateKey)
	fmt.Println(cert.SignedCertificateTimestamps)
}
