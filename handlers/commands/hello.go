package commands

import (
	"bensonMC/constants"
	"bensonMC/logger"
	"fmt"
	"github.com/Lukaesebrot/mojango"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func IAm(s *discordgo.Session, m *discordgo.MessageCreate) {
	mojangClient := mojango.New()
	messageArgs := strings.Split(m.Content, " ")

	userIGN := messageArgs[1]

	userUUID, err := mojangClient.FetchUUID(userIGN)
	if err != nil {
		logger.Error("could not fetch UUID from IGN", err, logrus.Fields{"IGN": userIGN})
		failed(s, m)
		return
	}

	// Set nickname
	if err := s.GuildMemberNickname(m.GuildID, m.Author.ID, userIGN); err != nil {
		logger.Error("could not set nickname", err, logrus.Fields{"userID": m.Author.ID})
		failed(s, m)
	}

	// Add role
	if err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, constants.MC_ROLE); err != nil {
		logger.Error("could not add role", err, logrus.Fields{"userID": m.Author.ID})
		failed(s, m)
		return
	}

	// Send success embed
	e := &discordgo.MessageEmbed{}
	e.Title = fmt.Sprintf("%s#%s | %s", m.Author.Username, m.Author.Discriminator, userIGN)
	e.Color = 16713016
	e.Image = &discordgo.MessageEmbedImage{URL: fmt.Sprintf(constants.CRAFATAR_BODYRENDER_FORMATSTRING, userUUID)}
	e.Timestamp = time.Now().UTC().Format(time.RFC3339)
	if _, err := s.ChannelMessageSendEmbed(m.ChannelID, e); err != nil {
		logger.Error("could not send success message", err, nil)
		failed(s, m)
		return
	}
	_ = s.MessageReactionAdd(m.ChannelID, m.ID, constants.CHECK_EMOJI)
}

func failed(s *discordgo.Session, m *discordgo.MessageCreate) {
	if err := s.MessageReactionAdd(m.ChannelID, m.ID, constants.X_EMOJI); err != nil {
		logger.Error("could not add reaction", err, nil)
	}
}