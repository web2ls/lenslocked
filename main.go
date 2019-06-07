package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>To get in touch please send me email</h1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<div><h3>FAQ</h3><p>This is how our web service is work</p></div>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<div>We are not found information for you(((((</div>")
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")

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
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	var h http.Handler = http.HandlerFunc(notFound)
	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}
