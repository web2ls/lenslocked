package controllers

import "lenslocked/views"

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
		Faq:     views.NewView("secondary", "views/static/faq.gohtml"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}
