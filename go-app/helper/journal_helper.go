package helper

import (
	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var JournalHelper = JournalHelp{}

type JournalHelp struct{}

func (JournalHelp) WriteEntry(userID, date string, entry models.JournalEntry) (err error) {
	// entry.UserID = userID
	// entry.Date = date

	// res := JournalHelper.GetEntry(userID, date)
	// if res.Error == "" && res.Entry.ID != primitive.NilObjectID {
	// 	_, err = MongoHelper.JournalEntriesCollection.ReplaceOne(MongoHelper.Context, res.Entry, entry)
	// } else if res.Error == config.Get().Messages.NoEntryFound {
	// 	quote, err := QuoteGenerator.GetRandomQuote(userID, date)
	// 	if err == nil {
	// 		entry.Quote = quote
	// 		_, err = MongoHelper.JournalEntriesCollection.InsertOne(MongoHelper.Context, entry)
	// 	}
	// }
	// return err
	//todo here
	return nil
}

func (JournalHelp) GetEntry(userID, date string) models.JournalEntryResponse {
	entry := models.JournalEntry{}
	err := MongoHelper.JournalEntriesCollection.FindOne(MongoHelper.Context, bson.D{{Key: "userID", Value: userID}, {Key: "date", Value: date}}).Decode(&entry)

	if err == mongo.ErrNoDocuments {
		return models.JournalEntryResponse{Entry: models.JournalEntry{}, Error: config.Get().Messages.NoEntryFound}
	} else if err != nil {
		return models.JournalEntryResponse{Entry: models.JournalEntry{}, Error: "A server error occurred, try again later."}
	}
	return models.JournalEntryResponse{Entry: entry}
}

func (JournalHelp) DeleteEntry(userID, date string) error {
	// res := JournalHelper.GetEntry(userID, date)
	// var err error
	// if res.Error == "" && res.Entry.ID != primitive.NilObjectID {
	// 	_, err = MongoHelper.JournalEntriesCollection.DeleteOne(MongoHelper.Context, bson.D{{"userID", userID}, {"date", date}})
	// }
	// return err
	//todo here
	return nil
}
