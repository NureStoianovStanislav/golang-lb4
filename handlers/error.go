package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderError(w http.ResponseWriter, err error) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/error.html"))
	message := fmt.Sprintf("%s", err)
	t.Execute(w, map[string]string{"Message": message})
}
