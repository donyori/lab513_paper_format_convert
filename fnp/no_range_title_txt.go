package fnp

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type FilenamePatternNoRangeTitleTxt struct{}

type FilenameInfoNoRangeTitleTxt struct {
	Directory string
	Numbers   [2]int32
	Title     string
}

var fpnrttRegexp *regexp.Regexp

func init() {
	fpnrttRegexp = regexp.MustCompile(
		"^[[:digit:]]+(\\-[[:digit:]]+)??\\..+\\.txt$")
}

func NewFilenamePatternNoRangeTitleTxt() *FilenamePatternNoRangeTitleTxt {
	return new(FilenamePatternNoRangeTitleTxt)
}

func (fpnrtt *FilenamePatternNoRangeTitleTxt) Parse(filename string) (
	filenameInfo interface{}, err error) {
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
		Directory: dir,
		Numbers:   [2]int32{no1, no2},
		Title:     filename[firstDot+1 : len(filename)-4],
	}
	return info, nil
}

func (fpnrtt *FilenamePatternNoRangeTitleTxt) Format(filenameInfo interface{}) (
	filename string, err error) {
	if filenameInfo == nil {
		return "", ErrNilFilenameInfo
	}
	info, ok := filenameInfo.(*FilenameInfoNoRangeTitleTxt)
	if !ok {
		return "", ErrFilenameInfoNotMatchPattern
	}
	f := fmt.Sprintf("%s.%s.txt",
		formatNoRange(info.Numbers[0], info.Numbers[1]), info.Title)
	return filepath.Join(info.Directory, f), nil
}
