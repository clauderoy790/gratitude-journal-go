package helper

import (
	"log"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/clauderoy790/gratitude-journal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var QuoteGenerator QuoteGen = QuoteGen{}

type QuoteGen struct{}

func (*QuoteGen) GetRandomQuote(userID, date string) (quote repository.Quote, err error) {
	mixedId, err := combineToStr(userID, date, 7)

	if err != nil {
		log.Fatal(err)
	}
	s1 := rand.NewSource(int64(mixedId))
	r1 := rand.New(s1)
	randNb := r1.Intn(int(QuotesHelper.QuotesCount())) + 1
	err = MongoHelper.QuotesCollection.FindOne(MongoHelper.Context, bson.D{{"quoteID", randNb}}).Decode(&quote)
	if err != nil {
		log.Fatal(err)
	}

	return quote, err
}

func combineToStr(id, date string, maxLen int) (nb int, err error) {

	// Make a Regex to say we only want numbers
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	// Remove all characters for combined Id + Date
	idNumbers := reg.ReplaceAllString(id, "")
	idNumbers = clamp(idNumbers, maxLen)
	processedString := idNumbers + reg.ReplaceAllString(date, "")

	nb, err = strconv.Atoi(processedString)
	return nb, err
}

func clamp(str string, max int) string {
	if len(str) > max {
		str = str[:max]
	} else {
		str = str[:len(str)]
	}
	return str
}
