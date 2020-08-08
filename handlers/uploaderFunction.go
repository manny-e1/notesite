package handlers

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(mf *multipart.File, name string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "asset","files", name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	io.Copy(file, *mf)
	return nil
}

func Trimmer(name string) string{
	var m int
	for k, v := range name {
		if string(v) == "\\" {
			m = k + 1
		}
	}
	name = string(name[m:])
	return name
}