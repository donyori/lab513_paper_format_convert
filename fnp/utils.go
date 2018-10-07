package fnp

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var invalidFilenameCharacterPattern *regexp.Regexp

func init() {
	invalidFilenameCharacterPattern = regexp.MustCompile(`[/<>:"\\\|\?\*]`)
}

func parseNoRange(noRangeString string) (number1, number2 int32, err error) {
	idx := strings.IndexRune(noRangeString, '-')
	if idx < 0 {
		no, err := strconv.ParseInt(noRangeString, 10, 32)
		if err != nil {
			return 0, 0, err
		}
		number1 = int32(no)
		number2 = number1
	} else {
		noStr1 := noRangeString[:idx]
		noStr2 := noRangeString[idx+1:]
		no1, err := strconv.ParseInt(noStr1, 10, 32)
		if err != nil {
			return 0, 0, err
		}
		no2, err := strconv.ParseInt(noStr2, 10, 32)
		if err != nil {
			return 0, 0, err
		}
		number1 = int32(no1)
		number2 = int32(no2)
	}
	return number1, number2, nil
}

func formatNoRange(number1, number2 int32) string {
	if number1 == number2 {
		return fmt.Sprintf("%d", number1)
	} else {
		return fmt.Sprintf("%d-%d", number1, number2)
	}
}

func replaceInvalidFilenameCharacters(filename string) string {
	return invalidFilenameCharacterPattern.ReplaceAllLiteralString(
		filename, "_")
}
