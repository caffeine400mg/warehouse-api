package warehouse

type Warehouse struct {
	Products []Product
	Profit   int
}

type Product struct {
	ID       int      `json:"product_id"`
	Name     string   `json:"product_name"`
	Price    int      `json:"product_price"`
	Quantity int      `json:"product_quantity"`
	Category Category `json:"product_category"`
}

func (w *Warehouse) CreateProduct(p Product) error {
	if w == nil {
		return ErrNilWarehouse
	}

	for i := range w.Products {
		if w.Products[i].Name == p.Name {
			return ErrProductAlreadyExist
		}
	}

	w.Products = append(w.Products, p)
	return nil
}

func (w *Warehouse) AddProduct(id int, amount int) (Product, error) {
	if w == nil {
		return Product{}, ErrNilWarehouse
	}

	if err := EditAmountValidate(amount); err != nil {
		return Product{}, err
	}

	for i, v := range w.Products {
		if v.ID == id {
			w.Products[i].Quantity += amount
			return w.Products[i], nil
		}
	}
	return Product{}, ErrProductNotFound
}

func (w *Warehouse) RemoveProduct(id int, amount int) (Product, error) {
	if w == nil {
		return Product{}, ErrNilWarehouse
	}
	if amount <= 0 {
		return Product{}, ErrInvalidAmount
	}
	for i := range w.Products {
		if w.Products[i].ID == id {
			if w.Products[i].Quantity < amount {
				return Product{}, ErrNotEnoughProducts
			}
			w.Products[i].Quantity -= amount
			return w.Products[i], nil
		}
	}
	return Product{}, ErrProductNotFound
}

func (w *Warehouse) SellProduct(id int, amount int) (Product, error) {
	if w == nil {
		return Product{}, ErrNilWarehouse
	}
	if amount <= 0 {
		return Product{}, ErrInvalidAmount
	}
	for i, v := range w.Products {
		if w.Products[i].ID == id {
			if amount > w.Products[i].Quantity {
				return Product{}, ErrNotEnoughProducts
			}
			w.Products[i].Quantity -= amount
			w.Profit += v.Price * amount
			return w.Products[i], nil
		}
	}
	return Product{}, ErrProductNotFound
}

func ShowAllProducts() {}
func ShowProfit()      {}
