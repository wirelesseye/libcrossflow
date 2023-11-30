package sharespace

import (
	"libcrossflow/config"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type ShareSpace struct {
	Name   string
	Config config.ShareSpaceConfig
}

type FileInfo struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func GetShareSpace(name string) (ShareSpace, bool) {
	shareSpaceConfig, ok := config.GetConfig().GetRawData().ShareSpaces[name]
	if ok {
		return ShareSpace{
			Name:   name,
			Config: shareSpaceConfig,
		}, true
	} else {
		return ShareSpace{}, false
	}
}

func GetShareSpaceNames() []string {
	shareSpaceConfigs := config.GetConfig().GetRawData().ShareSpaces
	shareSpaceNames := make([]string, len(shareSpaceConfigs))

	i := 0
	for key := range shareSpaceConfigs {
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
		realPath = shareSpace.Config.Files[name]
	} else {
		name, relPath := split[0], split[1]
		rootPath := shareSpace.Config.Files[name]
		realPath = filepath.Join(rootPath, relPath)
	}

	return realPath
}

func (shareSpace ShareSpace) GetFileInfo(path string) (FileInfo, error) {
	if path == "" {
		return FileInfo{
			Type: "sharespace",
			Name: shareSpace.Name,
		}, nil
	}

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

		for name, path := range shareSpace.Config.Files {
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
