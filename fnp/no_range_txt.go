package fnp

import (
	"path/filepath"
	"regexp"
	"strings"
)

type FilenamePatternNoRangeTxt struct{}

type FilenameInfoNoRangeTxt struct {
	Dir     string
	Numbers [2]int32
}

var fpnrtRegexp *regexp.Regexp

func init() {
	fpnrtRegexp = regexp.MustCompile(`^[[:digit:]]+(\-[[:digit:]]+)??\.txt$`)
}

func NewFilenamePatternNoRangeTxt() *FilenamePatternNoRangeTxt {
	return new(FilenamePatternNoRangeTxt)
}

func (fpnrt *FilenamePatternNoRangeTxt) Parse(filename string) (
	filenameInfo FilenameInfo, err error) {
	if fpnrt == nil {
		return nil, ErrNilFilenamePattern
	}
	dir, filename := filepath.Split(filename)
	if filename == "" {
		return nil, ErrEmptyFilename
	}
	if !fpnrtRegexp.MatchString(filename) {
		return nil, ErrFilenameNotMatchPattern
	}
	noStr := filename[:strings.IndexRune(filename, '.')]
	no1, no2, err := parseNoRange(noStr)
	if err != nil {
		return nil, err
	}
	info := &FilenameInfoNoRangeTxt{
		Dir:     dir,
		Numbers: [2]int32{no1, no2},
	}
	return info, nil
}

func (fpnrt *FilenamePatternNoRangeTxt) Format(filenameInfo FilenameInfo) (
	filename string, err error) {
	if filenameInfo == nil {
		return "", ErrNilFilenameInfo
	}
	info, ok := filenameInfo.(*FilenameInfoNoRangeTxt)
	if !ok {
		return "", ErrFilenameInfoNotMatchPattern
	}
	f := formatNoRange(info.Numbers[0], info.Numbers[1]) + ".txt"
	f = replaceInvalidFilenameCharacters(f)
	return filepath.Join(info.Dir, f), nil
}

func (finrt *FilenameInfoNoRangeTxt) Directory() string {
	if finrt == nil {
		return ""
	}
	return finrt.Dir
}

func (finrt *FilenameInfoNoRangeTxt) Number() (number int32, isSupported bool) {
	if finrt == nil {
		return 0, true
	}
	return finrt.Numbers[0], true
}

func (finrt *FilenameInfoNoRangeTxt) NumberRange() (
	numberRange [2]int32, isSupported bool) {
	if finrt == nil {
		return [2]int32{0, 0}, true
	}
	return finrt.Numbers, true
}

func (finrt *FilenameInfoNoRangeTxt) Title() (title string, isSupported bool) {
	return "", false
}
