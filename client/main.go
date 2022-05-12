package main

import (
	"crypto/tls"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)

	os.Setenv("SSL_CERT_FILE", "server.crt")

	conf := &tls.Config{}

	conn, err := tls.Dial("tcp", "localhost:443", conf)

	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	println("tls conn opened")

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println("here is what we got")
	println(string(buf[:n]))
}
