package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type logEntry struct {
	Timestamp  time.Time `json:"timestamp"`
	Method     string    `json:"method"`
	Path       string    `json:"path"`
	StatusCode int       `json:"status_code"`
	Latency    string    `json:"latency"`
	ClientIP   string    `json:"client_ip"`
	RequestID  string    `json:"request_id"`
}

func LoggingMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		value, _ := c.Get(RequestIDKey)
		requestID, _ := value.(string)

		entry := logEntry{
			Timestamp:  time.Now(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			StatusCode: c.Writer.Status(),
			Latency:    time.Since(start).String(),
			ClientIP:   c.ClientIP(),
			RequestID:  requestID,
		}

		data, err := json.Marshal(entry)
		if err == nil {
			log.Println(string(data)) // log ke stdout
			// bisa juga post ke endpoint log external
			// http.Post("http://log-service:4000/log", "application/json", bytes.NewBuffer(data))
		} else {
			log.Println("[LOG ERROR] failed to marshal log entry")
		}
	}
}
