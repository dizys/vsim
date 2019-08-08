package cmd

import (
	"github.com/spf13/cobra"
	"vsim/core"
	"vsim/version"
)

func Run() error {
	var rootCmd = &cobra.Command{Use: "vsim [VIDEO_A] [VIDEO_B]",
		Short:   "vsim is a tool for calculating video similarity",
		Long:    "A convenient tool for calculating similarities between videos",
		Args:    cobra.MinimumNArgs(2),
		PreRunE: checkFFMPEGAvailabilityPreRunE,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var videoAPath = args[0]
			var videoBPath = args[1]

			videoA := core.CreateVideo(videoAPath)
			videoB := core.CreateVideo(videoBPath)

			comparer := &core.VideoComparer{videoA, videoB}

			err = comparer.OpenVideos()
			defer comparer.CloseVideos()

			if err != nil {
				return
			}

			err = comparer.CheckConsistency()

			if err != nil {
				return
			}

			err = comparer.Compare()

			return err
		},
		Version: version.Version}

	rootCmd.AddCommand(cmdSetup, cmdClean)

	return rootCmd.Execute()
}
