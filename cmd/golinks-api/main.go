package main

import (
	"flag"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goccy/go-json"

	log "github.com/sirupsen/logrus"

	"github.com/orsanawwad/golinks/pkg/db"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

type CreateLinkRequest struct {
	db.Link `json:"link"`
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	golinkdb := db.New()

	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/", func(r chi.Router) {
		r.Route("/{shortUrlKey}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				if shortUrlKey := chi.URLParam(r, "shortUrlKey"); shortUrlKey != "" {
					link := golinkdb.GetLink(shortUrlKey)
					if link == "" {
						w.Write([]byte("Not found"))
						return
					}
					http.Redirect(w, r, link, http.StatusTemporaryRedirect)
				}
			})
		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			w.Write([]byte("Root."))
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Body == nil {
				http.Error(w, "Please send a request body", 400)
				return
			}
			var link CreateLinkRequest
			err := json.NewDecoder(r.Body).Decode(&link)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			_, err = url.ParseRequestURI(link.URL)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			golinkdb.CreateLink(&link.Link)
			w.Write([]byte("OK"))
		})
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	http.ListenAndServe(":7777", r)

}
