package role

import "github.com/sandertv/gophertunnel/minecraft/text"

// Nitro represents the role specification for the nitro role.
type Nitro struct{}

// Name returns the name of the role.
func (Nitro) Name() string {
	return "nitro"
}

// Chat returns the formatted chat message using the name and message provided.
func (Nitro) Chat(name, message string) string {
	return text.Colourf("<grey>[<gold>Nitro</gold>]</grey> <gold>%s</gold><grey>:</grey> <white>%s</white>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Nitro) Tag(name string) string {
	return text.Colourf("<gold>%s</gold>", name)
}
