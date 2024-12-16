package handlers

import (
	"html/template"
	"lb4/models"
	"net/http"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := models.ProcessPayment()
		if err != nil {
			renderError(w, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/payment.html"))
		t.Execute(w, nil)
	}
}
