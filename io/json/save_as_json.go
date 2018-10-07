package json

import (
	stdjson "encoding/json"
	"os"

	"github.com/donyori/lab513_paper_format_convert/model"
)

func SaveAsJson(docModel *model.DocumentModel, filename string,
	isPretty bool) error {
	jsonModel, err := BuildJsonDocumentModel(docModel)
	if err != nil {
		return err
	}
	var data []byte
	if isPretty {
		data, err = stdjson.MarshalIndent(jsonModel, "", "    ")
	} else {
		data, err = stdjson.Marshal(jsonModel)
	}
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}
