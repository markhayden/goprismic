package block

import (
	"github.com/markhayden/goprismic/fragment/link"
)

type Block interface {
	Decode(interface{}) error
	AsHtml() string
	AsMarkdown(cnt int) string
	AsText() string
	ParentHtmlTag() string
	ResolveLinks(link.Resolver)
}
