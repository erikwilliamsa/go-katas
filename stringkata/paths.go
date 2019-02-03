package stringkata

import (
	"strings"
)

func ChangePath(current, destination string) string {

	curdirs, destdirs := splitAndIgnoreEmpty(current), splitAndIgnoreEmpty(destination)

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

func splitAndIgnoreEmpty(s string) []string {

	result := []string{}
	if dirs := strings.Split(s, "/"); len(dirs) > 0 && dirs[0] == "" {
		result = dirs[1:]

	} else {
		result = dirs
	}

	return result
}
