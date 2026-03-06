package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"warehousehttp/internal/warehouse"
)

type Server struct {
	Warehouse *warehouse.Warehouse
}

func NewServer(w *warehouse.Warehouse) Server {
	return Server{
		Warehouse: w,
	}
}

func (s *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	var bodyRequestDTO ProductDTO

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&bodyRequestDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct, err := DTOtoProduct(bodyRequestDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.Warehouse.CreateProduct(newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(buffer.Bytes())
}

func (s Server) ShowAllProducts(w http.ResponseWriter, r *http.Request) {
	if s.Warehouse == nil {
		http.Error(w, "NIL WAREHOUSE", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	defer r.Body.Close()
	for _, v := range s.Warehouse.Products {
		json.NewEncoder(w).Encode(v)
	}
}

func (s *Server) AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	idString, err := PathIndexValue(r.URL.Path, 2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var amount AmountDTOstruct

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := s.Warehouse.AddProduct(id, amount.AmountDTO)
	if err != nil {
		switch {
		case errors.Is(err, warehouse.ErrProductNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		case errors.Is(err, warehouse.ErrInvalidAmount):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var buffer bytes.Buffer

	if err = json.NewEncoder(&buffer).Encode(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}

func (s *Server) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	idString, err := PathIndexValue(r.URL.Path, 2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var amount AmountDTOstruct

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, warehouse.ErrInvalidID.Error(), http.StatusBadRequest)
		return
	}

	product, err := s.Warehouse.RemoveProduct(id, amount.AmountDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var buffer bytes.Buffer

	err = json.NewEncoder(&buffer).Encode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}

func (s *Server) SellProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	idString, err := PathIndexValue(r.URL.Path, 2)
	if err != nil {
		http.Error(w, ErrInvalidPath.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var amount AmountDTOstruct

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := s.Warehouse.SellProduct(id, amount.AmountDTO)
	if err != nil {
		switch {
		case errors.Is(err, warehouse.ErrInvalidAmount),
			errors.Is(err, warehouse.ErrNotEnoughProducts):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, warehouse.ErrProductNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var buffer bytes.Buffer
	if err = json.NewEncoder(&buffer).Encode(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}

func (s *Server) ShowProfit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
		return
	}

	var buffer bytes.Buffer
	var balance AmountDTOstruct
	balance.AmountDTO = s.Warehouse.Profit

	if err := json.NewEncoder(&buffer).Encode(balance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}
