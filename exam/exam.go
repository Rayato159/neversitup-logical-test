package exam

import (
	"regexp"
)

func ValidatePincode(pin string) bool {
	match, err := regexp.MatchString(`^\d{6,}`, pin)
	if err != nil {
		return false
	}
	if !match {
		return false
	}

	dup := 0
	for i := 1; i < len(pin)-1; i++ {
		if pin[i] == pin[i-1] {
			dup++
		} else {
			dup = 0
		}
		if dup > 1 {
			return false
		}
	}

	series := 0
	for i := 1; i < len(pin)-1; i++ {
		if pin[i] == pin[i-1]-1 || pin[i] == pin[i-1]+1 {
			series++
		} else {
			series = 0
		}
		if series > 1 {
			return false
		}
	}

	checkList := make([]string, 0)
	count := 0
	for i := 0; i < len(pin)-1; i++ {
		for j := i + 1; j < len(pin)-1; j++ {
			if pin[j] == pin[j+1] {
				count++
				if count > 0 {
					checkList = append(checkList, string(pin[j]))
					count = 0
					break
				}
			}
		}
	}
	if len(checkList) > 2 {
		return false
	}
	return true
}
