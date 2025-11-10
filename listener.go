package main

import (
	"log/slog"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func commandListener(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()

	msg := messageGenerator(reasonList)

	if data.CommandName() == "no" {
		err := event.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent(msg).
			Build(),
		)
		if err != nil {
			slog.Error("error on sending response", slog.Any("err", err))
		}
	}
}
