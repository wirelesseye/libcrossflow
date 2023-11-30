package sharespace

import (
	"libcrossflow/config"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type ShareSpace config.ShareSpace

type FileInfo struct {
	Type string `json:"type"`
	Name string `json:"name"`
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
	slices.Sort(shareSpaceNames)
	return shareSpaceNames
}

func (shareSpace ShareSpace) GetRealPath(path string) string {
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

	return realPath
}

func (shareSpace ShareSpace) GetFileInfo(path string) (FileInfo, error) {
	realPath := shareSpace.GetRealPath(path)

	fi, err := os.Stat(realPath)
	if err != nil {
		return FileInfo{}, err
	}

	var ty string
	if fi.IsDir() {
		ty = "dir"
	} else {
		ty = "file"
	}

	return FileInfo{
		Type: ty,
		Name: fi.Name(),
	}, nil
}

func (shareSpace ShareSpace) ListFiles(path string) ([]FileInfo, error) {
	if path == "" {
		files := []FileInfo{}

		for name, path := range shareSpace.Files {
			fi, err := os.Stat(path)
			if err != nil {
				return []FileInfo{}, err
			}

			var ty string
			if fi.IsDir() {
				ty = "dir"
			} else {
				ty = "file"
			}

			files = append(files, FileInfo{
				Type: ty,
				Name: name,
			})
		}

		return files, nil
	}

	realPath := shareSpace.GetRealPath(path)

	files := []FileInfo{}
	entries, err := os.ReadDir(realPath)
	if err != nil {
		return []FileInfo{}, err
	}

	for _, e := range entries {
		var ty string
		if e.IsDir() {
			ty = "dir"
		} else {
			ty = "file"
		}

		files = append(files, FileInfo{
			Type: ty,
			Name: e.Name(),
		})
	}

	return files, nil
}
