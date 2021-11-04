package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	repo Repository
}

func (s *RepositoryTestSuite) SetupTest() {
	cfg := config.Get()
	db, err := ConnectToDatabase(&cfg)
	s.NoError(err)
	s.repo = NewRepository(db)

}

func (s *RepositoryTestSuite) TearDownTest() {
}

func (suite *RepositoryTestSuite) Test_repository_GetUser() {
	tests := []struct {
		name             string
		email            string
		password         string
		createUserBefore bool
		wantErr          bool
	}{
		{
			name:    "get non existing user returns error",
			email:   "testemail",
			wantErr: true,
		},
		{
			name:             "test get user runs properly with a freshly created used",
			email:            "test-email",
			password:         "testPW",
			createUserBefore: true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if tt.createUserBefore {
				userID, err := suite.repo.CreateUser(tt.email, tt.password)
				suite.NoError(err)
				suite.NotEqual(0, userID)
			}
			got, err := suite.repo.GetUser(tt.email)
			if tt.wantErr {
				suite.Error(err)
			} else {
				suite.NotEqual(got.ID, 0)
				suite.Equal(got.Email, tt.email)
				suite.Equal(got.PasswordHash, tt.password)
				suite.NoError(err)
				err = suite.repo.DeleteUser(tt.email)
				suite.NoError(err)
			}
		})
	}
}

func (suite *RepositoryTestSuite) Test_repository_CreateUser() {
	tests := []struct {
		name             string
		email            string
		password         string
		createUserBefore bool
		wantErr          bool
	}{
		{
			name:    "successful creation when non existing user",
			email:   "testemail",
			wantErr: false,
		},
		{
			name:             "fails to create when user already exists",
			email:            "test-email",
			password:         "testPW",
			createUserBefore: true,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var err error
			if tt.createUserBefore {
				userID, err := suite.repo.CreateUser(tt.email, tt.password)
				suite.NoError(err)
				suite.NotEqual(0, userID)
			}
			_, err = suite.repo.CreateUser(tt.email, tt.password)
			if tt.wantErr {
				suite.Error(err)
			} else {
				got, err := suite.repo.GetUser(tt.email)
				suite.NotEqual(got.ID, 0)
				suite.Equal(got.Email, tt.email)
				suite.Equal(got.PasswordHash, tt.password)
				suite.NoError(err)
				suite.NoError(err)
			}
			err = suite.repo.DeleteUser(tt.email)
			suite.NoError(err)
		})
	}
}

func (suite *RepositoryTestSuite) Test_repository_SaveJournalEntry() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	tests := []struct {
		name           string
		email          string
		date           time.Time
		entry          *JournalEntry
		overwrite      *JournalEntry
		quote          Quote
		overwriteQuote Quote
		wantErr        bool
	}{
		{
			name:    "fails to save an empty entry",
			wantErr: true,
		},
		{
			name: "saves new entry properly",
			date: today,
			entry: &JournalEntry{
				Grateful1:      "grateful1",
				Grateful2:      "grateful2",
				Grateful3:      "grateul-3",
				TodayGreat1:    "great1",
				TodayGreat2:    "great2",
				TodayGreat3:    "great3",
				Affirmation1:   "aff1",
				Affirmation2:   "aff2",
				HappenedGreat1: "happenedGreat1",
				HappenedGreat2: "happenedGreat2",
				HappenedGreat3: "happenedGreat3",
				Better1:        "b1",
				Better2:        "b2",
			},
			quote: Quote{
				Message: "this is a quote",
				Author:  "this is an author",
			},
		},
		{
			name: "overwrites pre existing entry",
			date: today,
			entry: &JournalEntry{
				Grateful1:      "grateful1",
				Grateful2:      "grateful2",
				Grateful3:      "grateul-3",
				TodayGreat1:    "great1",
				TodayGreat2:    "great2",
				TodayGreat3:    "great3",
				Affirmation1:   "aff1",
				Affirmation2:   "aff2",
				HappenedGreat1: "happenedGreat1",
				HappenedGreat2: "happenedGreat2",
				HappenedGreat3: "happenedGreat3",
				Better1:        "b1",
				Better2:        "b2",
			},
			quote: Quote{
				Message: "salut",
				Author:  "me",
			},
			overwrite: &JournalEntry{
				Grateful1:      "test overwrite",
				Grateful2:      "over 2",
				Grateful3:      "over 3",
				TodayGreat1:    "over great1",
				TodayGreat2:    "over great2",
				TodayGreat3:    "over great3",
				Affirmation1:   "over aff1",
				Affirmation2:   "over aff2",
				HappenedGreat1: "over happened Great1",
				HappenedGreat2: "over happened Great2",
				HappenedGreat3: "over happened Great3",
				Better1:        "b1",
				Better2:        "b2",
			},
			overwriteQuote: Quote{
				Message: "overWrite msg",
				Author:  "me ovw",
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			if tt.wantErr {
				err := suite.repo.SaveJournalEntry(tt.entry)
				suite.Error(err)
			} else {
				_, _ = suite.repo.CreateUser(tt.email, "rerer")
				usr, err := suite.repo.GetUser(tt.email)
				suite.NoError(err)

				// initialize entry data
				tt.entry.Date = tt.date.Local()
				tt.entry.UserID = usr.ID
				qID, err := suite.repo.SaveQuote(tt.quote.Message, tt.quote.Author)
				suite.NoError(err)
				tt.entry.QuoteID = qID

				// Save entry
				err = suite.repo.SaveJournalEntry(tt.entry)
				suite.NoError(err)

				// Set entry's quote
				quote, err := suite.repo.GetQuote(qID)
				suite.NoError(err)
				tt.entry.Quote = quote
				if tt.overwrite != nil {
					tt.overwrite.Date = tt.date.Local()
					tt.overwrite.UserID = usr.ID
					qID, err := suite.repo.SaveQuote(tt.overwriteQuote.Message, tt.overwriteQuote.Author)
					suite.NoError(err)
					tt.overwrite.QuoteID = qID
					err = suite.repo.SaveJournalEntry(tt.overwrite)
					suite.NoError(err)
					quote, err := suite.repo.GetQuote(qID)
					suite.NoError(err)
					tt.overwrite.Quote = quote
					tt.entry = tt.overwrite
				}
				suite.NoError(err, "repository.SaveJournalEntry() error = %v, wantErr %v", err, tt.wantErr)
				savedEntry, err := suite.repo.GetJournalEntry(tt.entry.UserID, tt.entry.Date)
				suite.NoError(err)
				suite.NotEqual(0, savedEntry.ID)

				// Set some gorm form that automatically has been set
				tt.entry.CreatedAt = savedEntry.CreatedAt
				tt.entry.UpdatedAt = savedEntry.UpdatedAt
				tt.entry.DeletedAt = savedEntry.DeletedAt
				tt.entry.ID = savedEntry.ID
				// suite.Equal(tt.entry.Date, savedEntry.Date)
				suite.Equal(tt.entry, savedEntry)
				fmt.Printf("actual: %#v\n", savedEntry)
				fmt.Printf("expected: %#v\n", tt.entry)

				if tt.entry != nil {
					err = suite.repo.DeleteJournalEntry(tt.entry.UserID, tt.date)
					suite.NoError(err)
					err = suite.repo.DeleteQuote(tt.quote.ID)
					suite.NoError(err)
				}
				err = suite.repo.DeleteUser(tt.email)
				suite.NoError(err)
				if tt.overwrite != nil {
					err = suite.repo.DeleteQuote(tt.overwrite.QuoteID)
					suite.NoError(err)
				}
				err = suite.repo.DeleteQuote(qID)
				suite.NoError(err)
			}
		})
	}
}

