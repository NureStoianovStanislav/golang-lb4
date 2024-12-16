package models

func UpdateCartItem(cartItemID, quantity string) error {
	_, err := db.Exec("UPDATE cart SET quantity = $1 WHERE id = $2", quantity, cartItemID)
	return err
}
