package main

import (
	"fmt"
	"handler/function"
	"log"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":8080"),
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/", function.Handle)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
