package commands

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// NewCmdRoot generates root cmd
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "webkeep <command> [flags]",
		Version: "0.0.1",
		Short:   "Webkeep is a archive tool for save a lot of websites to pdf light and fast",
		Long: heredoc.Doc(
			`A archive tool for save a lot of websites to pdf light and fast
			More info at https://github.com/cjaewon/webkeep`,
		),
		Example: heredoc.Doc(`
			$ webkeep start https://google.com
		`),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	cmd.AddCommand(NewCmdStart())

	return cmd
}
