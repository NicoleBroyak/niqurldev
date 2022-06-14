package main

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/nicolebroyak/niqurldev/api"
	"github.com/nicolebroyak/niqurldev/tools/redishandler"
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
	server.GET("/!:url", api.InspectURL)
	server.GET("/:url", api.RedirectURL)
	server.NoRoute(api.NotFound)
	server.Run(":8081")
}
