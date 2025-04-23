package main

import (
	"github.com/bwmarrin/discordgo"
)

var FunGunsCommand = Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "funguns",
		Description: "Fun funguns mode",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		_, err := RCON_DIAL.Execute(
			"sv_cheats 1; mp_autoteambalance 0; mp_limitteams 0; " +
				"mp_maxrounds 100; mp_roundtime_defuse 30; " +
				"sv_infinite_ammo 1; mp_startmoney 1000000; " +
				"mp_afterroundmoney 1000000",
		)

		if err != nil {
			Respond(i.Interaction, "Error: "+err.Error())
			return
		}

		Respond(i.Interaction, "Bots added")
	},
}
