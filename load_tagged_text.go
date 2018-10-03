package lab513_paper_format_convert

import (
	"bufio"
	"os"
	"strings"
)

func LoadTaggedTextFile(filename string, taggedTextInfo *TaggedTextInfo,
	doesTrimSpace bool) (documentRoot *DocumentTreeNode, err error) {
	if taggedTextInfo == nil {
		taggedTextInfo = DefaultTaggedTextInfo
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	startTagLength := len(taggedTextInfo.ChapterStartTag)
	endTagLength := len(taggedTextInfo.ChapterEndTag)
	scanner := bufio.NewScanner(file)
	var root, current *DocumentTreeNode
	depth := 0
	for scanner.Scan() {
		// Process lines:
		line := scanner.Text()
		if doesTrimSpace {
			line = strings.TrimSpace(line)
		}
		if line == "" {
			// Ignore empty line.
			continue
		}
		if root == nil {
			title := ""
			if taggedTextInfo.DoesStartWithTitle {
				title = line
			}
			root, err = NewDocumentTreeNode(Document, title)
			if err != nil {
				return nil, err
			}
			current = root
			if taggedTextInfo.DoesStartWithTitle {
				continue
			}
		}
		if strings.HasPrefix(line, taggedTextInfo.ChapterStartTag) {
			level := 0
			for i := startTagLength; i < len(line); i += startTagLength {
				if line[i-startTagLength:i] != taggedTextInfo.ChapterStartTag {
					break
				}
				level++
			}
			if level > depth {
				for i := level - depth; i > 0; i-- {
					content := ""
					if i == 1 {
						content = line[level*startTagLength:]
					}
					node, err := NewDocumentTreeNode(Chapter, content)
					if err != nil {
						return nil, err
					}
					err = current.AppendChild(node)
					if err != nil {
						return nil, err
					}
					current = node
				}
			} else {
				for i := depth - level; i > 0; i-- {
					current = current.Parent
					if current == nil {
						return nil, ErrNilDocumentTreeNode
					}
				}
				node, err := NewDocumentTreeNode(
					Chapter, line[level*startTagLength:])
				if err != nil {
					return nil, err
				}
				err = current.AppendSibling(node)
				if err != nil {
					return nil, err
				}
				current = node
			}
			depth = level
		} else if strings.HasPrefix(line, taggedTextInfo.ChapterEndTag) {
			level := 0
			for i := endTagLength; i < len(line); i += endTagLength {
				if line[i-endTagLength:i] != taggedTextInfo.ChapterEndTag {
					break
				}
				level++
			}
			for i := depth - level; i >= 0; i-- {
				current = current.Parent
				if current == nil {
					return nil, ErrNilDocumentTreeNode
				}
			}
			depth = level
		} else {
			node, err := NewDocumentTreeNode(Paragraph, line)
			if err != nil {
				return nil, err
			}
			err = current.AppendChild(node)
			if err != nil {
				return nil, err
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return root, nil
}
