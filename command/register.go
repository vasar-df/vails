package command

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"golang.org/x/text/language"
)

// init will register all commands.
func init() {
	for _, c := range []cmd.Command{
		cmd.New("ping", "View the ping of yourself or another player.", []string{"ms"}, ping{}),
		cmd.New("discord", "Retrieve the link to our discord server.", nil, discord{}),
	} {
		cmd.Register(c)
	}
}

// locale returns the locale of a cmd.Source.
func locale(s cmd.Source) language.Tag {
	if p, ok := s.(*player.Player); ok {
		return p.Locale()
	}
	return language.English
}
