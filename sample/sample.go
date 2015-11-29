package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var templates = template.Must(template.New("Sample").ParseGlob("sample/templates/**/*.html"))

var errorTemplate = `
<html>
	<body>
		<h1>Error rendering template %s</h1>
		<p>%s</p>
	</body>
</html>`

func main() {
	mux := http.NewServeMux()

	name := "googlemap/googlemap"

	var pref, city string
	if os.Args[1] != "" {
		pref = os.Args[1]
	}
	if len(os.Args) == 3 && os.Args[2] != "" {
		city = os.Args[2]
	}

	var jsonPath string
	var title string
	if city != "" {
		jsonPath = pref + "/" + city + ".json"
		title = pref + city
	} else {
		jsonPath = "47都道府県/" + pref + ".json"
		title = pref
	}

	jsonURL := "geojson/" + jsonPath

	_, err := os.Stat(jsonURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := struct {
		Title   string
		JsonURL string
	}{
		Title:   title,
		JsonURL: jsonURL,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, name, data)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf(errorTemplate, name, err),
				http.StatusInternalServerError,
			)
		}
	})
	mux.Handle("/geojson/", http.StripPrefix("/geojson/", http.FileServer(http.Dir("geojson/"))))

	http.ListenAndServe(":3000", mux)
}
