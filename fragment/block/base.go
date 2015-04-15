package block

import (
	"fmt"
	"html"

	"github.com/markhayden/goprismic/fragment/link"
	"github.com/markhayden/goprismic/fragment/span"
)

// Common block properties
type BaseBlock struct {
	Type  string               `json:"type"`
	Text  string               `json:"text"`
	Spans []span.SpanInterface `json:"spans"`
}

func (b *BaseBlock) AsText() string {
	return b.Text
}

// Formats the block content as html, without enclosing tags
func (b *BaseBlock) FormatHtmlText() string {
	t := html.EscapeString(b.Text)
	// store one more to be able to compute offsets[len(text)]
	offsets := make([]int, len(b.Text)+1)
	// compute byte offsets for utf8 string
	off := 0
	index := 0
	for k, r := range b.Text {
		offsets[index] = k + off
		off += len([]rune(html.EscapeString(string([]rune{r})))) - 1
		index++
	}
	offsets[index] = len(b.Text) + off
	for _, s := range b.Spans {
		begin := s.HtmlBeginTag()
		end := s.HtmlEndTag()
		t = t[:offsets[s.GetStart()]] + begin + t[offsets[s.GetStart()]:offsets[s.GetEnd()]] + end + t[offsets[s.GetEnd()]:]
		for i := s.GetStart(); i < s.GetEnd(); i++ {
			offsets[i] += len(begin)
		}
		for i := s.GetEnd(); i < len(offsets); i++ {
			offsets[i] += len(begin) + len(end)
		}
	}
	return t
}

// Formats the block content as html, without enclosing tags
func (b *BaseBlock) FormatMarkdownText() string {
	t := html.EscapeString(b.Text)

	// store one more to be able to compute offsets[len(text)]
	offsets := make([]int, len(b.Text)+1)

	// compute byte offsets for utf8 string
	off := 0
	index := 0
	for k, r := range b.Text {
		offsets[index] = k + off
		off += len([]rune(html.EscapeString(string([]rune{r})))) - 1
		index++
	}

	offsets[index] = len(b.Text) + off
	for _, s := range b.Spans {
		begin := s.MarkdownBeginTag()
		end := s.MarkdownEndTag()

		t = t[:offsets[s.GetStart()]] + begin + t[offsets[s.GetStart()]:offsets[s.GetEnd()]] + end + t[offsets[s.GetEnd()]:]
		for i := s.GetStart(); i < s.GetEnd(); i++ {
			offsets[i] += len(begin)
		}

		for i := s.GetEnd(); i < len(offsets); i++ {
			offsets[i] += len(begin) + len(end)
		}
	}

	return t
}

func (b *BaseBlock) decodeBlock(enc interface{}) error {
	dec, ok := enc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("%+v is not a map", enc)
	}
	if v, found := dec["type"]; found {
		b.Type = v.(string)
	}
	if v, found := dec["text"]; found {
		b.Text = v.(string)
	}
	if v, found := dec["spans"]; found {
		dec2, ok := v.([]interface{})
		if !ok {
			return fmt.Errorf("%+v is not a slice", dec2)
		}
		b.Spans = make([]span.SpanInterface, 0, len(dec2))
		for _, v := range dec2 {
			dec3, ok := v.(map[string]interface{})
			if ok {
				var s span.SpanInterface
				switch dec3["type"] {
				case "strong":
					s = new(span.Strong)
				case "em":
					s = new(span.Em)
				case "hyperlink":
					s = new(span.Hyperlink)
				default:
					panic(fmt.Sprintf("Unknown span type %s", dec3["type"]))
				}
				err := s.Decode(v)
				if err == nil {
					b.Spans = append(b.Spans, s)
				}
			}
		}
	}
	return nil
}

// Resolves links
func (b *BaseBlock) ResolveLinks(r link.Resolver) {
	for _, v := range b.Spans {
		v.ResolveLinks(r)
	}
}
