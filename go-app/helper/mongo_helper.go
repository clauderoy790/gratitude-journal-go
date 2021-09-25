package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/clauderoy790/gratitude-journal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

var useLocalhost = true
var client *mongo.Client
var Context context.Context
var MongoHelper MongoHelp = MongoHelp{}

type MongoHelp struct {
	Context                  context.Context
	Db                       *gorm.DB 
	QuotesCollection         *mongo.Collection
	UsersCollection          *mongo.Collection
	JournalEntriesCollection *mongo.Collection
}

func (MongoHelp) Connect() {
	var err error
	cfg := config.Get()
	MongoHelper.Context = context.TODO()
	client, err = mongo.Connect(Context, options.Client().ApplyURI(getConnString(&cfg)))
	if err != nil {
		log.Fatalln("error connecting to the mongo client: ", err)
	}
	MongoHelper.Db = client.Database(cfg.Database.Host)
	MongoHelper.QuotesCollection = MongoHelper.Db.Collection("quotes")
	MongoHelper.UsersCollection = MongoHelper.Db.Collection("users")
	MongoHelper.JournalEntriesCollection = MongoHelper.Db.Collection("journalEntries")
}

func getConnString(cfg *config.Config) string {
	return fmt.Sprintf("mongodb+srv://%v:%v@%v/%v?retryWrites=true&w=majority", cfg.Database.User, cfg.Database.Password, cfg.Database.Cluster, cfg.Database.Host)
}

func (MongoHelp) Disconnect() {
	client.Disconnect(Context)
}
