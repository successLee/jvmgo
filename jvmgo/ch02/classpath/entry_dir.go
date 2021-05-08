package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absPath string
}

func newDirEntry(path string) *DirEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absPath}
}

func (e *DirEntry) readClass(className string) ([]byte, Entry, error) {
	file := filepath.Join(e.absPath, className)
	data, err := ioutil.ReadFile(file)

	return data, e, err
}

func (e *DirEntry) String() string {
	return e.absPath
}
