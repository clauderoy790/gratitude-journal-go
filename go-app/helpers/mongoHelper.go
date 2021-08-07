package helpers

import (
	"context"
	"fmt"
	"github.com/clauderoy790/gratitude-journal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var isLocalhost = true
var client *mongo.Client
var Context context.Context
var MongoHelper MongoHelp = MongoHelp{}

type MongoHelp struct {
	Context                  context.Context
	Db                       *mongo.Database
	QuotesCollection         *mongo.Collection
	UsersCollection          *mongo.Collection
	JournalEntriesCollection *mongo.Collection
}

func (MongoHelp) Connect() {
	var err error
	cfg := config.Get()
	cfg.UseLocalhost = isLocalhost
	MongoHelper.Context = context.TODO()
	client, err = mongo.Connect(Context, options.Client().ApplyURI(getConnString(cfg)))
	if err != nil {
		panic(err)
	}
	MongoHelper.Db = client.Database(cfg.Database.Name)
	MongoHelper.QuotesCollection = MongoHelper.Db.Collection("quotes")
	MongoHelper.UsersCollection = MongoHelper.Db.Collection("users")
	MongoHelper.JournalEntriesCollection = MongoHelper.Db.Collection("journalEntries")
}

func getConnString(cfg config.Config) string {
	if cfg.UseLocalhost {
		return "mongodb://localhost:27017"
	}
	return fmt.Sprintf("mongodb+srv://%v:%v@%v/%v?retryWrites=true&w=majority", cfg.Database.User, cfg.Database.Password, cfg.Database.Cluster, cfg.Database.Name)
}

func (MongoHelp) Disconnect() {
	client.Disconnect(Context)
}
