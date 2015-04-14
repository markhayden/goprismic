package span

import (
	"github.com/markhayden/goprismic/fragment/link"
)

type SpanInterface interface {
	GetStart() int
	GetEnd() int
	HtmlBeginTag() string
	HtmlEndTag() string
	Decode(interface{}) error
	ResolveLinks(link.Resolver)
}
