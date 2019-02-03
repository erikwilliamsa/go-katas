package stringkata

import (
	"strings"
)

func ChangePath(current, destination string) string {

	var curdirs, destdirs []string

	if dirs := strings.Split(current, "/"); len(dirs) > 0 && dirs[0] == "" {
		curdirs = dirs[1:]

	} else {
		curdirs = dirs
	}

	if dirs := strings.Split(destination, "/"); len(dirs) > 0 && dirs[0] == "" {
		destdirs = dirs[1:]

	} else {
		destdirs = dirs
	}

	lastMatchIndex := -1
	changedir := ""
	for i, v := range curdirs {
		if len(destdirs) > i && v == destdirs[i] {
			lastMatchIndex = i
		} else if len(curdirs) > 1 {

			changedir += "../"
		}
	}

	for i := lastMatchIndex + 1; i < len(destdirs); i++ {
		changedir += destdirs[i]

		if len(destdirs) > i+1 {
			changedir += "/"
		}
	}

	return changedir
}
