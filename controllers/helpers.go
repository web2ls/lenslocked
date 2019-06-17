package controllers

import "net/http"
import "github.com/gorilla/schema"

func parseForms(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}
