package main

import (
	"fmt"
	"github.com/shokishimo/OneTap/server"
	"log"
	"net/http"
)

func main() {
	srvMux := server.ServeMux()
	fmt.Println("Server is running")
	if err := http.ListenAndServe(":3000", srvMux); err != nil {
		log.Fatal("Error in running the server")
		fmt.Printf("e: %v\n", err)
	}
}
