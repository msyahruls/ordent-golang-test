package controllers

import (
	"ecommerce/models"
	"ecommerce/repositories"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	ProductRepo repositories.ProductRepository
}

func NewProductController(repo repositories.ProductRepository) *ProductController {
	return &ProductController{ProductRepo: repo}
}

// CreateProduct handles creating a new product.
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, "Invalid request payload", nil)
		return
	}

	if err := c.ProductRepo.Create(&product); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Error creating product", nil)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, true, "Product created successfully", product)
}

// GetAllProducts handles fetching all products.
func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.ProductRepo.GetAll()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Error fetching products", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Products retrieved successfully", products)
}

// GetProductByID handles fetching a single product by its ID.
func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	product, err := c.ProductRepo.GetByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, false, "Product not found", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Product retrieved successfully", product)
}

// UpdateProduct handles updating an existing product.
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, "Invalid request payload", nil)
		return
	}

	product.ID = id // Ensure the ID is set to the correct one from the URL
	if err := c.ProductRepo.Update(&product); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Error updating product", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Product updated successfully", product)
}

// DeleteProduct handles deleting a product by its ID.
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, "Invalid product ID", nil)
		return
	}

	if err := c.ProductRepo.Delete(id); err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Error deleting product", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Product deleted successfully", nil)
}
