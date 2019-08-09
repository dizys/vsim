package core

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/panjf2000/ants"
	"go.uber.org/atomic"
	"sync"
	"vsim/utils"
)

type VideoComparer struct {
	VideoA *Video
	VideoB *Video
}

func (comparer *VideoComparer) OpenVideos() (err error) {
	err = comparer.VideoA.Open()

	if err != nil {
		return
	}

	err = comparer.VideoB.Open()

	if err != nil {
		return
	}

	return
}

func (comparer *VideoComparer) CloseVideos() {
	comparer.VideoA.Close()
	comparer.VideoB.Close()
}

func (comparer *VideoComparer) CheckConsistency() (err error) {
	length := comparer.VideoA.GetLength()

	if length <= 0 {
		return fmt.Errorf("video not opened yet")
	}

	if length != comparer.VideoB.GetLength() {
		return fmt.Errorf("inconsistent frame numbers between two videos")
	}

	frameA := comparer.VideoA.GetFrameAt(0)
	frameB := comparer.VideoB.GetFrameAt(0)

	if !frameA.IsLoaded() {
		err = frameA.Load()

		if err != nil {
			return
		}
	}

	if !frameB.IsLoaded() {
		err = frameB.Load()

		if err != nil {
			return
		}
	}

	imageA, err := frameA.GetImage()

	if err != nil {
		return
	}

	imageB, err := frameB.GetImage()

	if err != nil {
		return
	}

	if !utils.ImageBoundsMatch(imageA.Bounds(), imageB.Bounds()) {
		return fmt.Errorf("video sizes mismatched")
	}

	return
}

func (comparer *VideoComparer) GetLength() int {
	return comparer.VideoA.GetLength()
}

func (comparer *VideoComparer) CompareFrameAt(index int) (avgDiff float64, err error) {
	frameA := comparer.VideoA.GetFrameAt(index)
	frameB := comparer.VideoB.GetFrameAt(index)

	if !frameA.IsLoaded() {
		err = frameA.Load()

		if err != nil {
			return
		}
	}

	if !frameB.IsLoaded() {
		err = frameB.Load()

		if err != nil {
			return
		}
	}

	imageA, err := frameA.GetImage()

	if err != nil {
		return
	}

	imageB, err := frameB.GetImage()

	if err != nil {
		return
	}

	diff := utils.GetImageDiff(imageA, imageB)

	width := imageA.Bounds().Dx()
	height := imageA.Bounds().Dy()

	avgDiff = diff / float64(width*height)
	return
}

func (comparer *VideoComparer) Compare() (err error) {
	length := comparer.GetLength()

	var totalDiff float64

	bar := pb.StartNew(length)

	for i := 0; i < length; i++ {
		diff, err := comparer.CompareFrameAt(i)

		if err != nil {
			return err
		}

		totalDiff += diff

		bar.Increment()
	}

	bar.Finish()

	fmt.Printf("Total diff sum: %.2f\n", totalDiff)

	return
}

func (comparer *VideoComparer) CompareInPool(size int) {
	if size <= 0 {
		size = 4
	}

	length := comparer.GetLength()

	var totalDiff atomic.Float64

	bar := pb.StartNew(length)

	var group sync.WaitGroup

	pool, _ := ants.NewPoolWithFunc(size, func(i interface{}) {
		n := i.(int)

		diff, err := comparer.CompareFrameAt(n)

		if err != nil {
			panic(err)
		}

		totalDiff.Add(diff)

		bar.Increment()
		group.Done()
	})
	defer pool.Release()

	for i := 0; i < length; i++ {
		group.Add(1)
		_ = pool.Invoke(i)
	}

	group.Wait()

	bar.Finish()

	fmt.Printf("Total diff sum: %.2f\n", totalDiff.Load())
}
