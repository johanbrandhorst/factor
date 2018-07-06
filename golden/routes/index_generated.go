// This file was created with https://jsgo.io/dave/html2vecty
package main

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

func main() {
	vecty.RenderBody(&Page{})
}

type Page struct {
	vecty.Core
}

func (p *Page) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Main(
			elem.Heading1(
				vecty.Text("Great success!"),
			),
			elem.Figure(
				elem.Image(
					vecty.Markup(
						vecty.Attribute("alt", "Borat"),
						prop.Src("great-success.png"),
					),
				),
				elem.FigureCaption(
					vecty.Text("HIGH FIVE!"),
				),
			),
			elem.Paragraph(
				elem.Strong(
					vecty.Text("Try editing this file (routes/index.html) to test hot module reloading."),
				),
			),
		),
		elem.Style(
			vecty.Text("h1,\n\tfigure,\n\tp {\n\t\ttext-align: center;\n\t\tmargin: 0 auto;\n\t}\n\n\th1 {\n\t\tfont-size: 2.8em;\n\t\ttext-transform: uppercase;\n\t\tfont-weight: 700;\n\t\tmargin: 0 0 0.5em 0;\n\t}\n\n\tfigure {\n\t\tmargin: 0 0 1em 0;\n\t}\n\n\timg {\n\t\twidth: 100%;\n\t\tmax-width: 400px;\n\t\tmargin: 0 0 1em 0;\n\t}\n\n\tp {\n\t\tmargin: 1em auto;\n\t}\n\n\t@media (min-width: 480px) {\n\t\th1 {\n\t\t\tfont-size: 4em;\n\t\t}\n\t}"),
		),
	)
}
