package handlers

import (
	"bensonMC/constants"
	"bensonMC/handlers/commands"
	"bensonMC/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	lowerContent := strings.ToLower(m.Content)

	if strings.Contains(lowerContent, "benson") {
		go benson(s, m)
	}

	if strings.HasPrefix(lowerContent, constants.PREFIX) {

		if strings.HasPrefix(lowerContent, fmt.Sprintf("%siam", constants.PREFIX)) {
			commands.IAm(s, m)

			return
		} else if strings.HasPrefix(lowerContent, fmt.Sprintf("%shelp", constants.PREFIX)) {
			commands.Help(s, m)

			return
		}
	}
}

func benson(s *discordgo.Session, m *discordgo.MessageCreate) {
	memeChannels := []string{
		"774111207778287638", // general
		"774111225227378688", // botspam
	}

	for _, c := range memeChannels {
		if m.ChannelID == c {
			msg, err := s.ChannelMessageSend(m.ChannelID, constants.BENSON_URL)
			if err != nil {
				logger.Error("could not send benson image", err, nil)
				return
			}
			time.Sleep(10*time.Second)
			_ = s.ChannelMessageDelete(m.ChannelID, msg.ID)
		}
	}
}