package fnp

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type FilenamePatternNoRangeTitlePdfTxt struct{}

type FilenameInfoNoRangeTitlePdfTxt FilenameInfoNoRangeTitleTxt

var fpnrtptRegexp *regexp.Regexp

func init() {
	fpnrtptRegexp = regexp.MustCompile(
		"^[[:digit:]]+(\\-[[:digit:]]+)??\\..+\\.pdf\\.txt$")
}

func NewFilenamePatternNoRangeTitlePdfTxt() *FilenamePatternNoRangeTitlePdfTxt {
	return new(FilenamePatternNoRangeTitlePdfTxt)
}

func (fpnrtpt *FilenamePatternNoRangeTitlePdfTxt) Parse(filename string) (
	filenameInfo interface{}, err error) {
	if fpnrtpt == nil {
		return nil, ErrNilFilenamePattern
	}
	dir, filename := filepath.Split(filename)
	if filename == "" {
		return nil, ErrEmptyFilename
	}
	if !fpnrtptRegexp.MatchString(filename) {
		return nil, ErrFilenameNotMatchPattern
	}
	firstDot := strings.IndexRune(filename, '.')
	noStr := filename[:firstDot]
	no1, no2, err := parseNoRange(noStr)
	if err != nil {
		return nil, err
	}
	info := &FilenameInfoNoRangeTitlePdfTxt{
		Directory: dir,
		Numbers:   [2]int32{no1, no2},
		Title:     filename[firstDot+1 : len(filename)-8],
	}
	return info, nil
}

func (fpnrtpt *FilenamePatternNoRangeTitlePdfTxt) Format(
	filenameInfo interface{}) (filename string, err error) {
	if filenameInfo == nil {
		return "", ErrNilFilenameInfo
	}
	info, ok := filenameInfo.(*FilenameInfoNoRangeTitlePdfTxt)
	if !ok {
		return "", ErrFilenameInfoNotMatchPattern
	}
	f := fmt.Sprintf("%s.%s.pdf.txt",
		formatNoRange(info.Numbers[0], info.Numbers[1]), info.Title)
	return filepath.Join(info.Directory, f), nil
}
