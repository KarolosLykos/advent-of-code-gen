package main

import (
	"fmt"
	"os"

	"github.com/KarolosLykos/advent-of-code-gen/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
