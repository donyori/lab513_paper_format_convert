package json

import (
	"errors"

	"github.com/donyori/lab513_paper_format_convert/model"
)

type JsonParagraphModel struct {
	Kind    string `json:"kind"`
	Content string `json:"content"`
}

type JsonChapterModel struct {
	Kind    string        `json:"kind"`
	Title   string        `json:"title"`
	Content []interface{} `json:"content"`
}

type JsonDocumentModel struct {
	Number  int32         `json:"number"`
	Title   string        `json:"title"`
	Content []interface{} `json:"content"`
}

func NewJsonParagraphModel(content string) *JsonParagraphModel {
	return &JsonParagraphModel{Kind: "Paragraph", Content: content}
}

func NewJsonChapterModel(title string) *JsonChapterModel {
	return &JsonChapterModel{
		Kind:    "Chapter",
		Title:   title,
		Content: make([]interface{}, 0, 8),
	}
}

func BuildJsonDocumentModel(docModel *model.DocumentModel) (
	jsonDocModel *JsonDocumentModel, err error) {
	if docModel == nil {
		return nil, model.ErrNilDocumentModel
	}
	currentNode := docModel.Root()
	jdm := &JsonDocumentModel{
		Number:  docModel.Number(),
		Title:   currentNode.Content(),
		Content: make([]interface{}, 0, 8),
	}
	currentNode = currentNode.FirstChild()
	modelStack := make([]interface{}, 0, 6)
	modelStack = append(modelStack, jdm)
	stackIdx := 0
	for currentNode != nil {
		// Create new model:
		var newModel interface{}
		switch currentNode.Kind() {
		case model.Chapter:
			newModel = NewJsonChapterModel(currentNode.Content())
		case model.Paragraph:
			newModel = NewJsonParagraphModel(currentNode.Content())
		case model.Document:
			return nil, errors.New("nested document is not supported")
		}
		// Add new model to its parent:
		topModel := modelStack[stackIdx]
		switch tm := topModel.(type) {
		case *JsonDocumentModel:
			tm.Content = append(tm.Content, newModel)
		case *JsonChapterModel:
			tm.Content = append(tm.Content, newModel)
		case *JsonParagraphModel:
			return nil, errors.New("nested paragraph is not supported")
		}
		// Update currentNode, modelStack and stackIdx:
		cn, deltaLevel, e := currentNode.Next()
		if e == model.ErrNoNextDocumentTreeNode {
			currentNode = nil
		} else if e != nil {
			return nil, e
		} else {
			currentNode = cn
			if deltaLevel < 0 {
				stackIdx += deltaLevel
			} else if deltaLevel == 1 {
				stackIdx++
				if stackIdx >= len(modelStack) {
					modelStack = append(modelStack, newModel)
				} else {
					modelStack[stackIdx] = newModel
				}
			}
		}
	}
	return jdm, nil
}
