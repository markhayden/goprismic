package fragment

import (
	"github.com/markhayden/goprismic/fragment/link"
)

type Interface interface {
	Decode(string, interface{}) error
	AsText() string
	AsHtml() string
	AsMarkdown() string
	ResolveLinks(link.Resolver)
}
