package main

import (
	"github.com/bwmarrin/discordgo"
	"go.kyoto.codes/zen/v3/slice"
)

var BotKickCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "botkick",
		Description: "Kick bots",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "team",
				Description: "Bot team",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: slice.Map(
					[]string{
						"c",
						"t",
					},
					func(s string) *discordgo.ApplicationCommandOptionChoice {
						return &discordgo.ApplicationCommandOptionChoice{
							Name:  s,
							Value: s,
						}
					}),
				Required: true,
			}, {
				Name:        "num",
				Description: "Number of bots to remove (or all)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
		},
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		args := GetCommandArgs(i.ApplicationCommandData().Options)

		team := "t"
		if args["team"] != nil {
			team = args["team"].StringValue()
		}

		var num = 5
		if args["num"] != nil {
			num = int(args["num"].IntValue())
		}

		onecmd := "bot_kick ct;"
		if team == "t" {
			onecmd = "bot_kick t;"
		}

		cmd := ""
		for k := 0; k < num; k++ {
			cmd += onecmd
		}

		_, err := RCON_DIAL.Execute(cmd)
		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
			return
		}

		Respond(i.Interaction, "Bots kicked")
	},
}
