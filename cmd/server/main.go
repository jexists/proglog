package main

import (
	"fmt"
	"log"
	"proj_proglog/internal/server"
)

func main() {
	fmt.Print("????")
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
