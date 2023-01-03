package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/KarolosLykos/advent-of-code-gen/internal/config"
)

var (
	debugFlag bool
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "aoc",
		Short:   "AoC helper CLI",
		Long:    `A cli tool that generates, runs and tests advent of code puzzles.`,
		Example: `aoc`,
		Version: "v0.0.1",
	}

	cmd.PersistentFlags().BoolVar(&debugFlag, "debug", false, "aoc [command] -d")

	cmd.AddCommand(NewInitCmd()) // init command
	cmd.AddCommand(NewGenCmd())  // gen subcommand
	cmd.AddCommand(NewSessionCmd())

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
	cobra.OnInitialize(initConfig, initializeLogger)
}

func initializeLogger() {
	if debugFlag {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func initConfig() {
	if err := initConfigurationFile(); err != nil {
		logrus.Fatal(err)
	}
}

func initConfigurationFile() error {
	if err := config.SetViperConfig(); err != nil {
		return fmt.Errorf("could not set viper config: %v", err)
	}

	return viper.ReadInConfig()
}
