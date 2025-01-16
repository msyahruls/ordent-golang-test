package controllers

import (
	"ecommerce/models"
	"ecommerce/repositories"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderController struct {
	repo repositories.OrderRepository
}

func NewOrderController(repo repositories.OrderRepository) *OrderController {
	return &OrderController{
		repo: repo,
	}
}

// CreateOrder handles creating a new order.
func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, "Invalid request payload", nil)
		return
	}

	if err := c.repo.Create(&order); err != nil {
		fmt.Println(err)
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Error creating order", nil)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, true, "Order created successfully", order)
}

// GetAllOrders handles fetching all orders.
func (c *OrderController) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := c.repo.GetAll()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Error fetching orders", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Orders retrieved successfully", orders)
}

// GetOrderByID handles fetching a single order by its ID.
func (c *OrderController) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, false, "Invalid order ID", nil)
		return
	}

	order, err := c.repo.GetByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, false, "Order not found", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Order retrieved successfully", order)
}
