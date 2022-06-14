package cli

import (
	"github.com/nicolebroyak/niqurldev/internal/cli"
	"github.com/nicolebroyak/niqurldev/tools/redishandler"
)

func main() {
	redishandler.Start()
	defer redishandler.Client.Close()
	cli.Start()
}
