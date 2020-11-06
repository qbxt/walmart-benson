package commands

import "github.com/bwmarrin/discordgo"

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "wip")
}
