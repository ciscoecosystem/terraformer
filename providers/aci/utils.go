package aci

import (
	"regexp"
	"strconv"
	"strings"

	container "github.com/Jeffail/gabs"
)

func GetMOName(dn string) string {
	arr := strings.Split(dn, "/")
	// Get the last element
	last_ele := arr[len(arr)-1]
	// split on -
	dash_split := strings.Split(last_ele, "-")
	// join except first element as that will be rn
	return strings.Join(dash_split[1:], "-")
}

func filterChildrenDn(dn, parentDns string) string {
	parentDnList := strings.Split(parentDns, ":")

	for _, parentDn := range parentDnList {
		if strings.HasPrefix(dn, parentDn) {
			return dn
		}
	}

	return ""
}

func StrtoInt(s string, startIndex int, bitSize int) (int64, error) {
	return strconv.ParseInt(s, startIndex, bitSize)
}

func stripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(word, "\""), "\"")
	}
	return word
}

func G(cont *container.Container, key string) string {
	return stripQuotes(cont.S(key).String())
}

func GetParentDn(dn string, rn string) string {
	arr := strings.Split(dn, rn)
	return arr[0]
}

func replaceSpecialCharsDn(dn string) string {
	doubleDash := regexp.MustCompile("[/]")
	removeChars := regexp.MustCompile("[\\[\\]]")
	res := doubleDash.ReplaceAllString(dn, "__")
	res = removeChars.ReplaceAllString(res, "")
	return res
}
