package core

type Video struct {
	filepath string
}

func createVideo(filepath string) *Video {
	return &Video{filepath}
}

func (video *Video) getFilePath() string {
	return video.filepath
}

func (video *Video) open() {

}
