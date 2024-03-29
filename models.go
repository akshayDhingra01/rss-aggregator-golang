package main

import (
	"time"

	"github.com/akshayDhingra01/rss-aggregator-golang/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModeifiedAt time.Time `json:"modified_at"`
	Name        string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:          dbUser.ID,
		CreatedAt:   dbUser.CreatedAt,
		ModeifiedAt: dbUser.ModeifiedAt,
		Name:        dbUser.Name,
	}
}
