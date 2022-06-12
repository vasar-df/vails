package role

import (
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/vasar-network/vails"
)

// Owner represents the role specification for the owner role.
type Owner struct{}

// Name returns the name of the role.
func (Owner) Name() string {
	return "owner"
}

// Chat returns the formatted chat message using the name and message provided.
func (Owner) Chat(name, message string) string {
	return text.Colourf("<grey>[<dark-purple>Owner</dark-purple>]</grey> <dark-purple>%s</dark-purple><dark-grey>:</dark-grey> <dark-purple>%s</dark-purple>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Owner) Tag(name string) string {
	return text.Colourf("<dark-purple>%s</dark-purple>", name)
}

// Inherits returns the role that this role inherits from.
func (Owner) Inherits() vails.Role {
	return Manager{}
}
