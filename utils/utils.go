package utils

import (
	"strings"
)

// StrtoUC ...
func StrtoUC(str string) string {
	return strings.ToUpper(str)
}

// StrToTitle ..
func StrToTitle(str string) string {
	return strings.Title(str)
}

// FirstToUC first letter to upper case
func FirstToUC(str string) string {
	letters := []string{}
	for i := range str {
		//log.Println(1, l)
		if i == 0 {
			letters = append(letters, strings.ToUpper(string(str[i])))
		} else {
			letters = append(letters, string(str[i]))
		}
	}
	var cname string
	for _, l := range letters {
		cname = cname + l

	}
	return cname
}
