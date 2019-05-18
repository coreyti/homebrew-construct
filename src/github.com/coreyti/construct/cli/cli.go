package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/coreyti/construct/commands/add"
	"github.com/coreyti/construct/commands/version"
	"github.com/cppforlife/go-cli-ui/ui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config string
var command = NewCLI()

// NewCLI returns a new "construct command"
func NewCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "construct",
		Short: "The Construct CLI.",
		Long:  `A CLI with utilities for MacOS workstations.`,
	}

	ui := ui.NewConfUI(ui.NewNoopLogger())

	cmd.AddCommand(add.NewCommand(add.NewOptions(ui)))
	cmd.AddCommand(version.NewCommand(version.NewOptions(ui)))

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the command.
func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// command.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.construct/config.yml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// command.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configDir := strings.Join([]string{home, ".construct"}, "/")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, os.ModePerm)
	}

	configFile := strings.Join([]string{configDir, "config.yml"}, "/")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		f, err := os.Create(configFile)
		if err != nil {
			fmt.Printf("Initializing config file; err: %s", err)
			os.Exit(1)
		}

		f.Close()
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Using config file; err: %s", err)
		os.Exit(1)
	}
}
