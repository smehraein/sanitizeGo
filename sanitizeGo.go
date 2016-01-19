/*
Package sanitizeGo implements simple input verification tools for a basic go server.
*/
package sanitizeGO

import (
	"regexp"
)

// Checks if input string is a valid email
func Email(s string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,30})@(\w{1,}).([a-z]{2,4})$`, s); !m {
		return false
	} else {
		return true
	}
}

// Checks if input string is a valid English name
func Name(s string) bool {
	if m, _ := regexp.MatchString(`^[a-zA-Z]+$`, s); !m {
		return false
	} else {
		return true
	}
}

// Given a slice of choices, checks that a given string is one of them
func Select(s string, choices []string) bool {
	for _, v := range choices {
		if v == s {
			return true
		}
	}
	return false
}

// Given a slice of ints, checks that a given int is one of them
func Radio(i int, choices []int) bool {
	for _, v := range choices {
		if v == i {
			return true
		}
	}
	return false
}
