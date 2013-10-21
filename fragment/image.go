package fragment

import(
	"fmt"
)

type ImageView struct {
	Url        string `json:"url"`
	Alt        string `json:"alt"`
	Copyright  string `json:"copyright"`
	Dimensions struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}
}

func (i *ImageView) Decode(enc interface{}) error {
	dec, ok := enc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("%+v is not a map", enc)
	}
	if v, found := dec["url"]; found {
		i.Url = v.(string)
	}
	if v, found := dec["alt"]; found {
		i.Alt = v.(string)
	}
	if v, found := dec["copyright"]; found {
		i.Copyright = v.(string)
	}
	if d, found := dec["dimensions"]; found {
		dim, ok := d.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%+v is not a map", d)
		}
		if v, found := dim["width"]; found {
			i.Dimensions.Width = int(v.(float64))
		}
		if v, found := dim["height"]; found {
			i.Dimensions.Height = int(v.(float64))
		}
	}
	return nil
}

func (i *ImageView) AsHtml() string {
	return fmt.Sprintf("<img src=\"%s\" width=\"%d\" height=\"%d\"/>", i.Url, i.Dimensions.Width, i.Dimensions.Height);
}

type Image struct {
	Main  ImageView            `json:"main"`
	Views map[string]ImageView `json:"views"`
}

func (i *Image) GetView(view string) (*ImageView, bool) {
	v, found := i.Views[view]
	return &v, found
}

func (i *Image) Decode(enc interface{}) error {
	dec, ok := enc.(map[string]interface{})
	if !ok {
		return fmt.Errorf("%+v is not a map", enc)
	}
	if v, found := dec["main"]; found {
		(&i.Main).Decode(v)
	}
	if v, found := dec["views"]; found {
		views, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%+v is not a map", v)
		}
		i.Views = make(map[string]ImageView)
		for k, view := range views {
			iv := &ImageView{}
			iv.Decode(view)
			i.Views[k] = *iv
		}
	}
	return nil
}