package handlers

import (
	"encoding/json"
	"fmt"
	"my_server/config" // Note: This string must match your module name in go.mod + /config
	"net/http"
)

type User struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go API Server")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/about created successfully")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Welcome to Contact Page")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method Not alllowed", http.StatusMethodNotAllowed)
		return
	}
	// Get the name from URL parameters, e.g., /delete?name=Harsh
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}
	// Execute the real SQL Delete statement
	result, err := config.DB.Exec("DELETE FROM users WHERE name = ?", name)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//check if rows was actully deleted
	rowsaffectes, _ := result.RowsAffected()
	if rowsaffectes == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Deleted successfully")

}

func UpdateUserAgeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedData User
	err := json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil || updatedData.Name == "" || updatedData.Age <= 0 {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Execute a clean SQL Update command
	result, err := config.DB.Exec("UPDATE users SET age = ? WHERE name = ?", updatedData.Age, updatedData.Name)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found to update", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User age updated successfully!")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := config.DB.Query("SELECT name, age FROM users")
	if err != nil {
		http.Error(w, "Database query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	dbUsers := make([]User, 0)

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Name, &u.Age)
		if err != nil {
			http.Error(w, "Row scanning error", http.StatusInternalServerError)
			return
		}
		dbUsers = append(dbUsers, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbUsers)
}

func CreateUserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid data structure", http.StatusBadRequest)
		return
	}

	if newUser.Name == "" {
		http.Error(w, "Name field cannot be empty", http.StatusBadRequest)
		return
	}
	if newUser.Age <= 0 {
		http.Error(w, "Age must be greater than 0", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (name, age) VALUES (?, ?)", newUser.Name, newUser.Age)
	if err != nil {
		http.Error(w, "Database insertion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User created successfully!")
}
