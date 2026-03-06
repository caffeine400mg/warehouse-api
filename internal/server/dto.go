package server

import (
	"math/rand"
	"warehousehttp/internal/warehouse"
)

type ProductDTO struct {
	Name     string `json:"product_name"`
	Price    int    `json:"product_price"`
	Category string `json:"product_category"`
}

type AmountDTOstruct struct {
	AmountDTO int `json:"amount"`
}

func DTOtoProduct(p ProductDTO) (warehouse.Product, error) {
	if err := ValidateDTO(p); err != nil {
		return warehouse.Product{}, err
	}

	category, _ := EncodeCategoryDTO(p.Category)
	newProduct := warehouse.Product{
		ID:       rand.Int(),
		Name:     p.Name,
		Price:    p.Price,
		Quantity: 0,
		Category: category,
	}
	return newProduct, nil
}

func ValidateDTO(p ProductDTO) error {
	category, err := EncodeCategoryDTO(p.Category)
	switch {
	case p.Name == "":
		return warehouse.ErrInvalidName
	case p.Price <= 0:
		return warehouse.ErrInvalidPrice
	case category < 1 || category > 3:
		return err
	default:
		return nil
	}
}

func EncodeCategoryDTO(category string) (warehouse.Category, error) {
	switch category {
	case "food":
		return 1, nil
	case "tools":
		return 2, nil
	case "clothes":
		return 3, nil
	default:
		return 0, warehouse.ErrInvalidCategory
	}
}
