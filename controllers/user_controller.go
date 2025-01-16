package controllers

import (
	"ecommerce/models"
	"ecommerce/repositories"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type UserController struct {
	UserRepo repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) *UserController {
	return &UserController{UserRepo: repo}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Save the user
	if err := c.UserRepo.Create(&user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}
	utils.JSONResponse(w, http.StatusCreated, true, "Users registered successfully", data)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Find the user by email
	user, err := c.UserRepo.GetByEmail(credentials.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Validate the password
	if err := utils.CheckPassword(credentials.Password, user.Password); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Error generating token", nil)
		return
	}

	data := map[string]interface{}{
		"token": token,
	}
	utils.JSONResponse(w, http.StatusOK, true, "Users logged successfully", data)
}

func (c *UserController) ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve claims from the request context
	claims, ok := utils.GetClaimsFromContext(r.Context())
	if !ok {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Unable to retrieve claims", nil)
		return
	}

	// Use claims (e.g., user_id, email) to perform necessary logic
	userID := claims["user_id"]
	email := claims["email"]

	// Respond with user information
	utils.JSONResponse(w, http.StatusOK, true, "Access granted", map[string]interface{}{
		"user_id": userID,
		"email":   email,
	})
}

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserRepo.GetAll()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, false, "Error fetching users", nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, true, "Users retrieved successfully", users)
}
