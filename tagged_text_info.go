package lab513_paper_format_convert

type TaggedTextInfo struct {
	DoesStartWithTitle bool
	ChapterStartTag    string
	ChapterEndTag      string
}

const (
	DefaultChapterStartTag string = "；"
	DefaultChapterEndTag   string = "。"
)

var DefaultTaggedTextInfo *TaggedTextInfo = &TaggedTextInfo{
	DoesStartWithTitle: true,
	ChapterStartTag:    DefaultChapterStartTag,
	ChapterEndTag:      DefaultChapterEndTag,
}

func NewTaggedTextInfo(doesStartWithTitle bool, chapterStartTag string,
	chapterEndTag string) *TaggedTextInfo {
	if chapterStartTag == "" {
		chapterStartTag = DefaultChapterStartTag
	}
	if chapterEndTag == "" {
		chapterEndTag = DefaultChapterEndTag
	}
	return &TaggedTextInfo{
		DoesStartWithTitle: doesStartWithTitle,
		ChapterStartTag:    chapterStartTag,
		ChapterEndTag:      chapterEndTag,
	}
}
