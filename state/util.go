package state

import "strings"

func getFragments(str string, sep string) []string {
	if str == "" {
		return nil
	}
	return strings.Split(str, sep)
}