func (suite *RepositoryTestSuite) Test_repository_GetJournalEntry() {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	yestserday := today.Add(-time.Hour * 24)
	tomorrow := today.Add(time.Hour * 24)

	tests := []struct {
		name      string
		email     string
		callCount int
		dates     []time.Time
	}{
		{
			name:      "call three times in a row",
			email:     "testmail@gmail.com",
			callCount: 3,
			dates:     []time.Time{yestserday, tomorrow},
		},
		{
			name:      "call once",
			email:     "testmail2@gmail.com",
			callCount: 1,
			dates:     []time.Time{today},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			_, _ = suite.repo.CreateUser(tt.email, "password")
			user, _ := suite.repo.GetUser(tt.email)
			for _, date := range tt.dates {
				for i := 0; i < tt.callCount; i++ {
					entry, err := suite.repo.GetJournalEntry(user.ID, date)
					suite.NoError(err)
					suite.NotEqual(0, entry.ID)
					suite.NotEqual(0, entry.QuoteID)
				}
				err := suite.repo.DeleteJournalEntry(user.ID, date)
				suite.NoError(err)
			}
			err := suite.repo.DeleteUser(tt.email)
			suite.NoError(err)
		})
	}
}

func (suite *RepositoryTestSuite) Test_repository_GetRandomQuote() {

	for i := 0; i < 60; i++ {
		suite.Run(fmt.Sprintf("execution n%d", (i+1)), func() {
			email := "test@email.com"
			_ = suite.repo.DeleteUser(email)
			userID, err := suite.repo.CreateUser(email, "fsdfsdf")
			suite.NoError(err)
			now := time.Now()
			today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
			repo, ok := suite.repo.(*repository)
			suite.True(ok)
			quote, err := repo.getRandomQuote(userID, today)
			suite.NoError(err, "error getting random quote")
			suite.NotEqual(0, quote.ID)
			suite.NotEmpty(quote.Message)
			suite.NotEmpty(quote.Author)
			_ = suite.repo.DeleteUser(email)
		})
	}
}

func (suite *RepositoryTestSuite) TestConnectToDatabase() {
	tests := []struct {
		cfg     config.Config
		wantErr bool
	}{
		{
			cfg:     config.Get(),
			wantErr: false,
		},
		{
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := ConnectToDatabase(&tt.cfg)
		if tt.wantErr {
			suite.NotNil(got)
			suite.Error(err)
		} else {
			suite.NoError(err)
		}
	}
}

func Test_RepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
