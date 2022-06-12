package role

import (
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Media represents the role structure for the Media role.
type Media struct{}

// Name returns the name of the role.
func (Media) Name() string {
	return "media"
}

// Chat returns the formatted chat message using the name and message provided.
func (Media) Chat(name, message string) string {
	return text.Colourf("<grey><i>[<aqua>Media</aqua>]</grey></i><aqua> %s</aqua><grey>:</grey> <white>%s</white>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Media) Tag(name string) string {
	return text.Colourf("<i><aqua>%s</aqua></i>", name)
}
