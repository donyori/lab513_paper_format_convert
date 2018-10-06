package fnp

import (
	"path/filepath"
	"regexp"
	"strings"
)

type FilenamePatternNoRangeTxt struct{}

type FilenameInfoNoRangeTxt struct {
	Directory string
	Numbers   [2]int32
}

var fpnrtRegexp *regexp.Regexp

func init() {
	fpnrtRegexp = regexp.MustCompile("^[[:digit:]]+(\\-[[:digit:]]+)??\\.txt$")
}

func NewFilenamePatternNoRangeTxt() *FilenamePatternNoRangeTxt {
	return new(FilenamePatternNoRangeTxt)
}

func (fpnrt *FilenamePatternNoRangeTxt) Parse(filename string) (
	filenameInfo interface{}, err error) {
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
		Directory: dir,
		Numbers:   [2]int32{no1, no2},
	}
	return info, nil
}

func (fpnrt *FilenamePatternNoRangeTxt) Format(filenameInfo interface{}) (
	filename string, err error) {
	if filenameInfo == nil {
		return "", ErrNilFilenameInfo
	}
	info, ok := filenameInfo.(*FilenameInfoNoRangeTxt)
	if !ok {
		return "", ErrFilenameInfoNotMatchPattern
	}
	f := formatNoRange(info.Numbers[0], info.Numbers[1]) + ".txt"
	return filepath.Join(info.Directory, f), nil
}
