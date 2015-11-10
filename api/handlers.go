package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
)

var (
	ren = render.New(render.Options{
		Directory:  "app/templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
	})
)

// Index ...
func Index(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	ren.HTML(rw, http.StatusOK, "index", nil)
}

// Bench ...
func Bench(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var r Request

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = r.validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
