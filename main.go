package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "PONG")
	})
	r.POST("/sgfoot.com", WebHookLog("qweqwe"))
	r.Run(":6666")
}

func WebHookLog(secret string) gin.HandlerFunc {
	return Handler(secret, func(event string, payload *GitHubPayload, req *http.Request) error {
		// Log webhook
		log.Println("Received", event, "for ", payload.Repository.Name)

		// You'll probably want to do some real processing
		log.Println("Can clone repo at:", payload.Repository.CloneURL)

		// All is good (return an error to fail)
		return nil
	})
}
