package helpers

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	Separator = string(filepath.Separator)
)

/*
444 r--r--r--
600 rw-------
644 rw-r--r--
666 rw-rw-rw-
700 rwx------
744 rwxr--r--
755 rwxr-xr-x
777 rwxrwxrwx
*/

func Mkdir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Remove(path string) error {
	return os.RemoveAll(path)
}

func Move(src string, dst string) error {
	return os.Rename(src, dst)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

func Pwd() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

func Home() (dir string, err error) {
	dir = os.Getenv("HOME")
	if len(dir) > 0 && dir[0] == '/' && IsDir(dir) {
		return dir, nil
	}

	return "", err
}

func Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func Basename(path string) string {
	return filepath.Base(path)
}

func Dir(path string) string {
	return filepath.Dir(path)
}

func Ext(path string) string {
	ext := filepath.Ext(path)
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}

func ExtName(path string) string {
	return strings.TrimLeft(Ext(path), ".")
}

func TempDir(names ...string) string {
	tempDir := "/tmp"
	if Separator != "/" || !IsDir(tempDir) {
		tempDir = os.TempDir()
	}

	path := tempDir
	for _, name := range names {
		path += Separator + name
	}
	return path
}

func Glob(pattern string, onlyNames bool) ([]string, error) {
	if list, err := filepath.Glob(pattern); err == nil {
		if onlyNames {
			array := make([]string, len(list))
			for k, v := range list {
				array[k] = Basename(v)
			}
			return array, nil
		}
		return list, nil
	} else {
		return nil, err
	}
}

func FileGetContent(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return "", nil
	}

	return string(fd), nil
}

func FilePutContent(path string, content string, isAppend bool) error {
	dir := Dir(path)
	if !IsDir(dir) {
		if err := Mkdir(dir); err != nil {
			return err
		}
	}
	flag := os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	perm := os.ModePerm
	if isAppend {
		flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
		perm = os.ModeAppend | os.ModePerm
	}
	f, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()

	data := []byte(content)
	if n, err := f.Write(data); err != nil {
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	return nil
}
