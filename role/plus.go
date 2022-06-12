package role

import "github.com/sandertv/gophertunnel/minecraft/text"

// Plus represents the role specification for the plus role.
type Plus struct{}

// Name returns the name of the role.
func (Plus) Name() string {
	return "plus"
}

// Chat returns the formatted chat message using the name and message provided.
func (Plus) Chat(name, message string) string {
	return text.Colourf("<grey>[<black>+</black>]</grey> <black>%s</black><grey>:</grey> <white>%s</white>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Plus) Tag(name string) string {
	return text.Colourf("<black>%s</black>", name)
}
