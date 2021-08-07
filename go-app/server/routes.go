package server

import (
	"encoding/json"
	"fmt"
	"github.com/clauderoy790/gratitude-journal/config"
	"github.com/clauderoy790/gratitude-journal/helpers"
	"github.com/clauderoy790/gratitude-journal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"net/http"
)

func (server *Server) setupRoutes() {
	serverMux := mux.NewRouter()
	serverMux.HandleFunc("/", server.homeHandler).Methods("GET")
	// Publish scores AGS 2.0 (extension to the scores endpoint above)
	serverMux.Handle("/{sectionXid}/assignments/{assignmentXid}/attempt/{attemptId}/lineitems/{assetXid}/scores", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idmAuth.Wrap(s.ctx, s.publishScoresExtendedHandler(), s.ds).ServeHTTP(w, r)
	})).Methods("POST")

	// Publish scores Outcomes1.0, doesnt use IDM auth, uses sourcedID decode/validation logic
	serverMux.Handle("/lti/outcomev1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idmAuth.Wrap(s.ctx, s.publishOutcomesHandler(), s.ds).ServeHTTP(w, r)
	})).Methods("POST")

	// Delete line items
	serverMux.Handle("/{contextId}/lineitems/{lineItemId}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idmAuth.Wrap(s.ctx, s.deleteLineItemHandler(), s.ds).ServeHTTP(w, r)
	})).Methods("DELETE")

	// Fetch a line item
	serverMux.Handle("/{contextId}/lineitems/{lineItemId}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idmAuth.Wrap(s.ctx, s.fetchLineItemHandler(), s.ds).ServeHTTP(w, r)
	})).Methods("GET")

	// Update a line item
	serverMux.Handle("/{contextId}/lineitems/{lineItemId}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idmAuth.Wrap(s.ctx, deprecatedSubmissionKeyMiddleware.Wrap(s.updateLineItemHandler()), s.ds).ServeHTTP(w, r)
	})).Methods("PUT")

	// Get assignment scores (results)
	serverMux.Handle("/{contextId}/lineitems/{lineItemId}/results", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idmAuth.Wrap(s.ctx, s.resultLineItemHandler(), s.ds).ServeHTTP(w, r)
	})).Methods("GET")
	r := gin.Default()
	c.JSON()
	gin.H{}
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "welcome to daily gratitude",
	//	})
	//})
	//
	//r.POST("/login", func(c *gin.Context) {
	//	logRes := helpers.UserHelper.Login(c.PostForm("email"), c.PostForm("password"))
	//	c.JSON(200, logRes)
	//})
	//
	//r.POST("/register", func(c *gin.Context) {
	//	regReg := helpers.UserHelper.Register(c.PostForm("email"), c.PostForm("password"), c.PostForm("verifyPassword"))
	//	c.JSON(200, regReg)
	//})
	//
	//r.GET("/journal", func(c *gin.Context) {
	//	userID := c.Query("userID")
	//	date := c.Query("date")
	//
	//	if userID == "" || date == "" {
	//		c.JSON(200, gin.H{
	//			"error": "must provide userID and date",
	//		})
	//	} else {
	//		//Read entry
	//		journalRes := helpers.JournalHelper.GetEntry(userID, date)
	//		if journalRes.Error != "" {
	//			c.JSON(200, gin.H{
	//				"error": journalRes.Error,
	//			})
	//		} else {
	//			c.JSON(200, gin.H{
	//				"entry": journalRes.Entry,
	//			})
	//		}
	//	}
	//})
	//
	//r.PUT("/journal", func(c *gin.Context) {
	//	userID := c.PostForm("userID")
	//	date := c.PostForm("date")
	//	entry := c.PostForm("entry")
	//
	//	if userID == "" || date == "" {
	//		c.JSON(200, gin.H{
	//			"error": "must provide userID and date",
	//		})
	//	} else if entry != "" {
	//		jEntry := models.JournalEntry{}
	//		err := json.Unmarshal([]byte(entry), &jEntry)
	//		if err != nil {
	//			c.JSON(200, gin.H{
	//				"error": "entry is not in a valid format.",
	//			})
	//		} else {
	//			if err := helpers.JournalHelper.WriteEntry(userID, date, jEntry); err != nil {
	//				c.JSON(200, gin.H{
	//					"error": err,
	//				})
	//			} else {
	//				c.JSON(200, gin.H{
	//					"entry": jEntry,
	//				})
	//			}
	//		}
	//	}
	//})
	//
	//if err := r.Run(fmt.Sprintf(":%v", config.Get().App.Port)); err != nil {
	//	panic(err)
	//}
}
