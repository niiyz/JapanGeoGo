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

	sampleType := os.Args[1]
	var name, ext string
	if sampleType == "geojson" {
		name = "googlemap/googlemap"
		ext = ".json"
	} else {
		name = "d3/d3"
		ext = ".topojson"
	}

	var pref, city string
	if os.Args[2] != "" {
		pref = os.Args[2]
	}
	if len(os.Args) == 4 {
		city = os.Args[3]
	}

	var jsonPath string
	var title string
	var filename string
	if city != "" {
		filename = city
		jsonPath = pref + "/" + filename + ext
		title = pref + city
	} else {
		filename = pref
		jsonPath = "47都道府県/" + filename + ext
		title = pref
	}

	jsonURL := sampleType + "/" + jsonPath
	fmt.Println(jsonURL)
	_, err := os.Stat(jsonURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := struct {
		Title    string
		JsonURL  string
		FileName string
	}{
		Title:    title,
		JsonURL:  jsonURL,
		FileName: filename,
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
	mux.Handle("/"+sampleType+"/", http.StripPrefix("/"+sampleType+"/", http.FileServer(http.Dir(sampleType+"/"))))

	http.ListenAndServe(":3000", mux)
}
