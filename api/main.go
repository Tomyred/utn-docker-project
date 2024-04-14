package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Date struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
}

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/save", saveDate)

	http.HandleFunc("/dates", getDates)

	port := ":8080"

	fmt.Printf("Server listening on PORT %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error initializing server: %s", err)
	}
}

func saveDate(w http.ResponseWriter, r *http.Request) {
	postgres, err := NewDB()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer postgres.db.Close()

	query := `INSERT INTO dates (date) VALUES ($1);`

	date := time.Now().Format("2006-01-02 15:04:05")

	_, qerror := postgres.db.Query(query, date)

	if qerror != nil {
		http.Error(w, qerror.Error(), http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("Fecha %v guardada", date)

	w.Write([]byte(msg))
}

func getDates(w http.ResponseWriter, r *http.Request) {

	postgres, err := NewDB()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer postgres.db.Close()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := `SELECT * FROM dates;`

	rows, err := postgres.db.Query(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var dates []Date
	for rows.Next() {
		var date Date
		err := rows.Scan(&date.ID, &date.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dates = append(dates, date)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dates)
}
