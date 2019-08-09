package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"vsim/core"
	"vsim/version"
)

func Run() error {
	var poolSize int

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

			fmt.Printf("Decoding videos...\n")
			err = comparer.OpenVideos()
			defer comparer.CloseVideos()

			if err != nil {
				return
			}

			fmt.Printf("Checking consisitency between two videos...\n")
			err = comparer.CheckConsistency()

			if err != nil {
				return
			}

			fmt.Printf("Comparing frames...\n")

			if poolSize == 0 {
				err = comparer.Compare()
			} else {
				comparer.CompareInPool(poolSize)
			}

			return err
		},
		Version: version.Version}

	rootCmd.Flags().IntVarP(&poolSize, "pool", "p", 0, "coroutine pool size")
	rootCmd.AddCommand(cmdSetup, cmdClean)

	return rootCmd.Execute()
}
