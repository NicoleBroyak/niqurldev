package cli

import (
	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

func createApp() *grumble.App {
	app := grumble.New(&grumble.Config{
		Name: "NiQurl",
		Description: "\n\n*    .##....##.####..#######..##.....##.########..##......\n" +
			"*    .###...##..##..##.....##.##.....##.##.....##.##......\n" +
			"*    .####..##..##..##.....##.##.....##.##.....##.##......\n" +
			"*    .##.##.##..##..##.....##.##.....##.########..##......\n" +
			"*    .##..####..##..##..##.##.##.....##.##...##...##......\n" +
			"*    .##...###..##..##....##..##.....##.##....##..##......\n" +
			"*    .##....##.####..#####.##..#######..##.....##.########\n\n" +
			"            NiQurl - Simple URL Shortening App\n\n\n" +
			"NiQurl is running. \nTo shorten url use command " +
			"\"make https://example.com\"\nUse help command to read documentation\n\n\n",
		Prompt:                "NiQurl>",
		PromptColor:           color.New(color.BgMagenta, color.Bold, color.FgBlack),
		HelpHeadlineColor:     color.New(color.FgMagenta),
		HelpHeadlineUnderline: true,
		HelpSubCommands:       true,
		Flags: func(f *grumble.Flags) {
			f.Int("g", "generate-fake-users", 0, "help string")
		},
	})
	commands(app)
	app.OnInit(initialize)
	return app
}
