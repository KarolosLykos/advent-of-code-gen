package cmd

import (
	"github.com/spf13/cobra"

	"github.com/KarolosLykos/advent-of-code-gen/internal/config"
)

var sessionFlag string

func NewSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Session Cookie value",
		Long:  "Set the session cookie value from the https://adventofcode.com website",
		RunE:  sessionCmd,
	}

	cmd.Flags().StringVarP(&sessionFlag, "session cookie", "s", "", "aoc session [-s module] (-s session-value)")

	return cmd
}

func sessionCmd(_ *cobra.Command, _ []string) error {
	_, err := config.GetUserConfig()
	if err != nil {
		return err
	}

	return nil
}
