package model

import (
	"errors"
)

type DocumentTreeNodeKind int8

type DocumentTreeNode struct {
	kind       DocumentTreeNodeKind
	parent     *DocumentTreeNode
	sibling    *DocumentTreeNode
	firstChild *DocumentTreeNode
	content    string
}

const (
	Document DocumentTreeNodeKind = iota
	Chapter
	Paragraph
)

const InvalidDocNodeKind DocumentTreeNodeKind = -1

var (
	ErrInvalidDocumentTreeNodeKind error = errors.New(
		"DocumentTreeNodeKind is invalid")
	ErrNilDocumentTreeNode       error = errors.New("DocumentTreeNode is nil")
	ErrDocumentTreeNodesNotMatch error = errors.New(
		"DocumentTreeNodes do NOT match")
	ErrDocumentTreeNodeNoParent error = errors.New(
		"DocumentTreeNode has no parent")
	ErrDocumentTreeNodeChildLevelHigherThanParent error = errors.New(
		"child's level is higher than parent's level")

	dtnkStrings = [...]string{
		"Document",
		"Chapter",
		"Paragraph",
	}
)

func (dtnk DocumentTreeNodeKind) IsValid() bool {
	return dtnk >= Document && dtnk <= Paragraph
}

func (dtnk DocumentTreeNodeKind) String() string {
	if !dtnk.IsValid() {
		return "Unknown"
	}
	return dtnkStrings[dtnk]
}

func NewDocumentTreeNode(kind DocumentTreeNodeKind, content string) (
	node *DocumentTreeNode, err error) {
	if !kind.IsValid() {
		return nil, ErrInvalidDocumentTreeNodeKind
	}
	return &DocumentTreeNode{kind: kind, content: content}, nil
}

func (dtn *DocumentTreeNode) Kind() DocumentTreeNodeKind {
	if dtn == nil {
		return InvalidDocNodeKind
	}
	return dtn.kind
}

func (dtn *DocumentTreeNode) Parent() *DocumentTreeNode {
	if dtn == nil {
		return nil
	}
	return dtn.parent
}

func (dtn *DocumentTreeNode) Sibling() *DocumentTreeNode {
	if dtn == nil {
		return nil
	}
	return dtn.sibling
}

func (dtn *DocumentTreeNode) FirstChild() *DocumentTreeNode {
	if dtn == nil {
		return nil
	}
	return dtn.firstChild
}

func (dtn *DocumentTreeNode) LastChild() *DocumentTreeNode {
	if dtn == nil {
		return nil
	}
	n := dtn.firstChild
	for n.sibling != nil {
		n = n.sibling
	}
	return n
}

func (dtn *DocumentTreeNode) Children() []*DocumentTreeNode {
	if dtn == nil {
		return nil
	}
	n := dtn.firstChild
	ns := make([]*DocumentTreeNode, 0, 8)
	for n != nil {
		ns = append(ns, n)
		n = n.sibling
	}
	return ns
}

func (dtn *DocumentTreeNode) Content() string {
	if dtn == nil {
		return ""
	}
	return dtn.content
}

func (dtn *DocumentTreeNode) SetContent(content string) error {
	if dtn == nil {
		return ErrNilDocumentTreeNode
	}
	dtn.content = content
	return nil
}

func (dtn *DocumentTreeNode) AppendChild(node *DocumentTreeNode) error {
	return dtn.InsertChild(node, nil)
}

func (dtn *DocumentTreeNode) InsertChild(node *DocumentTreeNode,
	sibling *DocumentTreeNode) error {
	if dtn == nil || node == nil {
		return ErrNilDocumentTreeNode
	}
	if !dtn.checkChildSibling(sibling) {
		return ErrDocumentTreeNodesNotMatch
	}
	if node.kind < dtn.kind {
		return ErrDocumentTreeNodeChildLevelHigherThanParent
	}
	if err := node.Remove(); err != nil {
		return err
	}
	node.parent = dtn
	node.sibling = sibling
	if dtn.firstChild == nil {
		dtn.firstChild = node
		return nil
	}
	n := dtn.firstChild
	for n.sibling != sibling {
		n = n.sibling
	}
	n.sibling = node
	return nil
}

func (dtn *DocumentTreeNode) AppendSibling(node *DocumentTreeNode) error {
	if dtn == nil || node == nil {
		return ErrNilDocumentTreeNode
	}
	if dtn.parent == nil {
		return ErrDocumentTreeNodeNoParent
	}
	if node.kind < dtn.parent.kind {
		return ErrDocumentTreeNodeChildLevelHigherThanParent
	}
	if err := node.Remove(); err != nil {
		return err
	}
	node.parent = dtn.parent
	node.sibling = nil
	n := dtn
	for n.sibling != nil {
		n = n.sibling
	}
	n.sibling = node
	return nil
}

func (dtn *DocumentTreeNode) Remove() error {
	if dtn == nil {
		return ErrNilDocumentTreeNode
	}
	if dtn.parent != nil {
		if dtn.parent.firstChild == dtn {
			dtn.parent.firstChild = dtn.sibling
		} else {
			n := dtn.parent.firstChild
			for n.sibling != dtn && n.sibling != nil {
				n = n.sibling
			}
			if n.sibling != nil {
				n.sibling = dtn.sibling
			}
		}
	}
	dtn.parent = nil
	dtn.sibling = nil
	return nil
}

func (dtn *DocumentTreeNode) checkChildSibling(sibling *DocumentTreeNode) bool {
	if dtn == nil {
		return false
	}
	if sibling == nil {
		return true
	}
	if sibling.parent != dtn {
		return false
	}
	for n := dtn.firstChild; n != nil; n = n.sibling {
		if n == sibling {
			return true
		}
	}
	return false
}
