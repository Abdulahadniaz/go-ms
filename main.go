package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		d, error := ioutil.ReadAll(r.Body)
		if error != nil {
			http.Error(rw, "Ooops something is not write", http.StatusBadRequest)
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye world!")
	})

	http.ListenAndServe(":9090", nil)
}
