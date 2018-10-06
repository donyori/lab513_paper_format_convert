package taggedtxt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/donyori/lab513_paper_format_convert/fnp"
)

func TestLoadTaggedTextFile(t *testing.T) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		t.Fatal("Cannot get GOPATH.")
	}
	dir := filepath.Join(goPath, "src", "github.com",
		"donyori", "lab513_paper_format_convert", "test_resource")
	cases := []struct {
		Filename string
		TTI      *TaggedTextInfo
	}{
		{
			"20-21.txt",
			NewTaggedTextInfo(
				fnp.NewFilenamePatternNoRangeTxt(), true, "", ""),
		},
		{
			"132.Centroid-based summarization of multiple documents.pdf.txt",
			NewTaggedTextInfo(
				fnp.NewFilenamePatternNoRangeTitlePdfTxt(), false, "", ""),
		},
		{
			"134.SuPor An Environment for AS of Texts in Brazilian Portuguese.txt",
			NewTaggedTextInfo(
				fnp.NewFilenamePatternNoRangeTitleTxt(), true, "", ""),
		},
	}
	for i, c := range cases {
		f := filepath.Join(dir, c.Filename)
		root, err := LoadTaggedTextFile(f, c.TTI, true)
		t.Log(i)
		if err != nil {
			t.Error(err)
		}
		if root != nil {
			children := root.Children()
			t.Log("Title:", root.Content())
			t.Log("1 level chapter:")
			for _, child := range children {
				t.Log("    ", child.Content())
			}
		}
	}
}
