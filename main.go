package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/donyori/lab513_paper_format_convert/fnp"
	"github.com/donyori/lab513_paper_format_convert/io/json"
	"github.com/donyori/lab513_paper_format_convert/io/taggedtxt"
)

func main() {
	in := flag.String("in", "", "Directory of input tagged text files. Required.")
	out := flag.String("out", "", "Directory of output json files. Required.")
	inTitlePattern := flag.Int("itp", 0,
		"Title pattern of input tagged text files: 1-NumberRangeTxt, 2-NumberRangeTitleTxt, 3-NumberRangeTitlePdfTxt. Required.")
	inStartWithTitle := flag.Bool("st", true,
		"True if input tagged text files start with the title of paper.")
	flag.Parse()
	if *in == "" || *out == "" || *inTitlePattern < 1 || *inTitlePattern > 3 {
		fmt.Println("Please give required parameters.")
		flag.Usage()
		return
	}
	var filenamePattern fnp.FilenamePattern
	switch *inTitlePattern {
	case 1:
		filenamePattern = fnp.NewFilenamePatternNoRangeTxt()
	case 2:
		filenamePattern = fnp.NewFilenamePatternNoRangeTitleTxt()
	case 3:
		filenamePattern = fnp.NewFilenamePatternNoRangeTitlePdfTxt()
	default:
		fmt.Println(`Bad parameter "itp".`)
		flag.Usage()
		return
	}
	files, err := ioutil.ReadDir(*in)
	if err != nil {
		fmt.Println("Error occurs:", err)
		return
	}
	outFilenamePattern := fnp.NewFilenamePatternTitleJson()
	count := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		docModel, err := taggedtxt.LoadTaggedTextFile(
			filepath.Join(*in, file.Name()),
			taggedtxt.NewTaggedTextInfo(
				filenamePattern, *inStartWithTitle, "", ""),
			true)
		if err != nil {
			fmt.Println("Error occurs:", err)
			fmt.Println("Input filename:", file.Name())
			return
		}
		outFilenameInfo := &fnp.FilenameInfoTitleJson{
			Dir: *out,
			T:   docModel.Root().Content(),
		}
		outFilename, err := outFilenamePattern.Format(outFilenameInfo)
		if err != nil {
			fmt.Println("Error occurs:", err)
			fmt.Println("Input filename:", file.Name())
			return
		}
		err = json.SaveAsJson(docModel, outFilename, true)
		if err != nil {
			fmt.Println("Error occurs:", err)
			fmt.Println("Input filename:", file.Name())
			return
		}
		count++
	}
	fmt.Println("Finish. Converted", count, "files.")
}
