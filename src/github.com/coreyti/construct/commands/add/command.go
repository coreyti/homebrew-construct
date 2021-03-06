package add

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCommand(o *Options) *cobra.Command {
	command := &cobra.Command{
		Use: "add <name> <repo>",

		Run: func(command *cobra.Command, args []string) {
			// TODO: validation
			name := args[0]
			repo := args[1]

			update := map[string]interface{}{
				"repositories": map[string]interface{}{
					name: repo,
				},
			}

			if err := viper.MergeConfigMap(update); err != nil {
				fmt.Printf("Merging config; error: %s", err)
				os.Exit(1)
			}

			if err := viper.WriteConfig(); err != nil {
				fmt.Printf("Writing config; error: %s", err)
				os.Exit(1)
			}

			// TODO: move to a separate command
			path := strings.Join([]string{o.cache, name}, "/")
			runner := NewRunner(path, repo)
			runner.Run()

			// TODO: move to a separate command
			runner.TmpInstall()
		},
	}

	o.Set(command)
	return command
}
