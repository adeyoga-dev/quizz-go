package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	port := "5000"
	fmt.Printf("Starting Go Learning Server on http://0.0.0.0:%s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
