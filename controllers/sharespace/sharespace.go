package sharespace

import (
	"libcrossflow/config"
	"os"
	"path/filepath"
	"strings"
)

type ShareSpace config.ShareSpace

func GetShareSpace(name string) (ShareSpace, bool) {
	shareSpaces := GetShareSpaces()
	val, ok := shareSpaces[name]
	return val, ok
}

func GetShareSpaces() map[string]ShareSpace {
	config := config.GetConfig()
	shareSpaces := map[string]ShareSpace{}

	configShareSpaces := config.GetRawData().ShareSpaces
	for name, shareSpace := range configShareSpaces {
		shareSpaces[name] = ShareSpace(shareSpace)
	}

	return shareSpaces
}

func GetShareSpaceNames() []string {
	shareSpaces := GetShareSpaces()

	shareSpaceNames := make([]string, len(shareSpaces))
	i := 0
	for key := range shareSpaces {
		shareSpaceNames[i] = key
		i++
	}

	return shareSpaceNames
}

func (shareSpace ShareSpace) ListFiles(path string) []string {
	split := strings.SplitN(path, "/", 2)

	var realPath string
	if len(split) < 2 {
		name := split[0]
		realPath = shareSpace.Files[name]
	} else {
		name, relPath := split[0], split[1]
		rootPath := shareSpace.Files[name]
		realPath = filepath.Join(rootPath, relPath)
	}

	fileNames := []string{}
	entries, _ := os.ReadDir(realPath)
	for _, e := range entries {
		fileNames = append(fileNames, e.Name())
	}
	return fileNames
}
