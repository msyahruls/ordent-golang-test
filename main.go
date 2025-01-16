package main

import (
	"ecommerce/controllers"
	"ecommerce/middlewares"
	"ecommerce/repositories"
	"ecommerce/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Ensure data directory exists
	if err := utils.EnsureDataDirectory(); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// Initialize file-based repositories
	userRepo := repositories.NewFileUserRepository()
	productRepo := repositories.NewFileProductRepository()
	orderRepo := repositories.NewFileOrderRepository()

	// Initialize controllers
	userController := controllers.NewUserController(userRepo)
	productController := controllers.NewProductController(productRepo)
	orderController := controllers.NewOrderController(orderRepo)

	// Set up routes
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/login", userController.Login).Methods("POST")
	router.HandleFunc("/register", userController.Register).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("").Subrouter() // Subrouter for middleware
	protected.Use(middlewares.Authenticate)        // Apply middleware to all routes in the subrouter

	protected.HandleFunc("/users", userController.GetAllUsers).Methods("GET")

	protected.HandleFunc("/products", productController.CreateProduct).Methods("POST")
	protected.HandleFunc("/products", productController.GetAllProducts).Methods("GET")
	protected.HandleFunc("/products/{id:[0-9]+}", productController.GetProductByID).Methods("GET")
	protected.HandleFunc("/products/{id:[0-9]+}", productController.UpdateProduct).Methods("PUT")
	protected.HandleFunc("/products/{id:[0-9]+}", productController.DeleteProduct).Methods("DELETE")

	protected.HandleFunc("/orders", orderController.CreateOrder).Methods("POST")
	protected.HandleFunc("/orders", orderController.GetAllOrders).Methods("GET")
	protected.HandleFunc("/orders/{id:[0-9]+}", orderController.GetOrderByID).Methods("GET")

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.Printf("Server running at http://localhost:%s", port)
	http.ListenAndServe(":"+port, router)
}
