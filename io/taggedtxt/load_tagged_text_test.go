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
		{
			"136.A neural attention model for abstractive sentence summarization.txt",
			NewTaggedTextInfo(
				fnp.NewFilenamePatternNoRangeTitleTxt(), true, "", ""),
		},
		{
			"174.Image collection summarization via dictionary learning for sparse representation.txt",
			NewTaggedTextInfo(
				fnp.NewFilenamePatternNoRangeTitleTxt(), true, "", ""),
		},
	}
	for i, c := range cases {
		f := filepath.Join(dir, c.Filename)
		docModel, err := LoadTaggedTextFile(f, c.TTI, true)
		t.Log(i)
		if err != nil {
			t.Error(err)
		}
		if docModel != nil {
			root := docModel.Root()
			children := root.Children()
			t.Log("Number:", docModel.Number())
			t.Log("Title:", root.Content())
			t.Log("1 level chapter:")
			for _, child := range children {
				t.Log("    ", child.Content())
			}
			children = children[0].Children()
			t.Log("2 level nodes:")
			for _, child := range children {
				t.Log("    ", child.Content())
			}
		}
	}
}
