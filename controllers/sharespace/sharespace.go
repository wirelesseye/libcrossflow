package sharespace

import (
	"errors"
	"libcrossflow/config"
	"os"
	"path/filepath"
	"strings"
)

type ShareSpace config.ShareSpace

type PathResult struct {
	Type string
	Files []string `json:"files,omitempty"`
}

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

func (shareSpace ShareSpace) GetPath(path string) (PathResult, error) {
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

	fi, err := os.Stat(realPath)
	if err != nil {
		return PathResult{}, err
	}

	switch mode := fi.Mode(); {
		case mode.IsDir():
			fileNames := []string{}
			entries, _ := os.ReadDir(realPath)
			for _, e := range entries {
				fileNames = append(fileNames, e.Name())
			}

			return PathResult {
				Type: "dir",
				Files: fileNames,
			}, nil
		case mode.IsRegular():
			return PathResult{
				Type: "file",
			}, nil
		default:
			return PathResult{}, errors.New("unexpected file mode")
	}
}
