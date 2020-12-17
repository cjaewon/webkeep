package main

import "github.com/cjaewon/webkeep/commands"

func main() {
	rootCmd := commands.NewCmdRoot()

	rootCmd.Execute()
}
