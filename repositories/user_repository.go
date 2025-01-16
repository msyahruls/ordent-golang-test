package repositories

import (
	"ecommerce/models"
	"ecommerce/utils"
	"fmt"
	"sync"
)

const userFilePath = "data/users.json"

type UserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type FileUserRepository struct {
	mu     sync.Mutex
	users  map[int]models.User
	nextID int
}

func NewFileUserRepository() *FileUserRepository {
	repo := &FileUserRepository{
		users:  make(map[int]models.User),
		nextID: 1,
	}
	// Load existing data
	utils.ReadFromFile(userFilePath, &repo.users)
	// Set nextID to max existing ID + 1
	for id := range repo.users {
		if id >= repo.nextID {
			repo.nextID = id + 1
		}
	}
	return repo
}

func (r *FileUserRepository) saveToFile() error {
	return utils.WriteToFile(userFilePath, r.users)
}

func (r *FileUserRepository) Create(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.users[r.nextID] = *user
	r.nextID++

	return r.saveToFile()
}

func (r *FileUserRepository) GetAll() ([]models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userList []models.User
	for _, user := range r.users {
		userList = append(userList, user)
	}
	return userList, nil
}

func (r *FileUserRepository) GetByEmail(email string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
