package camera

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	PHOTO      = "raspistill"
	VIDEO      = "raspivid"
	HFLIP      = "-hf"
	VFLIP      = "-vf"
	OUTFLAG    = "-o"
	PREVIEW    = false
	TIME_STAMP = "2006-01-02_15:04::05"
)

type Camera struct {
	//General Booleans
	AWBG, fullScreenPreview         bool
	horizontalFlip, previewPicture  bool
	verticalFlip, videoStablization bool
	ROI, simpleCapture              bool
	//Photo Specific
	captureVerbose, latest, timeOut bool
	raw, demo, penc                 bool
	//Video Specific

	//General Floats
	blueAWBG, redAWBG float64
	//General Int32
	brightness, cameraSelection         int16
	channelU, channelY, contrast, EV    int16
	ISO, mode, opacity, px, py, pw, ph  int16
	saturation, sharpness, shutterSpeed int16
	//Photo Specific
	photoWidth, photoHeight, jpgQuality int16
	timeLength, timeOutLength           int16
	//Video Specific
	frameRate int16
	//General Strings
	Annotate, AWB, dynamicRangeCompression string
	fileType, meteringMode, savePath       string
	simpleCommand                          string
	//Photo Specific
	latestFileName, fileName, thumb, encoding string
}

func New(path, name, fType string) *Camera {
	if name == "" {
		name = time.Now().Format(TIME_STAMP)
	}
	if fType == "" {
		fType = ".jpg"
	}
	return &Camera{horizontalFlip: false, verticalFlip: false, fileName: name, fileType: fType, savePath: path}
}

func (c *Camera) Hflip(b bool) *Camera {
	c.horizontalFlip = b
	return c
}

func (c *Camera) Vflip(b bool) *Camera {
	c.verticalFlip = b
	return c
}

func (c *Camera) Preview(b bool) *Camera {
	c.previewPicture = b
	return c
}

func (c *Camera) PreviewSize(x, y, w, h int16) *Camera {
	c.px = x
	c.py = y
	c.pw = w
	c.ph = h
	return c
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
	cmd := exec.Command(PHOTO, args...)
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
