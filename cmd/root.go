package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg string

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cli-template",
		Short:   "This cli template shows nothing",
		Long:    `This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.`,
		Example: `cli-template`,
		Version: "v0.0.3", // <--VERSION-->
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(cmd.UsageString())

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&cfg, "config", "", "config file (default is $HOME/.cobra.yaml)")

	cmd.AddCommand(NewDateCmd())    // version subcommand
	cmd.AddCommand(NewVersionCmd()) // version subcommand

	return cmd
}

// Execute invokes the command.
func Execute() error {
	if err := NewRootCmd().Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfg != "" {
		fmt.Println("here", cfg)
		viper.SetConfigFile(cfg)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".aof-gen.yaml".
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".aof-gen.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(viper.ConfigFileNotFoundError.Error(viper.ConfigFileNotFoundError{}))
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
	fmt.Println(viper.Get("aoc_session"))
}
