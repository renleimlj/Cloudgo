package main

import (
    "os"
    "Cloudgo/service"
    flag "github.com/spf13/pflag"
)

const (
    PORT string = "8080"
)

func main() {
	// To get port, if not given, it's 8080
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT
    }
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
	}
	// To call the func defined in Server.go
    service.NewServer(port)
}