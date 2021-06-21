package helpers

import (
	"bufio"
	"fmt"
	"gitlab.com/claude.roy790/gratitude-journal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"strings"
	"unicode"
)

var quotes []models.Quote
var quoteFile = "quotes.txt"
var QuotesHelper QuotesHelp = QuotesHelp{}

type QuotesHelp struct{}

func (QuotesHelp) RefreshQuotes() {
	quotes = readQuoteFile(quoteFile)
	rebuildQuoteCollection(quotes)
	quote,author := findLongestQuoteAndAuthor(quotes)
	fmt.Println("longest quote is: ",quote)
	fmt.Println("longest author is: ",author)
}

func findLongestQuoteAndAuthor(quotes []models.Quote) (string,string) {
	author:= ""
	message := ""
	for _,quote := range quotes {
		if len(quote.Message) > len(message) {
			message = quote.Message
		}

		if len(quote.Author) > len(author) {
			author = quote.Author
		}
	}

	return message,author
}

func rebuildQuoteCollection(quotes []models.Quote) {
	MongoHelper.QuotesCollection.Drop(MongoHelper.Context)

	err := MongoHelper.Db.CreateCollection(MongoHelper.Context, "quotes")
	if err != nil {
		fmt.Println("error creating collection:", err)
	}

	for _, q := range quotes {
		MongoHelper.QuotesCollection.InsertOne(MongoHelper.Context, q)
	}
}

func readQuoteFile(filename string) []models.Quote {
	quotes := []models.Quote{}

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "|")
		if len(split) != 2 {
			panic(fmt.Sprintf("Wrong quote format, fix this before continuing %v\n", line))
		} else {
			quote := models.Quote{primitive.NewObjectID(), formatQuote(split[0]), strings.TrimSpace(split[1])}
			quotes = append(quotes, quote)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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
