package models

type Cart struct {
	Items []CartItem
	Total float64
}

type CartItem struct {
	ID          int
	ProductName string
	Quantity    int
	Price       float64
}

func GetCart() (*Cart, error) {
	rows, err := db.Query(`
        SELECT c.id, p.name, c.quantity, p.price
        FROM cart c
        JOIN products p ON c.product_id = p.id
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CartItem
	var total float64
	for rows.Next() {
		var item CartItem
		rows.Scan(&item.ID, &item.ProductName, &item.Quantity, &item.Price)
		items = append(items, item)
		total += item.Price * float64(item.Quantity)
	}
	cart := Cart{
		Items: items,
		Total: total,
	}
	return &cart, nil
}

func AddToCart(productID, quantity string) error {
	_, err := db.Exec(`
		INSERT INTO cart (product_id, quantity)
        VALUES ($1, $2)
        ON CONFLICT (product_id) DO
        UPDATE SET quantity = cart.quantity + EXCLUDED.quantity
        `,
		productID, quantity,
	)
	return err
}
