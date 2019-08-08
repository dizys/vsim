package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	cmdClean = &cobra.Command{
		Use:   "clean",
		Short: "Clean temp directory to store decoded images",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return os.RemoveAll("./temp")
		},
	}
)
