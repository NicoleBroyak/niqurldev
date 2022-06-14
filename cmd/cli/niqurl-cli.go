package cli

import (
	"github.com/nicolebroyak/niqurldev/pkg/cli"
	"github.com/nicolebroyak/niqurldev/tools/redishandler"
)

func main() {
	redishandler.Start()
	defer redishandler.Client.Close()
	cli.Start()
}
