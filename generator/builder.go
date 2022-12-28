package generator

import (
	"time"

	"github.com/spf13/cobra"
)

// NewBuilder builds a cli usage generator, parsing particular
// flags (constant, max, interval) and building configuration
// while injecting the command (CMD) specific handler for generating
// events
//
// This exists to prevent duplication across the various event generator
// cli commands that will merely implement a unique handler
func NewBuilder(cmd *cobra.Command, h Handler) Generator {
	constant, _ := cmd.Flags().GetBool("constant")

	cfg := DefaultConfig()

	if !constant {
		cfg.MaxCount = 1
	}

	cfg.Handler = h

	if max, _ := cmd.Flags().GetInt("max"); max != 0 {
		cfg.MaxCount = max
	}

	if interval, _ := cmd.Flags().GetInt("interval"); interval != 0 {
		cfg.Interval = time.Second * time.Duration(interval)
	}

	return NewWitConfig(cfg)
}
