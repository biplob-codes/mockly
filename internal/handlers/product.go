package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/biplob-codes/mockly/internal/model"
	"github.com/biplob-codes/mockly/internal/store"
)

func findProductById(id int) (model.Product, bool) {
	for _, p := range store.Products {
		if p.Id == id {
			return p, true
		}
	}
	return model.Product{}, false
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	writeRes(w, http.StatusOK, store.Products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	product, found := findProductById(n)
	if found {
		writeRes(w, http.StatusOK, product)
		return
	}
	writeRes(w, http.StatusNotFound, "Product with the Id not found")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req model.NewProduct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if req.Name == "" {
		writeRes(w, http.StatusBadRequest, "Name is required")
		return
	}
	createdProduct := model.Product{
		Id:          len(store.Products) + 1,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Sku:         req.Sku,
		Rating:      req.Rating,
		InStock:     req.InStock,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
	writeRes(w, http.StatusCreated, createdProduct)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	var req model.NewProduct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	updatedProduct := model.Product{
		Id:          n,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Sku:         req.Sku,
		Rating:      req.Rating,
		InStock:     req.InStock,
	}
	writeRes(w, http.StatusOK, updatedProduct)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		writeRes(w, http.StatusBadRequest, "Invalid number")
		return
	}
	message := fmt.Sprintf("Product with the Id: %d deleted", n)
	writeRes(w, http.StatusOK, message)
}