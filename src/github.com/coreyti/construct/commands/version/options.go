package version

import (
	"github.com/cppforlife/go-cli-ui/ui"
)

type Options struct {
	ui ui.UI
}

func NewOptions(ui ui.UI) *Options {
	return &Options{ui}
}

func (o *Options) Run() error {
	o.ui.PrintBlock([]byte("construct CLI Version: 0.1.0\n"))
	return nil
}
