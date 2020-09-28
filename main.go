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

	config := gittp.ServerConfig{
		Path: *dirFlag,
	}

	handle, _ := gittp.NewGitServer(config)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), handle))
}