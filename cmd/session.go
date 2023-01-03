package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/KarolosLykos/advent-of-code-gen/internal/config"
)

var sessionFlag string

func NewSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Initialize AoC",
		Long:  "Initialize a new AoC project.",
		RunE:  sessionCmd,
	}

	cmd.Flags().StringVarP(&sessionFlag, "session cookie", "s", "", "aoc session [-s module] (-s session-value)")

	return cmd
}

func sessionCmd(_ *cobra.Command, _ []string) error {
	cfg, err := config.GetUserConfig()
	if err != nil {
		return err
	}

	fmt.Println(cfg)
	return nil
}
