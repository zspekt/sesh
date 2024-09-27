package tmuxinator

import (
	"strings"

	"github.com/joshmedeski/sesh/model"
)

func (t *RealTmuxinator) List() ([]*model.TmuxinatorConfig, error) {
	// passing the --newline -n flag results in more consistent output,
	// regardless of the amount of entries 'list' returns
	res, err := t.shell.ListCmd("tmuxinator", "list", "--newline")
	if err != nil {
		// NOTE: return empty list if error
		return []*model.TmuxinatorConfig{}, nil
	}
	// first line is not a tmuxinator proj, and last one is empty
	return parseTmuxinatorConfigsOutput(res[1 : len(res)-1])
}

func parseTmuxinatorConfigsOutput(list []string) ([]*model.TmuxinatorConfig, error) {
	sessions := make([]*model.TmuxinatorConfig, 0, len(list))
	for _, line := range list {
		if len(line) > 0 {
			session := &model.TmuxinatorConfig{
				Name: strings.TrimSpace(line),
			}
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}
