package fnp

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type FilenamePatternNoRangeTitleTxt struct{}

type FilenameInfoNoRangeTitleTxt struct {
	Dir     string
	Numbers [2]int32
	T       string
}

var fpnrttRegexp *regexp.Regexp

func init() {
	fpnrttRegexp = regexp.MustCompile(
		`^[[:digit:]]+(\-[[:digit:]]+)??\..+\.txt$`)
}

func NewFilenamePatternNoRangeTitleTxt() *FilenamePatternNoRangeTitleTxt {
	return new(FilenamePatternNoRangeTitleTxt)
}

func (fpnrtt *FilenamePatternNoRangeTitleTxt) Parse(filename string) (
	filenameInfo FilenameInfo, err error) {
	if fpnrtt == nil {
		return nil, ErrNilFilenamePattern
	}
	dir, filename := filepath.Split(filename)
	if filename == "" {
		return nil, ErrEmptyFilename
	}
	if !fpnrttRegexp.MatchString(filename) {
		return nil, ErrFilenameNotMatchPattern
	}
	firstDot := strings.IndexRune(filename, '.')
	noStr := filename[:firstDot]
	no1, no2, err := parseNoRange(noStr)
	if err != nil {
		return nil, err
	}
	info := &FilenameInfoNoRangeTitleTxt{
		Dir:     dir,
		Numbers: [2]int32{no1, no2},
		T:       filename[firstDot+1 : len(filename)-4],
	}
	return info, nil
}

func (fpnrtt *FilenamePatternNoRangeTitleTxt) Format(
	filenameInfo FilenameInfo) (filename string, err error) {
	if filenameInfo == nil {
		return "", ErrNilFilenameInfo
	}
	info, ok := filenameInfo.(*FilenameInfoNoRangeTitleTxt)
	if !ok {
		return "", ErrFilenameInfoNotMatchPattern
	}
	f := fmt.Sprintf("%s.%s.txt",
		formatNoRange(info.Numbers[0], info.Numbers[1]), info.T)
	f = replaceInvalidFilenameCharacters(f)
	return filepath.Join(info.Dir, f), nil
}

func (finrtt *FilenameInfoNoRangeTitleTxt) Directory() string {
	if finrtt == nil {
		return ""
	}
	return finrtt.Dir
}

func (finrtt *FilenameInfoNoRangeTitleTxt) Number() (
	number int32, isSupported bool) {
	if finrtt == nil {
		return 0, true
	}
	return finrtt.Numbers[0], true
}

func (finrtt *FilenameInfoNoRangeTitleTxt) NumberRange() (
	numberRange [2]int32, isSupported bool) {
	if finrtt == nil {
		return [2]int32{0, 0}, true
	}
	return finrtt.Numbers, true
}

func (finrtt *FilenameInfoNoRangeTitleTxt) Title() (
	title string, isSupported bool) {
	if finrtt == nil {
		return "", true
	}
	return finrtt.T, true
}
