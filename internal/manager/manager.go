package manager

import (
	"os"
	"path/filepath"
	"strings"

	"errors"

	"github.com/gabriel-vasile/mimetype"
)

type (
	Converter interface {
		Run(src string, basename string, dest string) error
	}

	ConvertManager struct {
		converter  Converter
		watchpoint string
		dest       string
	}
)

func New(converter *Converter, watchpoint string, dest string) (*ConvertManager, error) {
	if f, err := os.Stat(dest); os.IsNotExist(err) || !f.IsDir() {
		return nil, errors.New("output destination path is invalid:" + dest)
	}

	return &ConvertManager{
		converter:  *converter,
		watchpoint: watchpoint,
		dest:       dest,
	}, nil

}

func (cm ConvertManager) Run(src string) error {
	f, err := os.Stat(src)
	if os.IsNotExist(err) {
		return errors.New("input file path is invalid:" + src)
	}

	if f.IsDir() {
		return nil
	}

	mtype, err := mimetype.DetectFile(src)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(mtype.String(), "image/") {
		return nil
	}

	sub := strings.Replace(filepath.Dir(src), cm.watchpoint, "", 1)
	dest := cm.dest

	if len(sub) > 0 {
		dest = filepath.Join(dest, sub)

		if f, err := os.Stat(dest); os.IsNotExist(err) || !f.IsDir() {
			if err := os.MkdirAll(dest, os.ModePerm); err != nil {
				return err
			}
		}
	}

	if err := cm.converter.Run(src, filepath.Base(src[:len(src)-len(filepath.Ext(src))]), dest); err != nil {
		return err
	}

	return os.Remove(src)
}
