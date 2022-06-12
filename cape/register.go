package cape

import (
	"github.com/vasar-network/vails"
)

var (
	// capes contains all registered capes.
	capes []vails.Cape
	// capesByName maps a cape's name to the cape itself.
	capesByName = map[string]vails.Cape{}
)

// All returns all registered capes.
func All() []vails.Cape {
	return capes
}

// Register registers the cape provided.
func Register(cape vails.Cape) {
	capes = append(capes, cape)
	capesByName[cape.Name()] = cape
}

// ByName returns the cape with the given name. If no cape with the given name is registered, the second return value is
// false.
func ByName(name string) (vails.Cape, bool) {
	cape, ok := capesByName[name]
	return cape, ok
}

// init registers all capes in the capes folder.
func init() {
	Register(vails.NewCape("Vasar Series 1", "regular/vasar_series_one.png", false))
	Register(vails.NewCape("Vasar Series 2", "regular/vasar_series_two.png", false))

	Register(vails.NewCape("Portal", "plus/portal.png", true))
	Register(vails.NewCape("Spicy", "plus/spicy.png", true))
	Register(vails.NewCape("Happy", "plus/happy.png", true))
	Register(vails.NewCape("Sad", "plus/sad.png", true))
	Register(vails.NewCape("Fold", "plus/fold.png", true))

	Register(vails.NewCape("Optifine Red", "plus/optifine_red.png", true))
	Register(vails.NewCape("Optifine Pink", "plus/optifine_pink.png", true))
	Register(vails.NewCape("Optifine Cyan", "plus/optifine_cyan.png", true))
	Register(vails.NewCape("Optifine Green", "plus/optifine_green.png", true))
	Register(vails.NewCape("Optifine Dark", "plus/optifine_dark.png", true))

	Register(vails.NewCape("Vlone", "plus/vlone.png", true))
	Register(vails.NewCape("Crow", "plus/crow.png", true))
	Register(vails.NewCape("AK", "plus/ak.png", true))
	Register(vails.NewCape("Evil", "plus/evil.png", true))
	Register(vails.NewCape("Drag", "plus/drag.png", true))
	Register(vails.NewCape("Bland", "plus/bland.png", true))
	Register(vails.NewCape("Pumpkin", "plus/pumpkin.png", true))
	Register(vails.NewCape("Wave", "plus/wave.png", true))
	Register(vails.NewCape("CBA", "plus/cba.png", true))
}
