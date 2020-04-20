package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/sgfoot.com", func(c *gin.Context) {
		WebHookLog("qweqwe")
	})
	r.Run(":6666")
}

func WebHookLog(secret string) gin.HandlerFunc {
	return Handler(secret, func(event string, payload *GitHubPayload, req *http.Request) error {
		// Log webhook
		fmt.Println("Received", event, "for ", payload.Repository.Name)

		// You'll probably want to do some real processing
		fmt.Println("Can clone repo at:", payload.Repository.CloneURL)

		// All is good (return an error to fail)
		return nil
	})
}