// This file was created with https://jsgo.io/dave/html2vecty
package components

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
)

/*
type Markdown struct {
	vecty.Core
}
*/

func (p *Markdown) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Body(
			elem.Div(
				vecty.Markup(
					vecty.Style("float", "right"),
				),
				elem.TextArea(
					vecty.Markup(
						vecty.Style("font-family", "monospace"),
						vecty.Attribute("cols", "70"),
						vecty.Attribute("rows", "14"),
						vecty.Attribute("oninput", "texthandler"),
					),
					vecty.Text("vecty-field:Input"),
				),
			),
		),
	)
}
