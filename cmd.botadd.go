package main

import (
	"github.com/bwmarrin/discordgo"
	"go.kyoto.codes/zen/v3/slice"
)

var BotAddCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "botadd",
		Description: "Add bots",
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
				Description: "Number of bots to remove",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
			},
		},
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		args := GetCommandArgs(i.ApplicationCommandData().Options)

		team := "c"
		if args["team"] != nil {
			team = args["team"].StringValue()
		}

		num := 5
		if args["num"] != nil {
			num = int(args["num"].IntValue())
		}

		onecmd := "bot_add_ct;"
		if team == "t" {
			onecmd = "bot_add_t;"
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

		Respond(i.Interaction, "Bots added")
	},
}
