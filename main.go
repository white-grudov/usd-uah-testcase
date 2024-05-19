package main

import (
	"fmt"
	"log"
	"net/http"
	"usd-uah-testcase/db"
	"usd-uah-testcase/handlers"
	"usd-uah-testcase/internal"
	"usd-uah-testcase/utils"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

func main() {
	config := internal.LoadConfig()
	log.Println(config.DatabaseURL)

	err := db.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Start the cron job for sending emails
	c := cron.New()
	c.AddFunc("@daily", sendEmails)
	c.Start()
	defer c.Stop()

	r := mux.NewRouter()
	r.HandleFunc("/api/rate", handlers.GetRate).Methods("GET")
	r.HandleFunc("/api/subscribe", handlers.Subscribe).Methods("POST")

	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func sendEmails() {
	subscribers, err := db.GetSubscribers()
	if err != nil {
		log.Printf("Error getting subscribers: %v", err)
		return
	}

	rate, err := utils.FetchExchangeRate()
	if err != nil {
		log.Printf("Error fetching exchange rate: %v", err)
		return
	}

	for _, subscriber := range subscribers {
		err := utils.SendMail(subscriber.Email, "Daily Exchange Rate", fmt.Sprintf("Current USD to UAH rate: %f", rate))
		if err != nil {
			log.Printf("Error sending email to %s: %v", subscriber.Email, err)
		}
	}
}
