package fnp

import (
	"os"
	"path/filepath"
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

func TestNoRangeTitleTxtFormat(t *testing.T) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		t.Fatal("Cannot get GOPATH.")
	}
	filename := filepath.Join(goPath, "src", "github.com",
		"donyori", "lab513_paper_format_convert", "test_resource",
		"134.SuPor An Environment for AS of Texts in Brazilian Portuguese.txt")
	fp := new(FilenamePatternNoRangeTitleTxt)
	info, err := fp.Parse(filename)
	if err != nil {
		t.Fatal("Error on parse:", err)
	}
	(info.(*FilenameInfoNoRangeTitleTxt)).Title = "SuPor: An Environment for AS of Texts in Brazilian Portuguese"
	result, err := fp.Format(info)
	if err != nil {
		t.Error(err)
	}
	t.Log("Format result:", result)
}
