package models

import "time"

type Post struct {
	CollectionId   string    `json:"collectionId"`
	CollectionName string    `json:"collectionName"`
	ID             string    `json:"id"`
	Title          string    `json:"Title"`
	Content        string    `json:"content"`
	UserID         string    `json:"user_id"`
	Created        time.Time `json:"created"`
	Updated        time.Time `json:"updated"`
}
