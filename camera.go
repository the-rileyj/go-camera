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
	AWBG, fullScreenPreview, horizontalFlip bool
	previewPicture, verticalFlip            bool
	videoStablization, ROI, simpleCapture   bool
	//Photo Specific
	demo, keypress, latest, photoVerbose, raw bool
	signal, timeOut                           bool
	//Video Specific
	inline, timed, videoVerbose, penc bool
	//General Floats
	blueAWBG, redAWBG float64
	//General Int32
	brightness, cameraSelection         int32
	channelU, channelY, contrast, EV    int32
	ISO, mode, opacity, px, py, pw, ph  int32
	saturation, sharpness, shutterSpeed int32
	//Photo Specific Int32
	photoWidth, photoHeight, jpgQuality int32
	timeLength, timeOutLength           int32
	//Video Specific Int32
	bitRate, frameRate, videoWidth      int32
	videoHeight, quantisation, h264Prof int32
	timeOn, timeOff                     int32
	//General Strings
	Annotate, AWB, dynamicRangeCompression string
	fileName, fileType, meteringMode       string
	savePath, simpleCommand                string
	//Photo Specific Strings
	photoEncoding, latestFileName, thumb string
	exif                                 string
	//Video Specific Strings
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
