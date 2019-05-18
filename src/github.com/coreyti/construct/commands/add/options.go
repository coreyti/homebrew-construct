package add

import (
	"github.com/cppforlife/go-cli-ui/ui"
	"github.com/spf13/cobra"
)

type Options struct {
	ui ui.UI
}

func NewOptions(ui ui.UI) *Options {
	return &Options{ui}
}

func (o *Options) Set(command *cobra.Command) {
}
