package helper

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/clauderoy790/gratitude-journal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var quotes []repository.Quote
var quoteFile = "quotes.txt"
var QuotesHelper QuotesHelp = QuotesHelp{}

type QuotesHelp struct {
}

func (q *QuotesHelp) RefreshQuotes() {
	quotes = readQuoteFile(quoteFile)
	rebuildQuoteCollection(quotes)
}
func rebuildQuoteCollection(quotes []repository.Quote) {
	MongoHelper.QuotesCollection.Drop(MongoHelper.Context)

	err := MongoHelper.Db.CreateCollection(MongoHelper.Context, "quotes")
	if err != nil {
		fmt.Println("error creating collection:", err)
	}

	for _, q := range quotes {
		_, err = MongoHelper.QuotesCollection.InsertOne(MongoHelper.Context, q)
		if err != nil {
			log.Fatal()
		}
	}
}

func readQuoteFile(filename string) []repository.Quote {
	var quotes []repository.Quote

	//todo here
	// f, err := os.Open(filename)
	// defer f.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// scanner := bufio.NewScanner(f)
	// for i := 1; scanner.Scan(); i++ {
	// 	line := scanner.Text()
	// 	split := strings.Split(line, "|")
	// 	if len(split) != 2 {
	// 		log.Fatalf("Wrong quote format, fix this before continuing %v\n", line)
	// 	} else {
	// 		quote := repository.Quote{ID: primitive.NewObjectID(), QuoteID: i, Message: formatQuote(split[0]), Author: strings.TrimSpace(split[1])}
	// 		quotes = append(quotes, quote)
	// 	}
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	return quotes
}

//Format a quote so that every sentence starts with a capital letter, rest is lower and ends with a dot and capitalize I
func formatQuote(quote string) string {
	formattedQuote := ""

	quote = strings.TrimSpace(quote)
	quote = strings.ToUpper(quote[:1]) + quote[1:] //start with capital

	sentences := strings.Split(quote, ".")
	for _, sentence := range sentences {
		if len(sentence) == 0 {
			continue
		}

		capitalizedStart := false
		for i, ru := range sentence {
			character := string(ru)
			//Capitalize I
			if i > 0 && i < len(sentence)-1 && character == "i" && string(sentence[i-1]) == " " && string(sentence[i+1]) == " " {
				character = "I"
			}

			//Capitalize start of sentence
			if !capitalizedStart && unicode.IsLetter(ru) {
				character = strings.ToUpper(character)
				capitalizedStart = true
			}
			formattedQuote = formattedQuote + character
		}
		formattedQuote += "."
	}

	return formattedQuote
}
