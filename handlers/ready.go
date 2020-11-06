package handlers

import (
	"bensonMC/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	logger.Info(fmt.Sprintf("%s#%s is now fully logged in and ready", s.State.User.Username, s.State.User.Discriminator), nil)
}
