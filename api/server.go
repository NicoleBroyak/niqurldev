package main

import (
	"os"
	"path"

	"github.com/nicolebroyak/niqurl/internal/redishandler"

	"github.com/gin-gonic/gin"
)

func main() {
	redishandler.Start()
	defer redishandler.Client.Close()
	StartServer()
}

func StartServer() {
	server := gin.Default()
	serverPath, _ := os.Getwd()
	tmplPath := path.Join(serverPath, "api", "templates")
	server.LoadHTMLFiles(
		path.Join(tmplPath, "404.html"),
		path.Join(tmplPath, "inspectURL.html"),
	)
	server.GET("/!:url", inspectURL)
	server.GET("/:url", redirectURL)
	server.NoRoute(notFound)
	server.Run(":8081")
}
