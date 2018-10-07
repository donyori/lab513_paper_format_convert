package fnp

import (
	"path/filepath"
	"regexp"
)

type FilenamePatternTitleJson struct{}

type FilenameInfoTitleJson struct {
	Dir string
	T   string
}

var fptjRegexp *regexp.Regexp

func init() {
	fptjRegexp = regexp.MustCompile(`\.json$`)
}

func NewFilenamePatternTitleJson() *FilenamePatternTitleJson {
	return new(FilenamePatternTitleJson)
}

func (fptj *FilenamePatternTitleJson) Parse(filename string) (
	filenameInfo FilenameInfo, err error) {
	if fptj == nil {
		return nil, ErrNilFilenamePattern
	}
	dir, filename := filepath.Split(filename)
	if filename == "" {
		return nil, ErrEmptyFilename
	}
	if !fptjRegexp.MatchString(filename) {
		return nil, ErrFilenameNotMatchPattern
	}
	info := &FilenameInfoTitleJson{
		Dir: dir,
		T:   filename[:len(filename)-5],
	}
	return info, nil
}

func (fptj *FilenamePatternTitleJson) Format(filenameInfo FilenameInfo) (
	filename string, err error) {
	if filenameInfo == nil {
		return "", ErrNilFilenameInfo
	}
	info, ok := filenameInfo.(*FilenameInfoTitleJson)
	if !ok {
		return "", ErrFilenameInfoNotMatchPattern
	}
	f := replaceInvalidFilenameCharacters(info.T + ".json")
	return filepath.Join(info.Dir, f), nil
}

func (fitj *FilenameInfoTitleJson) Directory() string {
	if fitj == nil {
		return ""
	}
	return fitj.Dir
}

func (fitj *FilenameInfoTitleJson) Number() (number int32, isSupported bool) {
	return 0, false
}

func (fitj *FilenameInfoTitleJson) NumberRange() (
	numberRange [2]int32, isSupported bool) {
	return [2]int32{0, 0}, false
}

func (fitj *FilenameInfoTitleJson) Title() (title string, isSupported bool) {
	if fitj == nil {
		return "", true
	}
	return fitj.T, true
}
