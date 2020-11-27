package project2

import (
	"regexp"
)

func judgeEno(e_no string) bool {
	match,_ := regexp.MatchString("[A|D]([0-9]{3})",e_no)

	return match
}


