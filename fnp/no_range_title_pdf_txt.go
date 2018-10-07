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
		`^[[:digit:]]+(\-[[:digit:]]+)??\..+\.pdf\.txt$`)
}

func NewFilenamePatternNoRangeTitlePdfTxt() *FilenamePatternNoRangeTitlePdfTxt {
	return new(FilenamePatternNoRangeTitlePdfTxt)
}

func (fpnrtpt *FilenamePatternNoRangeTitlePdfTxt) Parse(filename string) (
	filenameInfo FilenameInfo, err error) {
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
		Dir:     dir,
		Numbers: [2]int32{no1, no2},
		T:       filename[firstDot+1 : len(filename)-8],
	}
	return info, nil
}

func (fpnrtpt *FilenamePatternNoRangeTitlePdfTxt) Format(
	filenameInfo FilenameInfo) (filename string, err error) {
	if filenameInfo == nil {
		return "", ErrNilFilenameInfo
	}
	info, ok := filenameInfo.(*FilenameInfoNoRangeTitlePdfTxt)
	if !ok {
		return "", ErrFilenameInfoNotMatchPattern
	}
	f := fmt.Sprintf("%s.%s.pdf.txt",
		formatNoRange(info.Numbers[0], info.Numbers[1]), info.T)
	f = replaceInvalidFilenameCharacters(f)
	return filepath.Join(info.Dir, f), nil
}

func (finrtpt *FilenameInfoNoRangeTitlePdfTxt) Directory() string {
	if finrtpt == nil {
		return ""
	}
	return finrtpt.Dir
}

func (finrtpt *FilenameInfoNoRangeTitlePdfTxt) Number() (
	number int32, isSupported bool) {
	if finrtpt == nil {
		return 0, true
	}
	return finrtpt.Numbers[0], true
}

func (finrtpt *FilenameInfoNoRangeTitlePdfTxt) NumberRange() (
	numberRange [2]int32, isSupported bool) {
	if finrtpt == nil {
		return [2]int32{0, 0}, true
	}
	return finrtpt.Numbers, true
}

func (finrtpt *FilenameInfoNoRangeTitlePdfTxt) Title() (
	title string, isSupported bool) {
	if finrtpt == nil {
		return "", true
	}
	return finrtpt.T, true
}
