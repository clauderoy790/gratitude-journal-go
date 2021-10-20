package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/repository"
)

const (
	quoteFile = "../../quotes.txt"
)

func main() {
	cfg := config.Get()
	db, err := repository.ConnectToDatabase(&cfg)
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(db)

	err = repo.DeleteAllQuotes()
	if err != nil {
		panic(err)
	}

	quotes, err := readQuoteFile(quoteFile)

	if err != nil {
		panic(err)
	}

	for _, quote := range quotes {
		_, err = repo.SaveQuote(quote.Message, quote.Author)
		if err != nil {
			panic(fmt.Sprintf("Failed to save quote: %v", err))
		}
	}

	count, err := repo.QuotesCount()
	if err != nil {
		panic(fmt.Sprintf("Error retrieving quotes count: ", err))
	}
	fmt.Printf("Successfully created %d quotes!\n", count)
}

func readQuoteFile(filename string) ([]repository.Quote, error) {
	var quotes []repository.Quote

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		split := strings.Split(line, "|")
		if len(split) != 2 {
			return nil, errors.New(fmt.Sprintf("Wrong quote format, fix this before continuing %v\n", line))
		} else {
			quote := repository.Quote{Message: formatQuote(split[0]), Author: strings.TrimSpace(split[1])}
			quotes = append(quotes, quote)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return quotes, nil
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
