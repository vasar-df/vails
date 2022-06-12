package role

import (
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Operator represents the role specification for the operator role.
type Operator struct{}

// Name returns the name of the role.
func (Operator) Name() string {
	return "operator"
}

// Chat returns the formatted chat message using the name and message provided.
func (Operator) Chat(name, message string) string {
	return text.Colourf("<grey>%s</grey><white>: %s</white>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Operator) Tag(name string) string {
	return text.Colourf("<grey>%s</grey>", name)
}
