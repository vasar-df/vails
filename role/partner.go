package role

import (
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/vasar-network/vails"
)

// Partner represents the role structure for the partner role.
type Partner struct{}

// Name returns the name of the role.
func (Partner) Name() string {
	return "partner"
}

// Chat returns the formatted chat message using the name and message provided.
func (Partner) Chat(name, message string) string {
	return text.Colourf("<grey><i>[<dark-blue>Partner</dark-blue>]</grey></i><dark-blue> %s</dark-blue><grey>:</grey> <white>%s</white>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Partner) Tag(name string) string {
	return text.Colourf("<i><dark-blue>%s</dark-blue></i>", name)
}

// Inherits returns the role that this role inherits from.
func (Partner) Inherits() vails.Role {
	return Plus{}
}
