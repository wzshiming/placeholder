package placeholder

import (
	"bytes"
	"fmt"
	"math"
	"text/template"
)

var temp = template.Must(template.New("").Parse(`<svg xmlns="http://www.w3.org/2000/svg" width="{{.Width}}" height="{{.Height}}" viewBox="0 0 {{.Width}} {{.Height}}" preserveAspectRatio="none"><rect width="{{.Width}}" height="{{.Height}}" fill="#eee"/><text text-anchor="middle" x="{{.WidthHalf}}" y="{{.HeightHalf}}" style="fill:#aaa;font-weight:bold;font-size:{{.FontSize}}px;font-family:Arial,Helvetica,sans-serif;dominant-baseline:central">{{.Text}}</text></svg>`))

// Placeholder Generate SVG placeholders
// #path:"/placeholder"#
type Placeholder struct {
}

// Simple style
// #route:"GET /{width}x{height}"#
func (p *Placeholder) Simple(width, height uint64, text string) (file []byte /* #content:"image/svg+xml"# */, err error) {
	fontSize := math.Round(math.Max(12, math.Min(math.Min(float64(width), float64(height))*0.75, 0.75*math.Max(float64(width), float64(height))/12)))
	buf := bytes.NewBuffer(nil)
	if text == "" {
		text = fmt.Sprintf("%dx%d", width, height)
	}
	err = temp.Execute(buf, struct {
		Width      uint64
		WidthHalf  uint64
		Height     uint64
		HeightHalf uint64
		FontSize   uint64
		Text       string
	}{
		width,
		width / 2,
		height,
		height / 2,
		uint64(fontSize),
		text,
	})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
