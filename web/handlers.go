// Package web will provide all web interfacing support.
package web

import (
	"net/http"

	"github.com/CerealBoy/shortoftime/time"
	_ "github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Web struct {
	Id int `json:"id"`

	router *mux.Router
	time   *time.Time
}

func New(t *time.Time) *Web {
	w := new(Web)
	w.time = t

	return w
}

func (w *Web) Serve(port string) *Web {
	w.router = mux.NewRouter()
	w.router.HandleFunc("/", w.indexHandler)
	w.router.HandleFunc("/app", w.mainHandler)
	w.router.HandleFunc("/api/v1", w.apiHandler)
	w.router.HandleFunc("/res/{resource}", w.resourceHandler)

	http.ListenAndServe(port, w.router)

	return w
}

func (w *Web) indexHandler(o http.ResponseWriter, r *http.Request) {
	o.Write([]byte("Derp!\n"))
}

func (w *Web) mainHandler(o http.ResponseWriter, r *http.Request) {
	http.ServeFile(o, r, w.time.GetMainFile())
}

func (w *Web) resourceHandler(o http.ResponseWriter, r *http.Request) {
	http.ServeFile(o, r, w.time.GetResource(r))
}

func (w *Web) apiHandler(o http.ResponseWriter, r *http.Request) {
	//
}
