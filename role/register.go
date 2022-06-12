package role

import (
	"github.com/vasar-network/vails"
	"golang.org/x/exp/slices"
)

var (
	// roles contains all registered vails.Role implementations.
	roles []vails.Role
	// rolesByName contains all registered vails.Role implementations indexed by their name.
	rolesByName = map[string]vails.Role{}
)

// All returns all registered roles.
func All() []vails.Role {
	return roles
}

// Register registers a role to the roles list. The hierarchy of roles is determined by the order of registration.
func Register(role vails.Role) {
	roles = append(roles, role)
	rolesByName[role.Name()] = role
}

// ByName returns the role with the given name. If no role with the given name is registered, the second return value
// is false.
func ByName(name string) (vails.Role, bool) {
	role, ok := rolesByName[name]
	return role, ok
}

// Staff returns true if the role provided is a staff role.
func Staff(role vails.Role) bool {
	return Tier(role) >= Tier(Trial{})
}

// Tier returns the tier of a role based on its registration hierarchy.
func Tier(role vails.Role) int {
	return slices.IndexFunc(roles, func(other vails.Role) bool {
		return role == other
	})
}

// init registers all implemented roles.
func init() {
	Register(Operator{})
	Register(Default{})

	Register(Voter{})
	Register(Nitro{})

	Register(Plus{})

	Register(Media{})
	Register(Famous{})
	Register(Partner{})

	Register(Trial{})
	Register(Mod{})
	Register(Admin{})
	Register(Manager{})
	Register(Owner{})
}
