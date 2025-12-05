package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dolmen-go/codegen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/KarolosLykos/advent-of-code-gen/internal/config"
	"github.com/KarolosLykos/advent-of-code-gen/internal/templates"
)

var (
	yearFlag int
	dayFlag  int
)

func NewGenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate new puzzle",
		Long:  "Gen Generate a puzzle from year and day inputs",
		RunE:  genCmd,
	}

	cmd.Flags().IntVarP(&yearFlag, "year", "y", time.Now().Year(), "aoc gen [-y year]")
	cmd.Flags().IntVarP(&dayFlag, "day", "d", time.Now().Day(), "aoc gen [-d day]")

	return cmd
}

func genCmd(cmd *cobra.Command, _ []string) error {
	if err := validateDates(); err != nil {
		return err
	}

	cfg, err := config.GetUserConfig()
	if err != nil {
		logrus.Error("could not get config file")
	}

	if _, err = os.Stat(cfg.ProjectDir); err != nil {
		return err
	}

	if err = cdToProject(cfg.ProjectDir); err != nil {
		logrus.Errorf("could not change directory: %v", err)

		return err
	}

	start, end := 1, 25
	if cmd.Flags().Changed("day") {
		start, end = dayFlag, dayFlag
	}

	for i := start; i < end; i++ {
		unlocked, err := checkIfDayIsUnlocked(cmd.Context(), cfg.Session, yearFlag, i)
		if err != nil {
			return err
		}

		if !unlocked {
			logrus.Infof("puzzle %d/%d is not available yet or doesn't exist", yearFlag, i)

			return nil
		}

		yearDir := cfg.ProjectDir + "/" + strconv.Itoa(yearFlag)
		if err = os.Mkdir(yearDir, 0o750); err != nil && !errors.Is(err, os.ErrExist) {
			logrus.Errorf("error creating directory: %v", err)

			return err
		}

		dayDir := yearDir + "/" + formatDay(i)
		if err = os.Mkdir(dayDir, 0o750); err != nil && !errors.Is(err, os.ErrExist) {
			logrus.Errorf("error creating directory: %v", err)

			return err
		}

		if err = createMainFile(dayDir, yearFlag, i); err != nil {
			logrus.Error(err)

			return err
		}

		if err = createTestFile(dayDir, yearFlag, i); err != nil {
			logrus.Error(err)

			return err
		}

		inputFile, err := getInputFile(cmd.Context(), cfg.Session, yearFlag, i)
		if err != nil {
			logrus.Errorf("could not get input file %d/%d", yearFlag, i)

			return err
		}

		f, err := os.Create(dayDir + "/input.txt")
		if err != nil {
			logrus.Error(err)
			return err
		}

		if _, err = f.WriteString(inputFile); err != nil {
			return err
		}
	}

	return nil
}

// validateDates validates years and days and setting default day.
func validateDates() error {
	if yearFlag <= 2020 || yearFlag > time.Now().Year() {
		return fmt.Errorf("invalid year: %d", yearFlag)
	}

	if dayFlag < 1 || dayFlag > 31 {
		return fmt.Errorf("invalid day: %d", dayFlag)
	}

	return nil
}

// formatDay zero pads single-digit days.
func formatDay(day int) string {
	yearStr := strconv.Itoa(day)
	if len(yearStr) == 1 {
		return "0" + yearStr
	}
	return yearStr
}

func createMainFile(dayDir string, yearFlag, dayFlag int) error {
	path := dayDir + "/main.go"
	if _, err := os.Stat(path); err == nil {
		logrus.Infof("solution %d/%d exists", dayFlag, yearFlag)

		return nil
	}

	tmpl := codegen.MustParse(templates.SolutionTemplate)

	return tmpl.CreateFile(path, map[string]interface{}{
		"Year": yearFlag,
		"Day":  formatDay(dayFlag),
	})
}

func createTestFile(dayDir string, yearFlag, dayFlag int) error {
	path := dayDir + "/main_test.go"
	if _, err := os.Stat(path); err == nil {
		logrus.Infof("test %d/%d exists", dayFlag, yearFlag)

		return nil
	}

	tmpl := codegen.MustParse(templates.TestTemplate)

	return tmpl.CreateFile(path, map[string]interface{}{
		"Year": yearFlag,
		"Day":  formatDay(dayFlag),
	})
}

func getInputFile(ctx context.Context, session string, year, day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	c := http.Client{Timeout: 3 * time.Second}

	cookie := &http.Cookie{Name: "session", Value: session, MaxAge: 0}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(cookie)
	req.Header.Set("User-Agent", "github.com/KarolosLykos/advent-of-code-gen by Karolos Lykos")

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("something went wrong %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func checkIfDayIsUnlocked(ctx context.Context, session string, year, day int) (bool, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	c := http.Client{Timeout: 3 * time.Second}

	cookie := &http.Cookie{Name: "session", Value: session, MaxAge: 0}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	req.AddCookie(cookie)
	req.Header.Set("User-Agent", "github.com/KarolosLykos/advent-of-code-gen by Karolos Lykos")

	resp, err := c.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}
