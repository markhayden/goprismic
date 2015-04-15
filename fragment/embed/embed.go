package embed

import (
	"fmt"
	"reflect"
	"strings"
)

// An embed (see http://oembed.com/)
type Embed struct {
	Type         string
	ProviderName string
	EmbedUrl     string
	Width        int
	Height       int
	Html         string
}

func (e *Embed) Decode(enc interface{}) error {
	dec, ok := enc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unable to decode embed fragment : %+v is a %s, not a map", enc, reflect.TypeOf(enc))
	}
	if v, found := dec["oembed"]; found {
		doc, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unable to decode embed fragment : %+v is a %s, not a map", enc, reflect.TypeOf(enc))
		}
		if v, found := doc["type"]; found && v != nil {
			e.Type = v.(string)
		}
		if v, found := doc["provider_name"]; found && v != nil {
			e.ProviderName = v.(string)
		}
		if v, found := doc["embed_url"]; found && v != nil {
			e.EmbedUrl = v.(string)
		}
		if v, found := doc["width"]; found && v != nil {
			e.Width = int(v.(float64))
		}
		if v, found := doc["height"]; found && v != nil {
			e.Height = int(v.(float64))
		}
		if v, found := doc["html"]; found && v != nil {
			e.Html = v.(string)
		}
	}
	return nil
}

func (e *Embed) AsHtml() string {
	if e.Html != "" {
		return "<div data-oembed=\"" + e.EmbedUrl + "\" data-oembed-type=\"" + strings.ToLower(e.Type) + "\" data-oembed-provider=\"" + e.ProviderName + "\">" + e.Html + "</div>"
	}
	return ""
}

func (e *Embed) AsMarkdown() string {
	if e.Html != "" {
		return "<div data-oembed=\"" + e.EmbedUrl + "\" data-oembed-type=\"" + strings.ToLower(e.Type) + "\" data-oembed-provider=\"" + e.ProviderName + "\">" + e.Html + "</div>"
	}
	return ""
}
