package cmd

import (
	"bytes"
	"fmt"
	"github.com/inhies/go-bytesize"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	Version      string
	stepSize     string
	stepDuration string
)

var rootCmd = &cobra.Command{
	Use:     "memory-eater",
	Short:   "Eats memory at a constant rate",
	Example: "memory-eater --step-size=100MB --step-duration=1s",
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		size, err := bytesize.Parse(stepSize)
		if err != nil {
			return fmt.Errorf("unable to parse step size: %w", err)
		}
		duration, err := time.ParseDuration(stepDuration)
		if err != nil {
			return fmt.Errorf("unable to parse step duration: %w", err)
		}
		var buff bytes.Buffer
		for {
			start := time.Now().UnixNano()
			buff.Write(bytes.Repeat([]byte{0}, int(size)))
			end := time.Now().UnixNano()
			durationForMemoryIncrease := end - start
			sleepDuration := time.Duration(int64(duration) - durationForMemoryIncrease)
			_, _ = fmt.Fprintf(os.Stderr, "\r Swallowed up %s              ", bytesize.ByteSize(buff.Len()))
			time.Sleep(sleepDuration)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVar(&stepSize, "step-size", "1MB", "")
	rootCmd.Flags().StringVar(&stepDuration, "step-duration", "1s", "")
}
