package core

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func CreateVideoFrame(imagePath string) *VideoFrame {
	return &VideoFrame{imagePath, nil}
}

type VideoFrame struct {
	imagePath string
	image     image.Image
}

func (frame *VideoFrame) Load() (err error) {
	if frame.IsLoaded() {
		return
	}

	file, err := os.Open(frame.imagePath)
	defer file.Close()

	if err != nil {
		return
	}

	frame.image, _, err = image.Decode(file)

	return err
}

func (frame *VideoFrame) Unload() {
	frame.image = nil
}

func (frame *VideoFrame) IsLoaded() bool {
	return frame.image != nil
}

func (frame *VideoFrame) GetImage() (image image.Image, err error) {
	if !frame.IsLoaded() {
		err = fmt.Errorf("image not loaded yet")
		return
	}

	return frame.image, nil
}
