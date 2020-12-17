package commands

import (
	"github.com/spf13/cobra"

	"github.com/cjaewon/webkeep/utils/parser"
)

// StartOptions ...
type StartOptions struct {
	URLs []string
}

// NewCmdStart generates start cmd
func NewCmdStart() *cobra.Command {
	opts := &StartOptions{}

	cmd := &cobra.Command{
		Use:   "start [<url> ..]",
		Short: "Archive a website",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.URLs = args

			return startRun(opts)
		},
	}

	return cmd
}

func startRun(opts *StartOptions) error {
	for _, url := range opts.URLs {
		text, err := parser.Fetch(url)
		if err != nil {
			return err
		}

	}

	return nil
}
