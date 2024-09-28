package tmuxinator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joshmedeski/sesh/model"
	"github.com/joshmedeski/sesh/shell"
)

func TestListConfigs(t *testing.T) {
	t.Run("List Tmuxinator Configs", func(t *testing.T) {
		mockShell := new(shell.MockShell)
		tmuxinator := &RealTmuxinator{shell: mockShell}
		mockShell.EXPECT().ListCmd("tmuxinator", "list", "--newline").Return([]string{
			"tmuxinator projects:",
			"dotfiles",
			"sesh",
			"home",
			"test1",
			"nix",
			"tmuxinator",
			"tmux",
			"sshnixpi",
		}, nil)
		expected := []*model.TmuxinatorConfig{
			{Name: "dotfiles"},
			{Name: "sesh"},
			{Name: "home"},
			{Name: "test1"},
			{Name: "nix"},
			{Name: "tmuxinator"},
			{Name: "tmux"},
			{Name: "sshnixpi"},
		}
		actual, err := tmuxinator.List()
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
