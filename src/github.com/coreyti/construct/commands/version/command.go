package version

import (
	"github.com/spf13/cobra"
)

func NewCommand(o *Options) *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Prints construct CLI version",
		RunE: func(_ *cobra.Command, _ []string) error {
			return o.Run()
		},
	}

	return command
}
