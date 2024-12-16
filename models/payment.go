package models

func ProcessPayment() error {
	_, err := db.Exec("DELETE FROM cart")
	return err
}
