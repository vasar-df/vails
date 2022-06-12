package role

import "github.com/sandertv/gophertunnel/minecraft/text"

// Trial represents the role specification for the trial role.
type Trial struct{}

// Name returns the name of the role.
func (Trial) Name() string {
	return "trial"
}

// Chat returns the formatted chat message using the name and message provided.
func (Trial) Chat(name, message string) string {
	return text.Colourf("<grey>[<dark-yellow>Trial</dark-yellow>]</grey> <dark-yellow>%s</dark-yellow><dark-grey>:</dark-grey> <dark-yellow>%s</dark-yellow>", name, message)
}

// Tag returns the formatted name-tag using the name provided.
func (Trial) Tag(name string) string {
	return text.Colourf("<dark-yellow>%s</dark-yellow>", name)
}
