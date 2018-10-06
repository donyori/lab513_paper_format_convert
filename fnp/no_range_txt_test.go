package fnp

import (
	"testing"
)

func TestNoRangeTxtParse(t *testing.T) {
	cases := []struct {
		Filename string
		Result   error
	}{
		{"", ErrEmptyFilename},
		{"03-05.txt", nil},
		{"03.jpg", ErrFilenameNotMatchPattern},
		{"03.title.txt", ErrFilenameNotMatchPattern},
		{"03-05-06.txt", ErrFilenameNotMatchPattern},
	}
	fp := new(FilenamePatternNoRangeTxt)
	for i, c := range cases {
		info, err := fp.Parse(c.Filename)
		if err == c.Result {
			t.Log("case", i, "pass")
		} else {
			t.Error("case", i, "fail")
		}
		if info != nil {
			fi := info.(*FilenameInfoNoRangeTxt)
			t.Logf("%+v", *fi)
		}
	}
}
