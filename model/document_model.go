package model

import (
	"errors"
)

type DocumentModel struct {
	number int32
	root   *DocumentTreeNode
}

var ErrDocumentTreeNodeKindNotDocument error = errors.New(
	"DocumentTreeNode's kind is not Document")

func NewDocumentModel(number int32, root *DocumentTreeNode) (
	docModel *DocumentModel, err error) {
	if root == nil {
		return nil, ErrNilDocumentTreeNode
	}
	if root.Kind() != Document {
		return nil, ErrDocumentTreeNodeKindNotDocument
	}
	return &DocumentModel{number: number, root: root}, nil
}

func (dm *DocumentModel) Number() int32 {
	if dm == nil {
		return 0
	}
	return dm.number
}

func (dm *DocumentModel) Root() *DocumentTreeNode {
	if dm == nil {
		return nil
	}
	return dm.root
}
