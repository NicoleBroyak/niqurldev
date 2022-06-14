package cli

import (
	"log"
	"redishandler"

	"github.com/desertbit/grumble"
)

func GFUflag(a *grumble.App, flags grumble.FlagMap) error {
	num := flags.Int("generate-fake-users")
	if num > 1000 || num < 1 {
		if num != 0 {
			log.Println(("Number of users has to be between 1 and 1000"))
		}
		return nil
	}
	return redishandler.GenerateFakeUsers(num)
}
