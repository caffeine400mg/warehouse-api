package main

import (
	"net/http"
	"warehousehttp/internal/server"
	"warehousehttp/internal/warehouse"
)

func main() {
	InitWarehouse := warehouse.Warehouse{
		Products: []warehouse.Product{},
		Profit:   0,
	}
	InitServer := server.NewServer(&InitWarehouse)

	http.HandleFunc("/products/all", InitServer.ShowAllProducts)
	http.HandleFunc("/products/create", InitServer.CreateProduct)
	http.HandleFunc("/products/add/", InitServer.AddProduct)
	http.HandleFunc("/products/remove/", InitServer.RemoveProduct)
	http.HandleFunc("/products/sell/", InitServer.SellProduct)
	http.HandleFunc("/warehouse/profit", InitServer.ShowProfit)

	http.ListenAndServe(":666", nil)
}
