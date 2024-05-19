package db

import (
	"database/sql"
	"fmt"
	"time"
	"usd-uah-testcase/models"

	"github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dataSourceName)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return nil
			}
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("could not connect to the database: %v", err)
}

func AddSubscriber(subscriber models.Subscriber) error {
	_, err := db.Exec("INSERT INTO subscribers (email) VALUES ($1)", subscriber.Email)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code.Name() == "unique_violation" {
			return fmt.Errorf("unique_violation")
		}
		return err
	}
	return nil
}

func GetSubscribers() ([]models.Subscriber, error) {
	rows, err := db.Query("SELECT id, email FROM subscribers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribers []models.Subscriber
	for rows.Next() {
		var subscriber models.Subscriber
		if err := rows.Scan(&subscriber.ID, &subscriber.Email); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}
	return subscribers, nil
}
