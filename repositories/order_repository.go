package repositories

import (
	"ecommerce/models"
	"ecommerce/utils"
	"fmt"
	"sync"
)

const orderFilePath = "data/orders.json"

type OrderRepository interface {
	Create(order *models.Order) error
	GetAll() ([]models.Order, error)
	GetByID(id int) (*models.Order, error)
}

type FileOrderRepository struct {
	mu       sync.Mutex
	orders   map[int]models.Order
	nextID   int
	products map[int]models.Product
}

func NewFileOrderRepository() *FileOrderRepository {
	repo := &FileOrderRepository{
		orders: make(map[int]models.Order),
		nextID: 1,
	}
	// Load existing data
	utils.ReadFromFile(orderFilePath, &repo.orders)
	// Set nextID to max existing ID + 1
	for id := range repo.orders {
		if id >= repo.nextID {
			repo.nextID = id + 1
		}
	}
	utils.ReadFromFile(productFilePath, &repo.products)
	// Set nextID to max existing ID + 1
	for id := range repo.products {
		if id >= repo.nextID {
			repo.nextID = id + 1
		}
	}
	return repo
}

func (r *FileOrderRepository) saveToFile() error {
	// Write orders file
	if err := utils.WriteToFile(orderFilePath, r.orders); err != nil {
		return fmt.Errorf("failed to write orders file: %v", err)
	}

	// Write products file only if orders were successfully written
	if err := utils.WriteToFile(productFilePath, r.products); err != nil {
		return fmt.Errorf("failed to write products file: %v", err)
	}

	return nil
}

func (r *FileOrderRepository) Create(order *models.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Retrieve the product details and price
	product, exists := r.products[order.ProductID]
	if !exists {
		return fmt.Errorf("product not found")
	}

	// Calculate total price
	totalPrice := float64(order.Quantity) * product.Price
	order.TotalPrice = totalPrice

	// Update the stock of the product
	if product.Stock < order.Quantity {
		return fmt.Errorf("not enough stock available")
	}

	// Reduce the stock
	product.Stock -= order.Quantity
	r.products[order.ProductID] = product

	// Save the order
	order.ID = r.nextID
	r.orders[order.ID] = *order
	r.nextID++

	// Save to file
	return r.saveToFile()
}

func (r *FileOrderRepository) GetByID(id int) (*models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, exists := r.orders[id]
	if !exists {
		return nil, fmt.Errorf("order not found")
	}

	return &order, nil
}

func (r *FileOrderRepository) GetAll() ([]models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var orders []models.Order
	for _, order := range r.orders {
		orders = append(orders, order)
	}

	return orders, nil
}
