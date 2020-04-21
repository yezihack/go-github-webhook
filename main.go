package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	scriptBash string // 脚本执行路径,必须有效
	urlPath    string // url path
	port       int    // 端口
	secret     string // 密码
	isRun      = make(chan struct{}, 1)
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "PONG")
	})
	r.POST("/"+urlPath, WebHookLog(secret))
	cmd := cli.NewApp()
	cmd.Flags = Flags()
	cmd.Action = Action()
	cmd.BashComplete = func(c *cli.Context) {
		fmt.Println("complete", c.NArg(), c.NumFlags())
	}
	if err := cmd.Run(os.Args); err != nil {
		log.Fatalln(err)
		return
	}
	select {
	case <-isRun:
		fmt.Println("start run", fmt.Sprintf(":%d", port))
		if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
			log.Fatalln(err)
		}
	default:
		fmt.Println("over")
	}
}

// hook
func WebHookLog(secret string) gin.HandlerFunc {
	return Handler(secret, func(event string, payload *GitHubPayload, req *http.Request) error {
		// Log webhook
		log.Println("Received:", event, "for:", payload.Repository.Name)

		// You'll probably want to do some real processing
		log.Println("Can clone repo at:", payload.Repository.CloneURL)

		// All is good (return an error to fail)
		CallScript(scriptBash)
		return nil
	})
}

// call script bash function
func CallScript(path string) {
	cmd := exec.Command("/bin/bash", path)
	data, err := cmd.Output()
	if err != nil {
		log.Println("CallScript", err)
		return
	}
	log.Println(string(data))
}

// action
func Action() cli.ActionFunc {
	return func(c *cli.Context) error {
		// Check that the script path is valid
		scriptBash = c.String("bash")
		if scriptBash == "" {
			return fmt.Errorf("The script path is null")
		}
		// Check that the file is valid
		if !com.IsFile(scriptBash) {
			return fmt.Errorf("The script path not valid")
		}
		// Check that the path is empty
		urlPath = c.String("path")
		if urlPath == "" {
			return fmt.Errorf("The url path is empty")
		}
		port = c.Int("port")
		secret = c.String("secret")
		if strings.TrimSpace(secret) == "" {
			return fmt.Errorf("The githut secret is empty")
		}
		isRun <- struct{}{}
		return nil
	}
}
