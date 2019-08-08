package ffmpeg

import (
	"os"
	"os/exec"
	"vsim/utils"
)

func hasLocalFFMPEG() (availability bool) {
	availability, _ = utils.FileExists("./ffmpeg")

	return
}

func hasGlobalFFMPEG() (availability bool) {
	availability = utils.IsCommandAvailable("ffmpeg")

	return
}

func IsAvailable() bool {
	if hasLocalFFMPEG() {
		return true
	}

	return hasGlobalFFMPEG()
}

func ConvertVideoToImages(src string, dest string) (err error) {
	cmd := exec.Command("./ffmpeg", "-v", "error", "-stats", "-i", src, "-f", "image2", "-q:v", "1", dest)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	return
}
