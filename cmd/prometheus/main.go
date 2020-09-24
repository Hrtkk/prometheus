package main

import (
	"log"
	"net/http"

	"github.com/Hrtkk/prometheus/pkg/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	http.HandleFunc("/view/", api.MakeHandler(api.ViewHandler))
	http.HandleFunc("/edit/", api.MakeHandler(api.EditHandler))
	http.HandleFunc("/save/", api.MakeHandler(api.SaveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
