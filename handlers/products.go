package handlers

import (
	"html/template"
	"lb4/models"
	"net/http"
)

func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		renderError(w, err)
		return
	}
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/catalog.html"))
	t.Execute(w, products)
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		price := r.FormValue("price")
		err := models.AddProduct(name, price)
		if err != nil {
			renderError(w, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/add_product.html"))
		t.Execute(w, nil)
	}
}
