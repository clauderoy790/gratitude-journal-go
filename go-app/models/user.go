package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID     			primitive.ObjectID 	`bson:"_id,omitempty"`
	Email			string             	`bson:"email,omitempty"`
	PasswordHash	string             	`bson:"passwordHash,omitempty"`
	DateCreated		time.Time			`bson:"dateCreated,omitempty"`
}

type LoginResult struct {
	UserId string	`json:"userId"`
	Success bool 	`json:"success"`
	Error string	`json:"error"`
}

type RegisterResult struct {
	UserID string `json:"userID"`
	Error string  `json:"error"`
}
