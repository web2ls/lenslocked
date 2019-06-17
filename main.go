package main

import (
	"fmt"
	"net/http"

	"lenslocked/controllers"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<div>We are not found information for you(((((</div>")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers()

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Handle("/", staticController.Home).Methods("GET")
	router.Handle("/contact", staticController.Contact).Methods("GET")
	router.Handle("/faq", staticController.Faq).Methods("GET")
	router.HandleFunc("/signup", usersController.New).Methods("GET")
	router.HandleFunc("/signup", usersController.Create).Methods("POST")

	var h http.Handler = http.HandlerFunc(notFound)
	router.NotFoundHandler = h
	http.ListenAndServe(":3000", router)
}
