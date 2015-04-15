package block

import (
	"fmt"
)

// A preformatted block
type Preformatted struct {
	BaseBlock
}

func (p *Preformatted) Decode(enc interface{}) error {
	return p.decodeBlock(enc)
}

func (p *Preformatted) AsHtml() string {
	return fmt.Sprintf("<pre>%s</pre>", p.FormatHtmlText())
}

func (p *Preformatted) AsMarkdown(cnt int) string {
	return fmt.Sprintf("<pre>%s</pre>", p.FormatMarkdownText())
}

func (p *Preformatted) ParentHtmlTag() string {
	return ""
}
