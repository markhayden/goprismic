package block

import (
	"github.com/markhayden/goprismic/fragment/embed"
)

// A embed block
type Embed struct {
	BaseBlock
	Embed *embed.Embed
}

func (e *Embed) Decode(enc interface{}) error {
	e.Embed = new(embed.Embed)
	err := e.Embed.Decode(enc)
	if err != nil {
		return err
	}
	return e.decodeBlock(enc)
}

func (e *Embed) AsHtml() string {
	return e.Embed.AsHtml()
}

func (e *Embed) AsMarkdown(cnt int) string {
	return e.Embed.AsMarkdown()
}

func (e *Embed) ParentHtmlTag() string {
	return ""
}
