package helper

import (
	"log"
	"regexp"
	"strconv"
)

func ParseFirstDigit(str *string) int {
	re := regexp.MustCompile("[0-9]+")
	digitMatch := re.FindStringSubmatch(*str)
	if digitMatch != nil {
		firstDigit, err := strconv.Atoi(digitMatch[0])
		if err != nil {
			log.Fatal(err)
		}
		return firstDigit
	}
	return 0
}
