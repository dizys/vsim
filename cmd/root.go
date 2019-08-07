package cmd

import (
	"github.com/spf13/cobra"
	"vsim/version"
)

func Run() error {
	var rootCmd = &cobra.Command{Use: "vsim [VIDEO_A] [VIDEO_B]",
		Short: "vsim is a tool for calculating video similarity",
		Long:  "A convenient tool for calculating similarities between videos",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var videoAPath = args[0]
			var videoBPath = args[1]

			return nil
		},
		Version: version.Version}

	return rootCmd.Execute()
}
