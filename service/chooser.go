package service

import (
	"strings"

	"gopkg.in/pipe.v2"

	"github.com/hirakiuc/daily-workflow/config"
)

type Chooser struct {
	Chooser     string
	ChooserOpts string
}

func NewChooser(conf *config.Config) *Chooser {
	return &Chooser{
		Chooser:     conf.Common.Chooser,
		ChooserOpts: conf.Common.ChooserOpts,
	}
}

func (s *Chooser) Choose(candidates []string) ([]string, error) {
	p := pipe.Line(
		pipe.Print(strings.Join(candidates, "\n")),
		//		pipe.Exec(s.Chooser, s.ChooserOpts),
	)

	output, err := pipe.CombinedOutput(p)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(output), "\n"), nil
}
