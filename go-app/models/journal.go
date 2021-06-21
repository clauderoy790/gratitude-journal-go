package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JournalEntry struct {
	ID     			primitive.ObjectID 	`bson:"_id,omitempty"`
	UserID			string             	`bson:"userID,omitempty"`
	Date			string				`bson:"date,omitempty"`
	Grateful1		string				`bson:"grateful1,omitempty"`
	Grateful2		string				`bson:"grateful2,omitempty"`
	Grateful3		string				`bson:"grateful3,omitempty"`
	TodayGreat1		string				`bson:"todayGreat1,omitempty"`
	TodayGreat2		string				`bson:"todayGreat2,omitempty"`
	TodayGreat3		string				`bson:"todayGreat3,omitempty"`
	Affirmation1	string				`bson:"affirmation1,omitempty"`
	Affirmation2	string				`bson:"affirmation2,omitempty"`
	HappenedGreat1	string				`bson:"happenedGreat1,omitempty"`
	HappenedGreat2	string				`bson:"happenedGreat2,omitempty"`
	HappenedGreat3	string				`bson:"happenedGreat3,omitempty"`
	Better1			string				`bson:"better1,omitempty"`
	Better2			string				`bson:"better2,omitempty"`
}

type JournalEntryResult struct {
	Entry JournalEntry 	`json:"entry"`
	Error string 		`json:"error"`
}