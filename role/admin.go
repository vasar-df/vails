package role

import (
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/vasar-network/vails"
)

// Admin represents the role specification for the admin role.
type Admin struct{}

// Name returns the name of the role.
func (Admin) Name() string {
	return "admin"
}

// Chat returns the formatted chat message using the name and message provided.
func (Admin) Chat(name, message string) string {
	return text.Colourf("<grey>[<dark-aqua>Admin</dark-aqua>]</grey> <dark-aqua>%s</dark-aqua><dark-grey>:</dark-grey> <dark-aqua>%s</dark-aqua>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Admin) Tag(name string) string {
	return text.Colourf("<dark-aqua>%s</dark-aqua>", name)
}

// Inherits returns the role that this role inherits from.
func (Admin) Inherits() vails.Role {
	return Mod{}
}
