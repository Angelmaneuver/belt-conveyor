package converter

import (
	"errors"
	"path/filepath"

	"gocv.io/x/gocv"
)

type (
	Option struct {
		quality int
	}

	WebpConverter struct {
		params []int
	}
)

func New(option *Option) (*WebpConverter, error) {
	var params []int

	if option != nil {
		params = []int{gocv.IMWriteWebpQuality, option.quality}
	}

	converter := &WebpConverter{
		params: params,
	}

	return converter, nil
}

func NewOption(quality int) *Option {
	return &Option{
		quality: quality,
	}
}

func (c WebpConverter) Run(src string, basename string, dest string) error {
	img := gocv.IMRead(src, gocv.IMReadUnchanged)
	if img.Empty() {
		return errors.New("failed to read image:" + src)
	}
	defer img.Close()

	output := filepath.Join(dest, basename+".webp")

	if len(c.params) == 2 {
		if success := gocv.IMWriteWithParams(output, img, c.params); !success {
			return errors.New("failed to write image:" + output)
		}
	} else {
		if success := gocv.IMWrite(output, img); !success {
			return errors.New("failed to write image:" + output)
		}
	}

	return nil
}
