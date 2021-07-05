package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quote struct {
	ID     		primitive.ObjectID 	`bson:"_id,omitempty"`
	QuoteID     int 				`bson:"quoteID,omitempty"`
	Message		string             	`bson:"message,omitempty"`
	Author 		string             	`bson:"author,omitempty"`
}

type QuoteResult struct {
	Message	string             	`json:"message"`
	Author 	string             	`json:"author"`
	Error 	string             	`json:"error"`
}
