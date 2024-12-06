package main

import (
	"encoding/json"
	"go-quote-gen/quote"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		quote := quote.GenerateSingleRandomQuote(quote.DefaultQuotesProvider)

		jsonInBytes, err := json.MarshalIndent(quote, "", "  ")
		if err != nil {
			log.Printf("quote handler: error json marshal -> %v", err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInBytes)
	})

	s := http.Server{
		Handler: r,
		Addr:    ":1234",
	}

	log.Println("Server running...")
	s.ListenAndServe()
}
