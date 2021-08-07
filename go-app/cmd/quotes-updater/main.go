package main

import "github.com/clauderoy790/gratitude-journal/helpers"

func main() {
	helpers.MongoHelper.Connect()
	helpers.QuotesHelper.RefreshQuotes()
	helpers.MongoHelper.Disconnect()
}