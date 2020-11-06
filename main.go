package main

import (
	"bensonMC/constants"
	"bensonMC/handlers"
	"bensonMC/logger"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.Init()

	bot, err := discordgo.New("Bot " + constants.TOKEN)
	if err != nil {
		logger.Error("could not make Discord instance", err, nil)
		return
	}

	bot.AddHandler(handlers.Ready)
	bot.AddHandler(handlers.MessageCreate)

	if err := bot.Open(); err != nil {
		logger.Error("could not open connection", err, nil)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = bot.Close()
}