package json

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/donyori/lab513_paper_format_convert/fnp"
	"github.com/donyori/lab513_paper_format_convert/io/taggedtxt"
)

func TestSaveAsJson(t *testing.T) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		t.Fatal("Cannot get GOPATH.")
	}
	dir := filepath.Join(goPath, "src", "github.com",
		"donyori", "lab513_paper_format_convert", "test_resource")
	input := filepath.Join(dir,
		"134.SuPor An Environment for AS of Texts in Brazilian Portuguese.txt")
	output_pretty := filepath.Join(dir, "save_as_json_pretty_test.json")
	output := filepath.Join(dir, "save_as_json_test.json")
	docModel, err := taggedtxt.LoadTaggedTextFile(
		input,
		taggedtxt.NewTaggedTextInfo(
			fnp.NewFilenamePatternNoRangeTitleTxt(),
			true, "", ""),
		true)
	if err != nil {
		t.Fatal(err)
	}
	err = SaveAsJson(docModel, output_pretty, true)
	if err != nil {
		t.Error(err)
	}
	err = SaveAsJson(docModel, output, false)
	if err != nil {
		t.Error(err)
	}
}
