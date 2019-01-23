package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	if len(os.Args) < 3 {
		err := fmt.Errorf("usage: %s [keyfile_name] [certfile_name]", os.Args[0])
		check(err)
	}

	x509keyFile := os.Args[1]
	x509certFile := os.Args[2]

	keyBytes, err := ioutil.ReadFile(x509keyFile)
	check(err)

	certBytes, err := ioutil.ReadFile(x509certFile)
	check(err)

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	check(err)

	fmt.Println(cert.PrivateKey)
	fmt.Println(cert.SignedCertificateTimestamps)
}
