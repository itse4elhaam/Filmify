package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type film struct {
	Title    string
	Director string
	Genre    string
}

func main() {
	fmt.Println("Hello World")

	// w is used to send response
	// "*http.Request" just contains info about the request
	handler := func(w http.ResponseWriter, r *http.Request) {
		// this writes a string to the w object that we take
		// io.WriteString(w, "Writing to the server")

		// io.WriteString(w, r.Method)

		// - this is how you connect frontend to backend in golang

		template := template.Must(template.ParseFiles("index.html"))

		// this is how mapping works, go is strongly typed language
		films := map[string][]film{
			"Films": {
				{Title: "Inception", Director: "Christopher Nolan", Genre: "Sci-Fi"},
				{Title: "OpenHiemer", Director: "Christopher Nolan", Genre: "History"},
				{Title: "Tenet", Director: "Christopher Nolan", Genre: "SciFi"},
			},
		}

		// first parameter is the write function and second is the data that we want to send down to the html
		template.Execute(w, films)

	}

	addFilmHandler := func(w http.ResponseWriter, r *http.Request) {
		// log.Print("got HTMX req")
		// log.Print(r.Header.Get("Hx-Request")) // will be set to true if HTMX sent the request

		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		genre := r.PostFormValue("genre")

		// this uses template fragments in go, it can be done by creating a block in the index.html file with a name
		// the block is in the loop and uses the same naming convention so its easy to pass the data from here
		// - the other method is tiresome but is implemented down from line 51 onwards

		template := template.Must(template.ParseFiles("index.html"))

		template.ExecuteTemplate(w, "film-info-element", film{Title: title, Director: director, Genre: genre})

		// %s is replaced with the var we pass at the end
		// filmInfoHtml := fmt.Sprintf("<div class='film-info'><h3>%s</h3><p>-</p><h6>%s</h6><p>%s</p></div>", title, director, genre)

		// _ is the err this function with return
		// filmInfoHtmlTemplate, _ := template.New("filmInfo").Parse(filmInfoHtml)

		// rendering them, no data needed so nil used
		// filmInfoHtmlTemplate.Execute(w, nil)
	}

	// calls handler, when / used
	http.HandleFunc("/", handler)
	http.HandleFunc("/add-film/", addFilmHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// listen and serve just starts the app
	// log fatal acts as a try catch here and logs out an error if needed
	log.Fatal(http.ListenAndServe(":8080", nil))
}
