package add

import (
	"fmt"
	"os"
	"strings"

	"github.com/cppforlife/go-cli-ui/ui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type Options struct {
	ui    ui.UI
	cache string
}

func NewOptions(ui ui.UI) *Options {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cache := strings.Join([]string{home, ".construct", "cache"}, "/")
	if _, err := os.Stat(cache); os.IsNotExist(err) {
		os.MkdirAll(cache, os.ModePerm)
	}

	return &Options{
		ui:    ui,
		cache: cache,
	}
}

func (o *Options) Set(command *cobra.Command) {
}
