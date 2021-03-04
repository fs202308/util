package xos

import (
	"io/ioutil"
)

const (
	MODE_DIR  = 0
	MODE_FILE = 1
)

func ListSubFiles(path string, mode int) ([]string, error) {
	var r []string
	d, err := ioutil.ReadDir(path)
	if err != nil {
		return r, err
	}
	for _, d := range d {
		if mode == MODE_DIR {
			if d.IsDir() {
				r = append(r, d.Name())
			}
		}
		if mode == MODE_FILE {
			if !d.IsDir() {
				r = append(r, d.Name())
			}
		}
	}
	return r, nil
}
