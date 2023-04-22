package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rottenhaddock",
	Short: "Roten Haddock is a tool for detecting phishing domains",
}

func Execute() error {
	return rootCmd.Execute()
}
