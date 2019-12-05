//author: liu.yang02@ucarinc.com
//date: 20190830

package vstrings

import "strings"

func ContainsIn(items []string, item string) bool {
	if len(items) == 0 {
		return false
	}

	for _, n := range items {
		if item == n {
			return true
		}
	}

	return false
}

func ContainsAny(s string, test ...string) bool {
	for _, t := range test {
		if strings.Contains(s, t) {
			return true
		}
	}

	return false
}
