package taggedtxt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/donyori/lab513_paper_format_convert/fnp"
	"github.com/donyori/lab513_paper_format_convert/model"
)

func LoadTaggedTextFile(filename string, taggedTextInfo *TaggedTextInfo,
	doesTrimSpace bool) (docModel *model.DocumentModel, err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			docModel = nil
			if e, ok := panicErr.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", panicErr)
			}
		}
	}()
	if taggedTextInfo == nil {
		taggedTextInfo = DefaultTaggedTextInfo
	}
	numberFromFilename, titleFromFilename, _ := parseFilename(
		filename, taggedTextInfo.FilenamePattern)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	startTagLength := len(taggedTextInfo.ChapterStartTag)
	endTagLength := len(taggedTextInfo.ChapterEndTag)
	scanner := bufio.NewScanner(file)
	var root, current *model.DocumentTreeNode
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
			title := titleFromFilename
			if taggedTextInfo.DoesStartWithTitle {
				title = line
			}
			root, err = model.NewDocumentTreeNode(model.Document, title)
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
					node, err := model.NewDocumentTreeNode(
						model.Chapter, content)
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
					current = current.Parent()
					if current == nil {
						return nil, model.ErrNilDocumentTreeNode
					}
				}
				node, err := model.NewDocumentTreeNode(
					model.Chapter, line[level*startTagLength:])
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
				current = current.Parent()
				if current == nil {
					return nil, model.ErrNilDocumentTreeNode
				}
			}
			depth = level
		} else {
			node, err := model.NewDocumentTreeNode(model.Paragraph, line)
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
	docModel, err = model.NewDocumentModel(numberFromFilename, root)
	if err != nil {
		return nil, err
	}
	return docModel, nil
}

func parseFilename(filename string, filenamePattern fnp.FilenamePattern) (
	number int32, title string, ok bool) {
	if filenamePattern == nil {
		return 0, "", false
	}
	info, err := filenamePattern.Parse(filename)
	if err != nil {
		return 0, "", false
	}
	n, isSupported := info.Number()
	if isSupported {
		number = n
	}
	t, isSupported := info.Title()
	if isSupported {
		title = t
	}
	return number, title, true
}
