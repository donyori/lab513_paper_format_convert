package fnp

import (
	"testing"
)

func TestNoRangeTitlePdfTxtParse(t *testing.T) {
	cases := []struct {
		Filename string
		Result   error
	}{
		{"", ErrEmptyFilename},
		{"03-05.txt", ErrFilenameNotMatchPattern},
		{"03.title.txt", ErrFilenameNotMatchPattern},
		{"03.title.pdf.txt", nil},
		{"03.a long-title test_here.pdf.txt", nil},
		{"03.title.jpg", ErrFilenameNotMatchPattern},
		{"03-05-06.txt", ErrFilenameNotMatchPattern},
	}
	fp := new(FilenamePatternNoRangeTitlePdfTxt)
	for i, c := range cases {
		info, err := fp.Parse(c.Filename)
		if err == c.Result {
			t.Log("case", i, "pass")
		} else {
			t.Error("case", i, "fail")
		}
		if info != nil {
			fi := info.(*FilenameInfoNoRangeTitlePdfTxt)
			t.Logf("%+v", *fi)
		}
	}
}
