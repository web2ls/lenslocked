package main

import (
	"fmt"
	"net/http"

	"lenslocked/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
	// if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
	// 	panic(err)
	// }
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
	// if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
	// 	panic(err)
	// }
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
	// if err := faqView.Template.ExecuteTemplate(w, faqView.Layout, nil); err != nil {
	// 	panic(err)
	// }
}

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
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/faq", faq)

	var h http.Handler = http.HandlerFunc(notFound)
	router.NotFoundHandler = h
	http.ListenAndServe(":3000", router)
}
