package files

import (
	"fmt"
	"io"
	"os"
	"path"
)

// CopyFile copies a single file from src to dst
func CopyFile(src, dst string) error {
	var err error
	var srcFile *os.File
	var dstFile *os.File
	var srcinfo os.FileInfo

	if srcFile, err = os.Open(src); err != nil {
		return err
	}
	defer srcFile.Close()

	if dstFile, err = os.Create(dst); err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// CopyDir copies a whole directory recursively
func CopyDir(src string, dst string) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcFile := path.Join(src, fd.Name())
		dstFile := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcFile, dstFile); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcFile, dstFile); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
