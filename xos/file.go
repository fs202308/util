package xos

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	MODE_DIR  = 0
	MODE_FILE = 1
)

// 返回下层所有文件或目录
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

// 返回符合特定后缀的文件 或者目录
func ListSubFilesRecur(path string, suffix string, mode int) (files []string, err error) {
	files = make([]string, 0, 30)
	err = filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if mode == MODE_FILE {
			if !fi.IsDir() && strings.HasSuffix(fi.Name(), suffix) {
				files = append(files, filename)
			}
			return nil
		} else if mode == MODE_DIR {
			if fi.IsDir() {
				files = append(files, filename)
			}
		}
		return nil
	})
	return files, err
}

// 返回所有文件或目录 递归
func ListAllFilesRecur(path string, mode int) (files []string, err error) {
	files = make([]string, 0, 30)
	err = filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if mode == MODE_FILE {
			if !fi.IsDir() {
				files = append(files, filename)
			}
			return nil
		} else if mode == MODE_DIR {
			if fi.IsDir() {
				files = append(files, filename)
			}
		}
		return nil
	})
	return files, err
}
