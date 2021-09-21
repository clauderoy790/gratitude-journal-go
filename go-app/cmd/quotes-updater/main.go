package main

import "github.com/clauderoy790/gratitude-journal/helper"

func main() {
	helper.MongoHelper.Connect()
	helper.QuotesHelper.RefreshQuotes()
	helper.MongoHelper.Disconnect()
}
