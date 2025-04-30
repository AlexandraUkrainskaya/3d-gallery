package main

import (
	"fmt"
	"net/http"
)

func (app *application) postWorkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "post a works")
}

func (app *application) showWorkHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
