package main

import (
	"fmt"
	"github.com/tanopwan/rolele"
	middleware "github.com/tanopwan/rolele/middleware/http"
	"log"
	"net/http"
)

func registerHandle(pattern string, allowActions []rolele.Action) {
	http.Handle(pattern, middleware.GetRoleleMiddleware(allowActions, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, pattern)
	})))
}

func main() {
	addr := ":8888"
	server := &http.Server{
		Addr: addr,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	poAllowActions := []rolele.Action{{ID: 1, Name: "all_po"}, {ID: 2, Name: "get_po"}}
	soAllowActions := []rolele.Action{{ID: 3, Name: "all_so"}, {ID: 4, Name: "get_so"}}

	registerHandle("/api/po", poAllowActions)
	registerHandle("/api/so", soAllowActions)

	log.Printf("Starting http server at %s\n", addr)
	log.Fatal(server.ListenAndServe())
}
