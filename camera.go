package camera

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	STILL      = "raspistill"
	HFLIP      = "-hf"
	VFLIP      = "-vf"
	OUTFLAG    = "-o"
	TIME_STAMP = "2006-01-02_15:04::05"
)

type Camera struct {
	horizontalFlip bool
	verticalFlip   bool
	fileName       string
	fileType       string
	savePath       string
}

func New(path, name, fType string) *Camera {
	if name == "" {
		name = time.Now().Format(TIME_STAMP)
	}
	if fType == "" {
		fType = ".jpg"
	}
	return &Camera{false, false, name, fType, path}
}

func (c *Camera) Hflip(b bool) {
	c.horizontalFlip = b
}

func (c *Camera) Vflip(b bool) {
	c.verticalFlip = b
}

func (c *Camera) Capture() (string, error) {
	args := make([]string, 0)
	args = append(args, OUTFLAG)
	fullPath := c.fileName
	if c.savePath != "" {
		fullPath = filepath.Join(c.savePath, c.fileName)
	}
	args = append(args, fullPath)
	if c.horizontalFlip {
		args = append(args, HFLIP)
	}
	if c.verticalFlip {
		args = append(args, VFLIP)
	}
	cmd := exec.Command(STILL, args...)
	_, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Wait()
	return fullPath, nil
}
