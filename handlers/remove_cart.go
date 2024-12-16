package handlers

import (
	"lb4/models"
	"net/http"
)

func DeleteCartItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		cartItemID := r.FormValue("cart_item_id")
		err := models.DeleteCartItem(cartItemID)
		if err != nil {
			renderError(w, err)
			return
		}
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}
