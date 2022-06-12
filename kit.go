package vails

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
)

// Kit contains all the items, armour, and effects obtained by a kit.
type Kit interface {
	// Items returns the items provided by the kit.
	Items(*player.Player) (items [36]item.Stack)
	// Armour contains the armour applied by using the kit.
	// The item stacks are ordered helmet, chestplate, leggings, and then boots.
	Armour(*player.Player) [4]item.Stack
	// Effects returns the effects applied by using the kit.
	Effects(*player.Player) []effect.Effect
}
