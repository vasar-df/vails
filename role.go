package vails

// Role represents a role in-game. These can vary, and are used to specify the permissions of the user. It also contains
// the name of the role, prefix, and colour.
type Role interface {
	// Name returns the name of the role, for example "Admin".
	Name() string
	// Chat returns the formatted chat message using the name and message provided.
	Chat(name, message string) string
	// Tag returns the formatted name-tag using the name provided.
	Tag(name string) string
}

// HeirRole represents a role that inherits from another role.
type HeirRole interface {
	// Inherits returns the role that this role inherits from.
	Inherits() Role
}
