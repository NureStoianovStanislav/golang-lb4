package models

func DeleteProduct(productID string) error {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", productID)
	return err
}
