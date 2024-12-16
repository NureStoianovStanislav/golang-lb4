package handlers

import (
	"lb4/models"
	"net/http"
)

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		productID := r.FormValue("product_id")
		err := models.DeleteProduct(productID)
		if err != nil {
			renderError(w, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
