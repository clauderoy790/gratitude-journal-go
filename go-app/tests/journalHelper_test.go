package tests

import (
	"fmt"
	"gitlab.com/claude.roy790/gratitude-journal/helpers"
	"gitlab.com/claude.roy790/gratitude-journal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var today = time.Now().Format("01-02-2006")
var testEntry models.JournalEntry = models.JournalEntry{
	Date: 			today,
	Grateful1:      "test1",
	Grateful2:      "test2",
	Grateful3:      "test3",
	TodayGreat1:    "great1",
	TodayGreat2:    "great2",
	TodayGreat3:    "",
	Affirmation1:   "affirmation1",
	Affirmation2:   "",
	HappenedGreat1: "wasGreat1",
	HappenedGreat2: "wasGreat2",
	HappenedGreat3: "wasGreat3",
	Better1:        "",
	Better2:        "better2",
}

func TestDeleteEntry(t *testing.T) {
	helpers.MongoHelper.Connect()
	defer helpers.MongoHelper.Disconnect()
	jh := helpers.JournalHelper
	uh := helpers.UserHelper
	uh.Register(email, password, password)
	user, _ := uh.GetUser(email)

	jh.WriteEntry(user.ID.Hex(), today, testEntry)
	jh.DeleteEntry(user.ID.Hex(), today)

	res := jh.GetEntry(user.ID.Hex(), today)
	if res.Error == "" {
		t.Fatalf("got no error while trying to get a freshly deleted entry")
	}
	if res.Entry != (models.JournalEntry{}) {
		t.Fatalf("got a non empty entry right after deleting it")
	}
}

func TestWriteEntry(t *testing.T) {
	helpers.MongoHelper.Connect()
	defer helpers.MongoHelper.Disconnect()
	jh := helpers.JournalHelper
	uh := helpers.UserHelper
	uh.Register(email, password, password)
	user, _ := uh.GetUser(email)
	fmt.Println("user id: ",user.ID)
	jh.DeleteEntry(user.ID.Hex(), today)
	if err := jh.WriteEntry(user.ID.Hex(), today, testEntry); err != nil {
		t.Fatalf("failed to write new entry: %v", err)
	}

	res := jh.GetEntry(user.ID.Hex(), today)
	if res.Error != "" {
		t.Fatalf("failed to get a freshly created entry %v",res.Error)
	}
	if res.Entry.UserID == "" || res.Entry.UserID == primitive.NilObjectID.Hex() {
		t.Fatalf("invalied userId in freshly creaty entry")
	}

	jh.DeleteEntry(user.ID.Hex(), today)
}

func TestReadEntry(t *testing.T) {
	helpers.MongoHelper.Connect()
	defer helpers.MongoHelper.Disconnect()
	jh := helpers.JournalHelper
	uh := helpers.UserHelper
	uh.Register(email, password, password)
	user, _ := uh.GetUser(email)
	jh.WriteEntry(user.ID.Hex(),today,testEntry)
	res := jh.GetEntry(user.ID.Hex(),today)
	if res.Error != "" {
		t.Fatalf("failed to read a freshly added entry. %v",res.Entry)
	}

	testEntry.UserID = user.ID.Hex()
	res.Entry.ID = testEntry.ID
	if res.Entry != testEntry{
		t.Fatalf("get entry returned different content from inserted entry e1: %v, e2: %v",res.Entry, testEntry)
	}

	jh.DeleteEntry(user.ID.Hex(), today)
}

func TestUpdateEntry(t *testing.T) {
	helpers.MongoHelper.Connect()
	defer helpers.MongoHelper.Disconnect()
	jh := helpers.JournalHelper
	uh := helpers.UserHelper
	uh.Register(email, password, password)
	user, _ := uh.GetUser(email)

	jh.WriteEntry(user.ID.Hex(),today,testEntry)
	updatedEntry := models.JournalEntry{
		UserID:			user.ID.Hex(),
		Date: 			today,
		Grateful1:      "updatedTest1",
		Grateful2:      "updatedtest2",
		Grateful3:      "updatedtest3",
		TodayGreat1:    "updatedgreat1",
		TodayGreat2:    "updatedgreat2",
		TodayGreat3:    "updated",
		Affirmation1:   "updatedaffirmation1",
		Affirmation2:   "updated",
		HappenedGreat1: "updatedwasGreat1",
		HappenedGreat2: "updatedwasGreat2",
		HappenedGreat3: "updatedwasGreat3",
		Better1:        "updated",
		Better2:        "updatedbetter2",
	}
	err := jh.WriteEntry(user.ID.Hex(),today,updatedEntry)
	if err != nil {
		t.Fatalf("failed to write an updated entry err:%v", err)
	}

	res := jh.GetEntry(user.ID.Hex(),today)
	if res.Error != "" {
		t.Fatalf("failed to read an updated entry")
	}

	res.Entry.ID = updatedEntry.ID
	if res.Entry != updatedEntry {
		t.Fatalf("get entry did not return the properly updated entry content")
	}

	jh.DeleteEntry(user.ID.Hex(), today)
}
