package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
)

var (
	reasonList []string

	discordBotToken = os.Getenv("DISCORD_TOKEN")
	commands        = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "no",
			Description: "returns various negative responses",
		},
	}
)

func main() {
	slog.Info("init start reading reasons.json")
	data, err := os.ReadFile("reasons.json")

	if err != nil {
		fmt.Println("error: Error reading file:", err)
		os.Exit(1)
	}

	dec := json.Unmarshal([]byte(data), &reasonList)
	if dec != nil {
		fmt.Println("error: Error unmarshaling:", dec)
		return
	}

	// disgo code here
	slog.Info("starting bot...")
	slog.Info("disgo version", slog.String("version", disgo.Version))

	client, err := disgo.New(discordBotToken,
		bot.WithDefaultGateway(),
		bot.WithEventListenerFunc(commandListener),
	)
	if err != nil {
		slog.Error("error while building disgo instance", slog.Any("err", err))
		return
	}

	defer client.Close(context.TODO())

	if _, err = client.Rest().SetGlobalCommands(client.ApplicationID(), commands); err != nil {
		slog.Error("error while registering commands", slog.Any("err", err))
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		slog.Error("error while connecting to gateway", slog.Any("err", err))
	}

	slog.Info("bot is now running.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s

}
