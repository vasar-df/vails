package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/vasar-network/vails/lang"
)

// ping is a command that shows your latency or another player's.
type ping struct {
	Targets []cmd.Target `optional:"" name:"target"`
}

// Run ...
func (p ping) Run(s cmd.Source, o *cmd.Output) {
	l := locale(s)
	if length := len(p.Targets); length > 1 {
		o.Error(lang.Translatef(l, "command.targets.exceed"))
	} else if length == 1 {
		if target, ok := p.Targets[0].(*player.Player); ok {
			o.Print(lang.Translatef(l, "command.ping.other", target.Name(), target.Latency().Milliseconds()*2))
		} else {
			o.Error(lang.Translatef(l, "command.target.unknown"))
		}
	} else {
		if u, ok := s.(*player.Player); ok {
			o.Print(lang.Translatef(l, "command.ping.self", u.Latency().Milliseconds()*2))
		} else {
			c, _ := cmd.ByAlias("ping")
			o.Error("Usage:§e", c.Usage())
		}
	}
}
