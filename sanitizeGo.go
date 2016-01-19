package sanitize

import (
	"regexp"
)

func Email(s string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,30})@(\w{1,}).([a-z]{2,4})$`, s); !m {
		return false
	} else {
		return true
	}
}

func Name(s string) bool {
	if m, _ := regexp.MatchString(`^[a-zA-Z]+$`, s); !m {
		return false
	} else {
		return true
	}
}

func Select(s string, choices []string) bool {
	for _, v := range choices {
		if v == s {
			return true
		}
	}
	return false
}

func Radio(i int, choices []int) bool {
	for _, v := range choices {
		if v == i {
			return true
		}
	}
	return false
}
