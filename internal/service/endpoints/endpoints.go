package endpoints

import "github.com/bwmarrin/discordgo"

type API struct{}

func New() *API {
	return &API{}
}

func (a *API) HandleInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
	case discordgo.InteractionMessageComponent:
	default:
	}
}
