package slices

import (
	"strings"
)

func SplitString(s, sep string) []string {
	if s != "" {
		return strings.Split(s, sep)
	} else {
		return []string{}
	}
}
