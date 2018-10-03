package lab513_paper_format_convert

import (
	"errors"
)

type DocumentTreeNodeKind int8

type DocumentTreeNode struct {
	Kind       DocumentTreeNodeKind
	Parent     *DocumentTreeNode
	Sibling    *DocumentTreeNode
	FirstChild *DocumentTreeNode
	Content    string
}

const (
	Document DocumentTreeNodeKind = iota
	Chapter
	Paragraph
)

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
	if kind.IsValid() {
		return nil, ErrInvalidDocumentTreeNodeKind
	}
	return &DocumentTreeNode{Kind: kind, Content: content}, nil
}

func (dtn *DocumentTreeNode) LastChild() *DocumentTreeNode {
	if dtn == nil {
		return nil
	}
	n := dtn.FirstChild
	for n.Sibling != nil {
		n = n.Sibling
	}
	return n
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
	if node.Kind < dtn.Kind {
		return ErrDocumentTreeNodeChildLevelHigherThanParent
	}
	if err := node.Remove(); err != nil {
		return err
	}
	node.Parent = dtn
	node.Sibling = sibling
	if dtn.FirstChild == nil {
		dtn.FirstChild = node
		return nil
	}
	n := dtn.FirstChild
	for n.Sibling != sibling {
		n = n.Sibling
	}
	n.Sibling = node
	return nil
}

func (dtn *DocumentTreeNode) AppendSibling(node *DocumentTreeNode) error {
	if dtn == nil || node == nil {
		return ErrNilDocumentTreeNode
	}
	if dtn.Parent == nil {
		return ErrDocumentTreeNodeNoParent
	}
	if node.Kind < dtn.Parent.Kind {
		return ErrDocumentTreeNodeChildLevelHigherThanParent
	}
	if err := node.Remove(); err != nil {
		return err
	}
	node.Parent = dtn.Parent
	node.Sibling = nil
	n := dtn
	for n.Sibling != nil {
		n = n.Sibling
	}
	n.Sibling = node
	return nil
}

func (dtn *DocumentTreeNode) Remove() error {
	if dtn == nil {
		return ErrNilDocumentTreeNode
	}
	if dtn.Parent != nil {
		if dtn.Parent.FirstChild == dtn {
			dtn.Parent.FirstChild = dtn.Sibling
		} else {
			n := dtn.Parent.FirstChild
			for n.Sibling != dtn && n.Sibling != nil {
				n = n.Sibling
			}
			if n.Sibling != nil {
				n.Sibling = dtn.Sibling
			}
		}
	}
	dtn.Parent = nil
	dtn.Sibling = nil
	return nil
}

func (dtn *DocumentTreeNode) checkChildSibling(sibling *DocumentTreeNode) bool {
	if dtn == nil {
		return false
	}
	if sibling == nil {
		return true
	}
	if sibling.Parent != dtn {
		return false
	}
	for n := dtn.FirstChild; n != nil; n = n.Sibling {
		if n == sibling {
			return true
		}
	}
	return false
}
