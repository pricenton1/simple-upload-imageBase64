package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simple-upload-file/models"
	"simple-upload-file/services"
	"strings"
)

type userController struct {
	userService services.IUserService
}

func NewUserController(db *sql.DB) {
	controller := userController{
		userService: services.NewUserService(db),
	}

	var id string
	http.HandleFunc(fmt.Sprintf("/api/user/%s", id), controller.GetUser)
	http.HandleFunc("/api/user", controller.CreateUser)
	http.HandleFunc(fmt.Sprintf("/api/user/update/%s", id), controller.UpdateUser)
	http.HandleFunc(fmt.Sprintf("/api/user/delete/%s", id), controller.DeleteUser)
}

// GetUser Controller
func (u *userController) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("Endpoint ", r.Method, r.RequestURI)

	var id string
	path := strings.Split(r.URL.Path, "/")[1:]
	id = path[2]

	user, err := u.userService.GetUser(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}

// CreateUser Controller
func (u *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint ", r.Method, r.RequestURI)
	var user models.User

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err := u.userService.CreateUser(&user)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "user successfully created",
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

}

// UpdateUser Controller
func (u *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Endpoint ", r.Method, r.RequestURI)

	var id string
	var user models.User
	// fmt.Println("INI PATH", p)
	path := strings.Split(r.URL.Path, "/")[1:]
	id = path[3]

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err := u.userService.UpdateUser(&user, id)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "user successfully updated",
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Deleteuser Controller
func (u *userController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var id string
	path := strings.Split(r.URL.Path, "/")[1:]
	id = path[3]

	idDelete, err := u.userService.DeleteUser(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"id":      idDelete,
		"code":    http.StatusOK,
		"message": "user successfully deleted",
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

}
