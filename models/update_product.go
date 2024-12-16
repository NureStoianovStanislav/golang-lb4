package models

func UpdateProduct(productID, name, price string) error {
	_, err := db.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", name, price, productID)
	return err
}
