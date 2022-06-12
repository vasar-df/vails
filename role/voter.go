package role

import "github.com/sandertv/gophertunnel/minecraft/text"

// Voter represents the role specification for the voter role.
type Voter struct{}

// Name returns the name of the role.
func (Voter) Name() string {
	return "voter"
}

// Chat returns the formatted chat message using the name and message provided.
func (Voter) Chat(name, message string) string {
	return text.Colourf("<grey>[<green>Voter</green>]</grey> <green>%s</green><grey>:</grey> <white>%s</white>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Voter) Tag(name string) string {
	return text.Colourf("<green>%s</green>", name)
}
