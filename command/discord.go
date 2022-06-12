package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/vasar-network/vails/lang"
)

// discord is a command that outputs the Vasar discord link.
type discord struct{}

// Run ...
func (discord) Run(s cmd.Source, o *cmd.Output) {
	o.Print(lang.Translatef(locale(s), "command.discord"))
}
