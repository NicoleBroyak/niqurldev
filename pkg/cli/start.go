package cli

import (
	"fmt"

	"github.com/desertbit/grumble"
)

func Start() {
	app := createApp()
	fmt.Println(app.Config().Description)
	grumble.Main(app)
}
