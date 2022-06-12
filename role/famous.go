package role

import (
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/vasar-network/vails"
)

// Famous represents the role structure for the famous role.
type Famous struct{}

// Name returns the name of the role.
func (Famous) Name() string {
	return "famous"
}

// Chat returns the formatted chat message using the name and message provided.
func (Famous) Chat(name, message string) string {
	return text.Colourf("<grey><i>[<purple>Famous</purple>]</grey></i><purple> %s</purple><grey>:</grey> <white>%s</white>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Famous) Tag(name string) string {
	return text.Colourf("<i><purple>%s</purple></i>", name)
}

// Inherits returns the role that this role inherits from.
func (Famous) Inherits() vails.Role {
	return Media{}
}
