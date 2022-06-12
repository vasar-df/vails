package worlds

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// Handler handles world events, such as sounds playing.
type Handler struct {
	world.NopHandler
}

// HandleSound ...
func (h *Handler) HandleSound(ctx *event.Context, s world.Sound, _ mgl64.Vec3) {
	if _, ok := s.(sound.Attack); ok {
		ctx.Cancel()
	}
}

// HandleLiquidHarden ...
func (h *Handler) HandleLiquidHarden(ctx *event.Context, _ cube.Pos, _, _, _ world.Block) {
	ctx.Cancel()
}
