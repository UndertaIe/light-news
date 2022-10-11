package main

import "strings"

func trim(s string) string {
	return strings.TrimSpace(s)
}

var hotKeys = [...]string{"热", "hot", "*", "!"}

func isHot(s string) bool {
	for _, hk := range hotKeys {
		if strings.Contains(s, hk) {
			return true
		}
	}
	return false
}

// data({"1":"2"}) ==> {"1":"2"}
func jsonp2json(b []byte) []byte {
	n := len(b)
	var l = 0
	var r = n - 1
	for i := 0; i < n; i++ {
		if b[i] == '{' {
			l = i
			break
		}
	}
	for i := n - 1; i > 0; i-- {
		if b[i] == '}' {
			r = i
			break
		}
	}
	if r < l {
		return b
	}
	return b[l : r+1]
}
