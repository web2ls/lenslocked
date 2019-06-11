package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"

	"github.com/gorilla/mux"
)

var homeView *View
var contactView *View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<div>We are not found information for you(((((</div>")
}

func main() {
	var err error
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	// homeTemplate, err = template.ParseFiles(
	// 	"views/home.gohtml",
	// 	"views/layouts/footer.gohtml",
	// )
	// contactTemplate, err = template.ParseFiles(
	// 	"views/contact.gohtml",
	// 	"views/layouts/footer.gohtml",
	// )

	if err != nil {
		panic(err)
	}

	// httprouter use
	// router := httprouter.New()
	// router.GET("/", home)
	// router.GET("/contact", contact)
	// router.GET("/faq", faq)
	// var h http.Handler = http.HandlerFunc(notFound)
	// router.NotFound = h

	// mux router use
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)

	var h http.Handler = http.HandlerFunc(notFound)
	router.NotFoundHandler = h
	http.ListenAndServe(":3000", router)
}
