package main

import (
	"os"
	"path"
	"github.com/NicoleBroyak/niqurldev/api"
	"github.com/gin-gonic/gin"
	"github.com/NicoleBroyak/niqurldev/tools/redishandler"
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
	server.GET("/!:url", api.inspectURL)
	server.GET("/:url", api.redirectURL)
	server.NoRoute(api.notFound)
	server.Run(":8081")
}
