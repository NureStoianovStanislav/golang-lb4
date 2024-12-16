package models

func DeleteCartItem(cartItemID string) error {
	_, err := db.Exec("DELETE FROM cart WHERE id = $1", cartItemID)
	return err
}
