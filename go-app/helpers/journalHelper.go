package helpers

import (
	"gitlab.com/claude.roy790/gratitude-journal/config"
	"gitlab.com/claude.roy790/gratitude-journal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var JournalHelper = JournalHelp{}

type JournalHelp struct{}

func (JournalHelp) WriteEntry(userId ,date string, entry models.JournalEntry)  (err error) {
	entry.UserID = userId
	entry.Date = date

	res := JournalHelper.GetEntry(userId,date)
	if res.Error == "" && res.Entry.ID != primitive.NilObjectID {
		_, err = MongoHelper.JournalEntriesCollection.ReplaceOne(MongoHelper.Context,res.Entry,entry)
	} else if res.Error == config.Get().Messages.NoEntryFound {
		_,err = MongoHelper.JournalEntriesCollection.InsertOne(MongoHelper.Context,entry)
	}
	return err
}

func (JournalHelp) GetEntry(userId,date string) models.JournalEntryResult{
	entry := models.JournalEntry{}
	err := MongoHelper.JournalEntriesCollection.FindOne(MongoHelper.Context, bson.D{{"userID", userId},{"date",date}}).Decode(&entry)

	if  err == mongo.ErrNoDocuments {
		return models.JournalEntryResult{models.JournalEntry{},config.Get().Messages.NoEntryFound}
	} else if  err != nil {
		return models.JournalEntryResult{models.JournalEntry{},"A server error occurred, try again later."}
	}
	return models.JournalEntryResult{entry,""}
}

func (JournalHelp) DeleteEntry(userId,date string) error {
	res := JournalHelper.GetEntry(userId,date)
	var err error
	if res.Error == "" && res.Entry.ID != primitive.NilObjectID {
		_,err = MongoHelper.JournalEntriesCollection.DeleteOne(MongoHelper.Context,bson.D{{"userID", userId},{"date",date}})
	}
	return err
}
