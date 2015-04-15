package block

import (
	"github.com/markhayden/goprismic/fragment/image"
)

// An image block
type Image struct {
	BaseBlock
	View *image.View
}

func (i *Image) Decode(enc interface{}) error {
	i.View = new(image.View)
	err := i.View.Decode(enc)
	if err != nil {
		return err
	}
	return i.decodeBlock(enc)
}

func (i *Image) AsHtml() string {
	return i.View.AsHtml()
}

func (i *Image) AsMarkdown(cnt int) string {
	return i.View.AsMarkdown()
}

func (i *Image) ParentHtmlTag() string {
	return ""
}
