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
