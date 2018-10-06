package taggedtxt

import (
	"github.com/donyori/lab513_paper_format_convert/fnp"
)

type TaggedTextInfo struct {
	FilenamePattern    fnp.FilenamePattern
	DoesStartWithTitle bool
	ChapterStartTag    string
	ChapterEndTag      string
}

const (
	DefaultChapterStartTag string = "；"
	DefaultChapterEndTag   string = "。"
)

var (
	DefaultFilenamePattern fnp.FilenamePattern = fnp.NewFilenamePatternNoRangeTxt()
	DefaultTaggedTextInfo  *TaggedTextInfo     = &TaggedTextInfo{
		FilenamePattern:    DefaultFilenamePattern,
		DoesStartWithTitle: true,
		ChapterStartTag:    DefaultChapterStartTag,
		ChapterEndTag:      DefaultChapterEndTag,
	}
)

func NewTaggedTextInfo(filenamePattern fnp.FilenamePattern,
	doesStartWithTitle bool, chapterStartTag string,
	chapterEndTag string) *TaggedTextInfo {
	if chapterStartTag == "" {
		chapterStartTag = DefaultChapterStartTag
	}
	if chapterEndTag == "" {
		chapterEndTag = DefaultChapterEndTag
	}
	return &TaggedTextInfo{
		FilenamePattern:    filenamePattern,
		DoesStartWithTitle: doesStartWithTitle,
		ChapterStartTag:    chapterStartTag,
		ChapterEndTag:      chapterEndTag,
	}
}
