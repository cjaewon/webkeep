package commands

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "webkeep",
		Version: "0.0.1",
		Short:   "Webkeep is a archive tool for save a lot of websites to pdf light and fast",
		Long: heredoc.Doc(
			`A archive tool for save a lot of websites to pdf light and fast
			More info at https://github.com/cjaewon/webkeep`,
		),
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

// Execute runs rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
