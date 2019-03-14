package common

import "strings"

func StringsHas(target []string, src string) bool {
	for _, t := range target {
		if strings.TrimSpace(t) == src {
			return true
		}
	}
	return false
}
