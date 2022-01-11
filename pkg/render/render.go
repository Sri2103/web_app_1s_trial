package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, html string) {

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("error getting template Cache:", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + html)

	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error in Parsling Template :", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err

		}

		matches, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			return myCache, err

		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./template/*.layout.html")
			if err != nil {
				return myCache, err

			}
		}

		myCache[name] = ts
	}

	return myCache, err

}
