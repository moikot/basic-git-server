package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/adamveld12/gittp"
)

func main() {
	dirFlag := flag.String("d", "./", "Path to the repositories.")
	httpPort := flag.Int("p", 8080, "Port to serve on.")

	flag.Parse()

	addr := fmt.Sprintf(":%d", *httpPort)
	log.Println("Starting server on ", addr)

	config := gittp.ServerConfig{
		Path: *dirFlag,
	}

	handle, _ := gittp.NewGitServer(config)
	log.Fatal(http.ListenAndServe(addr, handle))
}