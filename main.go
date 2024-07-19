package main

import (
	handler "exchange/handlers"

	"fmt"
	"net/http"
	"os"
	"time"
)

var port = ":8010"

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		port = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         port,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/", http.HandlerFunc(handler.MainHandler))
	mux.Handle("/crypto", http.HandlerFunc(handler.CryptoHandler))
	mux.Handle("/calculate", http.HandlerFunc(handler.CalcHandler))
	mux.Handle("/calculate/ok", http.HandlerFunc(handler.OkHandler))

	fmt.Println("Ready to Serve HTTP to PORT", port)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
