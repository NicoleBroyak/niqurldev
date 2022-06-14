package main

import (
	"cli"
	"redishandler"
)

func main() {
	redishandler.Start()
	defer redishandler.Client.Close()
	cli.Start()
}
