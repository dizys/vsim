package core

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"vsim/core/ffmpeg"
)

func CreateVideo(filepath string) *Video {
	return &Video{filepath, "", nil}
}

type Video struct {
	filepath   string
	decodedDir string
	frames     []*VideoFrame
}

func (video *Video) GetFilePath() string {
	return video.filepath
}

func (video *Video) Open() (err error) {
	_ = os.MkdirAll("./temp", os.ModePerm)

	basename := filepath.Base(video.filepath)
	ext := filepath.Ext(basename)
	filename := strings.TrimSuffix(basename, ext)

	video.decodedDir, err = ioutil.TempDir("./temp", filename)

	if err != nil {
		return err
	}

	outputFilepathPattern := filepath.Join(video.decodedDir, "image-%04d.jpg")

	err = ffmpeg.ConvertVideoToImages(video.filepath, outputFilepathPattern)

	if err != nil {
		return err
	}

	imagePaths, _ := filepath.Glob(filepath.Join(video.decodedDir, "image-*.jpg"))

	for _, path := range imagePaths {
		video.frames = append(video.frames, CreateVideoFrame(path))
	}

	return
}

func (video *Video) Close() {
	video.frames = nil
	_ = os.RemoveAll(video.decodedDir)
}

func (video *Video) IsOpened() bool {
	return video.frames != nil
}

func (video *Video) GetLength() int {
	return len(video.frames)
}

func (video *Video) GetFrameAt(index int) *VideoFrame {
	return video.frames[index]
}
