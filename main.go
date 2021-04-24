package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	links := make(map[string]string)

	links["/git"] = "https://github.com/orsanawwad"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.URL.Path)
		link, exist := links[r.URL.Path]
		if !exist {
			fmt.Fprintf(w, "Not found")
			return
		}
		http.Redirect(w, r, link, http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "Req: %s %s\n", r.Host, r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":7777", nil))
}
