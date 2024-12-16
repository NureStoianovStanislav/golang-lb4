package models

type Product struct {
	ID    int
	Name  string
	Price float64
}

func GetAllProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Price)
		products = append(products, p)
	}
	return products, nil
}

func AddProduct(name string, price string) error {
	_, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", name, price)
	return err
}
