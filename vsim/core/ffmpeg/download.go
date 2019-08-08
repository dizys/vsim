package ffmpeg

import (
	"archive/zip"
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const downloadUrl = "https://evermeet.cx/ffmpeg/ffmpeg-4.2.zip"

func Download() error {
	err := downloadFile("./ffmpeg.zip", downloadUrl)

	if err != nil {
		return err
	}

	fmt.Printf("Download finished. Unzipping...\n")
	err = unzip("./ffmpeg.zip", "./")
	_ = os.Remove("./ffmpeg.zip")

	fmt.Printf("Done.\n")
	return err
}

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 40))

	fmt.Printf("\rDownloading decoder... %s complete (total 23MB)", humanize.Bytes(wc.Total))
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	fmt.Print("\n")

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
