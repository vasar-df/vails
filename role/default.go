package role

import "github.com/sandertv/gophertunnel/minecraft/text"

// Default represents the role specification for the default role.
type Default struct{}

// Name returns the name of the role.
func (Default) Name() string {
	return "default"
}

// Chat returns the formatted chat message using the name and message provided.
func (Default) Chat(name, message string) string {
	return text.Colourf("<grey>%s</grey><white>: %s</white>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Default) Tag(name string) string {
	return text.Colourf("<grey>%s</grey>", name)
}
