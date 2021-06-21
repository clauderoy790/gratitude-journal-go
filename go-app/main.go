package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/claude.roy790/gratitude-journal/config"
	"gitlab.com/claude.roy790/gratitude-journal/helpers"
	"gitlab.com/claude.roy790/gratitude-journal/models"
)

func main() {
	//refreshQuotes()
	startServer()
}
func startServer() {
	helpers.MongoHelper.Connect()
	defer helpers.MongoHelper.Disconnect()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to daily gratitude",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		logRes := helpers.UserHelper.Login(c.PostForm("email"),c.PostForm("password"))
		c.JSON(200, logRes)
	})

	r.POST("/register", func(c *gin.Context) {
		regReg := helpers.UserHelper.Register(c.PostForm("email"),c.PostForm("password"),c.PostForm("verifyPassword"))
		c.JSON(200, regReg)
	})

	r.POST("/journal", func(c *gin.Context) {
		userId := c.PostForm("userID")
		date := c.PostForm("date")
		entry := c.PostForm("entry")

		if userId == "" || date == "" {
			c.JSON(200,gin.H{
				"error":  "must provide userID and date",
			})
		} else if entry != "" {
			jEntry := models.JournalEntry{}
			err := json.Unmarshal([]byte(entry),&jEntry)
			if err != nil {
				c.JSON(200,gin.H{
					"error":  "entry is not in a valid format.",
				})
			} else {
				if err := helpers.JournalHelper.WriteEntry(userId,date,jEntry);err != nil {
					c.JSON(200,gin.H{
						"error":  err,
					})
				} else {
					c.JSON(200,gin.H{
						"entry":  jEntry,
					})
				}
			}
		} else {
			//Read entry
			journalRes := helpers.JournalHelper.GetEntry(userId,date)
			if journalRes.Error != "" {
				c.JSON(200,gin.H{
					"error":  journalRes.Error,
				})
			} else {
				c.JSON(200,gin.H{
					"entry":  journalRes.Entry,
				})
			}
		}
	})

	if err := r.Run(fmt.Sprintf(":%v",config.Get().App.Port));err != nil {
		panic(err)
	}
}

func refreshQuotes() {
	helpers.MongoHelper.Connect()
	helpers.QuotesHelper.RefreshQuotes()
}
