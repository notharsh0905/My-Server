package main

/*
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string
	Age  int
}

//***var users []User

var db *sql.DB // This global variable holds our live database connection pool

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Do something here before reaching the handler!
		fmt.Printf("Recieved Request: %s %s\n", r.Method, r.URL.Path)

		// 2. Pass the torch to the actual handler
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) { //w = send data to browser, //Receive request through r
	fmt.Fprintln(w, "Welcome to Go API Server")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/about creted sucessfully")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Welcome to Contact Page")
}

// 1. GET Handler - Sends JSON data to Postman
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Check if the method is bad
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// 2. If it is a good GET request, Go skips the IF block and runs this:
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Query all users from the database table
	rows, err := db.Query("SELECT name, age FROM users")
	if err != nil {
		http.Error(w, "Database query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// 2. Create a temporary slice to hold data we scan from the database
	dbUsers := make([]User, 0)

	// 3. Loop through each row in the database table
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Name, &u.Age)
		if err != nil {
			http.Error(w, "Row scanning error", http.StatusInternalServerError)
			return
		}
		dbUsers = append(dbUsers, u) // Add this database row to our temp slice
	}

	// 4. Send the slice back to Postman as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbUsers)
}

// Create Functions
// 2. POST Handler - Receives JSON data from Postman and saves it
func createUserData(w http.ResponseWriter, r *http.Request) {
	// If it's NOT a POST request, block them!
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Our server needs to read this incoming JSON text and translate it back into our Go User struct blueprint. This process is called Decoding (the exact opposite of what we did in the get handler).
	// To decode it, we use our "encoding/json" package again with a slightly different line:
	// 2. Decode the incoming JSON from Postman
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser) // We use r.Body instead of w because the data is coming from the Request body and // We pass &newUser with an ampersand (&).
	if newUser.Name == "" {
		http.Error(w, "Name field cannot be empty", http.StatusBadRequest)
		return
	}

	if newUser.Age <= 0 {
		http.Error(w, "Age must be greater than 0", http.StatusBadRequest)
		return
	}
	//******* 3. Append the new user to our slice
	//*******users = append(users, newUser)

	// 3. Insert the new user into our SQLite database table
	_, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", newUser.Name, newUser.Age)
	if err != nil {
		http.Error(w, "Database insertion error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// 4.Send success message back to Postman
	fmt.Fprintln(w, "User created successfully!")
}

func main() {
	//
	//***Starting data
	//***users = append(users, User{Name: "Harsh", Age: 21})
	//

	var err error
	// 1. Open (or create) the SQLite database file
	db, err = sql.Open("sqlite3", "./storage.db")
	if err != nil {
		panic(err) // If it can't open, crash the app immediately
	}
	defer db.Close() // Keep the connection open until main() finishes

	// 2. SQL command to create a table if it isn't there already
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        age INTEGER
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", homeHandler) //Connects the URL / to the function homeHandler
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	//Our API Endpoints
	http.Handle("/create", loggerMiddleware(http.HandlerFunc(createUserData)))
	http.Handle("/users", loggerMiddleware(http.HandlerFunc(getUserHandler)))
	http.ListenAndServe(":8090", nil) //Starts the server and tells it to listen for requests on port 8090. //Without this line: The program ends immediately because the server never starts listening.
}
*/
