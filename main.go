package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize the MySQL database connection
func initDB() {
	var err error

	dsn := "root:Loyalist@2468@tcp(localhost:3306)/avidb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}
	log.Println("Connected to MySQL database.")
}

// Handler to return the current Toronto time
func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Load Toronto timezone
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		log.Println("Error loading timezone:", err)
		http.Error(w, "Unable to load timezone", http.StatusInternalServerError)
		return
	}

	// Get current time in Toronto
	torontoTime := time.Now().In(location)
	response := map[string]string{
		"current_time": torontoTime.Format("2006-01-02 15:04:05"),
	}

	// Log the current time to the database
	_, err = db.Exec("INSERT INTO toronto_time_log (timestamp) VALUES (?)", torontoTime)
	if err != nil {
		log.Printf("Database Insert Error: %v\n", err) // Improved logging for database errors
		http.Error(w, "Failed to log time", http.StatusInternalServerError)
		return
	}

	// Respond with the current time in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler to retrieve all logged times
func timeLogsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, timestamp FROM time_log")
	if err != nil {
		log.Println("Error retrieving logs:", err)
		http.Error(w, "Failed to retrieve logs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Parse the results
	logs := []map[string]string{}
	for rows.Next() {
		var id int
		var timestamp string
		if err := rows.Scan(&id, &timestamp); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		logs = append(logs, map[string]string{
			"id":        string(rune(id)),
			"timestamp": timestamp,
		})
	}

	// Respond with the logs in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func main() {
	// Initialize the database
	initDB()
	defer db.Close()

	// Set up HTTP handlers
	http.HandleFunc("/current-time", currentTimeHandler)
	http.HandleFunc("/time-logs", timeLogsHandler)

	// Start the server
	log.Println("Server started at http://localhost:80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
