package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {

	// composite
	if strings.Contains(path, pathListSeparator) {
	}

	if strings.Contains(path, ".jar") ||
		strings.Contains(path, ".JAR") ||
		strings.Contains(path, ".zip") ||
		strings.Contains(path, ".ZIP") {
		return newZipEntry(path)
	}

	// wild
	if strings.Contains(path, "*") {

	}

	return newDirEntry(path)
}
