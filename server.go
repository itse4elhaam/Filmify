package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type film struct{
	Title string
	Director string
}

func main() {
	fmt.Println("Hello World")

	// w is used to send response
	// "*http.Request" just contains info about the request
	handler := func(w http.ResponseWriter, r *http.Request) {
		// this writes a string to the w object that we take
		// * io.WriteString(w, "Writing to the server")

		// * io.WriteString(w, r.Method)


		// - this is how you connect frontend to backend in golang

		template := template.Must(template.ParseFiles("index.html"))

		// this is how mapping works, go is strongly typed language
		films := map[string][]film{
			"Films":{
				{Title: "Inception", Director : "Christopher Nolan"},
				{Title: "OpenHiemer", Director : "Christopher Nolan"},
				{Title: "Tenet", Director : "Christopher Nolan"},
			},
		}

		// first parameter is the write function and second is the data that we want to send down to the html 
		template.Execute(w, films)


	}


	// calls handler, when / used
	http.HandleFunc("/", handler)

	// listen and serve just starts the app
	// log fatal acts as a try catch here and logs out an error if needed
	log.Fatal(http.ListenAndServe(":8080", nil))
}
