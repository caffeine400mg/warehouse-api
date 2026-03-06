package warehouse

type Category int

const (
	Food Category = iota + 1
	Tools
	Clothes
)

func ShowCategory(p Product) (error, string) {
	switch p.Category {
	case 1:
		return nil, "Food"
	case 2:
		return nil, "Tools"
	case 3:
		return nil, "Clothes"
	default:
		return ErrInvalidCategory, ""
	}
}

func DecodeCategory(tempCategory int) (error, Category) {
	var category Category
	if tempCategory == 1 {
		category = Food
	}
	if tempCategory == 2 {
		category = Tools
	}
	if tempCategory == 3 {
		category = Clothes
	}
	return nil, Category(category) //////////////////////////
}

func EditAmountValidate(amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	return nil
}

func Validate(p Product) error {
	switch {
	case p.Name == "":
		return ErrInvalidName
	case p.Price <= 0:
		return ErrInvalidPrice
	case p.Category < Food || p.Category > Clothes:
		return ErrInvalidCategory
	default:
		return nil
	}
}

func (w Warehouse) SearchProduct(id int) (Product, error) {
	for i, v := range w.Products {
		if w.Products[i].ID == id {
			return v, nil
		}
	}
	return Product{}, ErrProductNotFound
}
