package repositories

import (
	"ecommerce/models"
	"ecommerce/utils"
	"fmt"
	"sync"
)

const productFilePath = "data/products.json"

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
}

type FileProductRepository struct {
	mu       sync.Mutex
	products map[int]models.Product
	nextID   int
}

func NewFileProductRepository() *FileProductRepository {
	repo := &FileProductRepository{
		products: make(map[int]models.Product),
		nextID:   1,
	}
	// Load existing data
	utils.ReadFromFile(productFilePath, &repo.products)
	// Set nextID to max existing ID + 1
	for id := range repo.products {
		if id >= repo.nextID {
			repo.nextID = id + 1
		}
	}
	return repo
}

func (r *FileProductRepository) saveToFile() error {
	return utils.WriteToFile(productFilePath, r.products)
}

func (r *FileProductRepository) Create(product *models.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.ID = r.nextID
	r.products[r.nextID] = *product
	r.nextID++

	return r.saveToFile()
}

func (r *FileProductRepository) GetAll() ([]models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var productList []models.Product
	for _, product := range r.products {
		productList = append(productList, product)
	}
	return productList, nil
}

// Implementation of Update and Delete in FileProductRepository
func (r *FileProductRepository) GetByID(id int) (*models.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	product, exists := r.products[id]
	if !exists {
		return nil, fmt.Errorf("product not found")
	}

	return &product, nil
}

func (r *FileProductRepository) Update(product *models.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.products[product.ID]
	if !exists {
		return fmt.Errorf("product not found")
	}

	r.products[product.ID] = *product
	return r.saveToFile()
}

func (r *FileProductRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.products[id]
	if !exists {
		return fmt.Errorf("product not found")
	}

	delete(r.products, id)
	return r.saveToFile()
}
