package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"vsim/core/ffmpeg"
)

var (
	cmdSetup = &cobra.Command{
		Use:   "setup",
		Short: "Automatically download and setup video decoder if missing",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			availability := ffmpeg.IsAvailable()

			if availability {
				fmt.Printf("Decoder already setup.\n")
				return
			}

			fmt.Printf("Start downloading ffmpeg...\n")
			err = ffmpeg.Download()
			return
		},
	}
)

func checkFFMPEGAvailabilityPreRunE(_ *cobra.Command, _ []string) (err error) {
	availability := ffmpeg.IsAvailable()

	if !availability {
		err = fmt.Errorf("no available decoder found. Please use `vsim setup` to download and setup ffmpeg locally. Or you could visit https://ffmpeg.org and setup ffmpeg manually as a global executable")
		return
	}

	return
}
