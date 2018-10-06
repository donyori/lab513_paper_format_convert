package fnp

import (
	"testing"
)

func TestNoRangeTitleTxtParse(t *testing.T) {
	cases := []struct {
		Filename string
		Result   error
	}{
		{"", ErrEmptyFilename},
		{"03-05.txt", ErrFilenameNotMatchPattern},
		{"03.title.txt", nil},
		{"03.a long-title test_here.txt", nil},
		{"03.title.jpg", ErrFilenameNotMatchPattern},
		{"03-05-06.txt", ErrFilenameNotMatchPattern},
	}
	fp := new(FilenamePatternNoRangeTitleTxt)
	for i, c := range cases {
		info, err := fp.Parse(c.Filename)
		if err == c.Result {
			t.Log("case", i, "pass")
		} else {
			t.Error("case", i, "fail")
		}
		if info != nil {
			fi := info.(*FilenameInfoNoRangeTitleTxt)
			t.Logf("%+v", *fi)
		}
	}
}
