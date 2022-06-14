package cli

import (
	"fmt"
	"log"
	"redishandler"

	"github.com/desertbit/grumble"
)

func commands(app *grumble.App) {
	settime(app)
	setlen(app)
	makeurl(app)
	settings(app)
}

func settime(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name:    "settime",
		Aliases: []string{"SETTIME"},
		Help:    "set USER_WAIT_TIME variable, value in ms, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("settime", "settime")
		},
		Run: cmdSetTime,
	})
}

func setlen(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name:    "setlen",
		Aliases: []string{"SETLEN"},
		Help:    "set SHORT_URL_LEN variable, must be > 0",
		Args: func(a *grumble.Args) {
			a.Int("setlen", "setlen")
		},
		Run: cmdSetLen,
	})
}

func makeurl(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name:    "make",
		Aliases: []string{"MAKE"},
		Help:    "shorten url",
		Args: func(a *grumble.Args) {
			a.String("url", "url")
		},
		Run: cmdMake,
	})
}

func settings(app *grumble.App) {
	app.AddCommand(&grumble.Command{
		Name: "settings",
		Aliases: []string{"setup",
			"config",
			"SETTINGS",
			"SETUP",
			"CONFING"},
		Help: "show settings",
		Run:  cmdSettings,
	})
}

func cmdSettings(c *grumble.Context) error {
	redishandler.CheckSettings()
	fmt.Println("Current settings")
	fmt.Printf(
		"short url length: %v characters\n",
		redishandler.Client.Get(redishandler.Ctx, "SHORT_URL_LEN"),
	)
	fmt.Printf(
		"user wait time: %v s \n",
		redishandler.Client.Get(redishandler.Ctx, "USER_WAIT_TIME"),
	)
	return nil
}

func cmdSetVar(cmd, v string, min, max int, c *grumble.Context) error {
	if c.Args.Int(cmd) < min || c.Args.Int(cmd) > max {
		err := fmt.Errorf("%v variable must be between %v and %v", v, min, max)
		return err
	}
	redishandler.Client.Set(redishandler.Ctx, v, c.Args.Int(cmd), 0)
	log.Printf("%v set to %v\n", v, c.Args.Int(cmd))
	return nil
}

func cmdSetLen(c *grumble.Context) error {
	return cmdSetVar("setlen", "SHORT_URL_LEN", 1, 20, c)
}

func cmdSetTime(c *grumble.Context) error {
	return cmdSetVar("settime", "USER_WAIT_TIME", 1, 1<<20, c)
}

func cmdMake(c *grumble.Context) error {
	url := c.Args.String("url")
	redishandler.CheckSettings()
	if redishandler.CheckZSet(url, "longurl") == true {
		redishandler.PrintShortURL(url)
		return nil
	}
	user := redishandler.RandomUser()
	if redishandler.CheckWaitTime(user) == false {
		shrt := redishandler.ShortURL(url)
		redishandler.InsertURL(url, shrt, user)
		return nil
	}
	return nil
}
