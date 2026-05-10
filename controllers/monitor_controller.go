package controllers

import (
	"bufio"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetLogs(c *gin.Context) {
	file, err := os.Open("app.log")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open log file"})
		return
	}
	defer file.Close()

	var logs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	// Only return the last 100 lines
	start := 0
	if len(logs) > 100 {
		start = len(logs) - 100
	}

	c.JSON(http.StatusOK, logs[start:])
}

func GetDocs(c *gin.Context) {
	content, err := os.ReadFile("API_DOCUMENTATION.md")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read documentation file"})
		return
	}
	c.Data(http.StatusOK, "text/plain; charset=utf-8", content)
}
