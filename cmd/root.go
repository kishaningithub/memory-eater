package cmd

import (
	"fmt"
	"github.com/inhies/go-bytesize"
	"github.com/kishaningithub/memory-eater/pkg/memoryeater"
	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"time"
)

var (
	Version      string
	stepSize     string
	stepDuration string
	profileCmd   bool
)

var rootCmd = &cobra.Command{
	Use:     "memory-eater",
	Short:   "Eats memory at a given constant rate",
	Example: "memory-eater --step-size=100MB --step-duration=1s",
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		if profileCmd {
			defer profile.Start(profile.ProfilePath("."), profile.MemProfile).Stop()
		}
		size, err := bytesize.Parse(stepSize)
		if err != nil {
			return fmt.Errorf("unable to parse step size: %w", err)
		}
		duration, err := time.ParseDuration(stepDuration)
		if err != nil {
			return fmt.Errorf("unable to parse step duration: %w", err)
		}
		memoryeater.Eat(size, duration)
		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVar(&stepSize, "step-size", "1MB", `Size of memory to eat during every step. Valid byte units are "B", "KB", "MB", "GB", "TB", "PB" and "EB"`)
	rootCmd.Flags().StringVar(&stepDuration, "step-duration", "1s", `Duration of every step. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"`)
	rootCmd.Flags().BoolVar(&profileCmd, "profile", false, "profile this cli tool")
}
