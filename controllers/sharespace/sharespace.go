package sharespace

import "libcrossflow/config"

func GetShareSpaces() *map[string]config.ShareSpace {
	config := config.GetConfig()
	return &config.GetRawData().ShareSpaces
}

func GetShareSpaceNames() []string {
	shareSpaces := *GetShareSpaces()

	shareSpaceNames := make([]string, len(shareSpaces))
	i := 0
	for key := range shareSpaces {
		shareSpaceNames[i] = key
		i++
	}

	return shareSpaceNames
}