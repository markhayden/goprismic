package block

import (
	"fmt"
)

// A list block (unordered)
type ListItem struct {
	BaseBlock
}

func (l *ListItem) Decode(enc interface{}) error {
	return l.decodeBlock(enc)
}

func (l *ListItem) AsHtml() string {
	return fmt.Sprintf("<li>%s</li>", l.FormatHtmlText())
}

func (l *ListItem) AsMarkdown(cnt int) string {
	return fmt.Sprintf("%d. %s", cnt+1, l.FormatMarkdownText())
}

func (l *ListItem) ParentHtmlTag() string {
	return "ul"
}

// A list block (ordered)
type OrderedListItem struct {
	BaseBlock
}

func (l *OrderedListItem) Decode(enc interface{}) error {
	return l.decodeBlock(enc)
}

func (l *OrderedListItem) AsHtml() string {
	return fmt.Sprintf("<li>%s</li>", l.FormatHtmlText())
}

func (l *OrderedListItem) AsMarkdown(cnt int) string {
	return fmt.Sprintf("* %s", l.FormatMarkdownText())
}

func (l *OrderedListItem) ParentHtmlTag() string {
	return "ol"
}
