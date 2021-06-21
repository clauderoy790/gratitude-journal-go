package helpers

import (
	"context"
	"fmt"
	"gitlab.com/claude.roy790/gratitude-journal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Context context.Context
var MongoHelper MongoHelp = MongoHelp{}

type MongoHelp struct {
	Context context.Context
	Db *mongo.Database
	QuotesCollection *mongo.Collection
	UsersCollection *mongo.Collection
	JournalEntriesCollection *mongo.Collection
}

func (MongoHelp) Connect() {
	var err error;
	cfg := config.Get()
	MongoHelper.Context = context.TODO()
	client, err = mongo.Connect(Context,options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%v:%v@%v/%v?retryWrites=true&w=majority",cfg.Database.User,cfg.Database.Password,cfg.Database.Cluster,cfg.Database.Name)))
	if err != nil {
		panic(err)
	}
	MongoHelper.Db = client.Database(cfg.Database.Name)
	MongoHelper.QuotesCollection  = MongoHelper.Db.Collection("quotes")
	MongoHelper.UsersCollection  = MongoHelper.Db.Collection("users")
	MongoHelper.JournalEntriesCollection  = MongoHelper.Db.Collection("journalEntries")
}

func (MongoHelp) Disconnect() {
	client.Disconnect(Context)
}