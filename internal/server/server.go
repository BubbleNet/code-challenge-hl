package server

import (
	"encoding/json"
	"fmt"
	"github.com/BubbleNet/code-challenge-hl/internal/database"
	"github.com/BubbleNet/code-challenge-hl/internal/scheduler"
	"github.com/BubbleNet/code-challenge-hl/pkg/session"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type (
	Server struct {
		db *database.Database
	}
)

// CreateAndServe creates and starts the rest api client.
func CreateAndServe() {
	s := Server{database.New()}
	go scheduler.Start()
	r := gin.Default()
	r.POST("/session", s.createSession)
	r.GET("/reminders/:session_id", s.getReminders)
	r.Run()
}

func (s *Server) createSession(c *gin.Context) {
	newJsonSession := new(session.JsonSession)
	err := c.BindJSON(newJsonSession)
	if err != nil {
		//  TODO: Return additional error information
		c.Status(http.StatusBadRequest)
		return
	}
	newSession, err := newJsonSession.ToSession()
	if err != nil {
		//  TODO: Return additional error information
		c.Status(http.StatusBadRequest)
		return
	}
	id, err := s.db.SetSession(newSession)
	if err != nil {
		//  TODO: Return additional error information
		c.Status(http.StatusBadRequest)
		return
	}
	c.Data(http.StatusOK,
		"application/json",
		[]byte(fmt.Sprintf(`{"session_id": "%s"}`, id.String())))
}

func (s *Server) getReminders(c *gin.Context) {
	sessionId := c.Param("session_id")
	u, err := uuid.Parse(sessionId)
	if err != nil {
		//  TODO: Return additional error information
		c.Status(http.StatusBadRequest)
		return
	}
	reminders, err := s.db.GetReminders(u)
	if err != nil {
		//  TODO: Return additional error information
		c.Status(http.StatusBadRequest)
		return
	}
	remindersJson, err := json.Marshal(reminders)
	if err != nil {
		//  TODO: Return additional error information
		c.Status(http.StatusBadRequest)
		return
	}
	c.Data(http.StatusOK,
		"application/json",
		remindersJson)
}
