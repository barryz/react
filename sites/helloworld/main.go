// Template generated by reactGen

package main

import (
	r "github.com/myitcv/gopherjs/react"

	"honnef.co/go/js/dom"
)

//go:generate reactGen

var document = dom.GetWindow().Document()

func main() {
	domTarget := document.GetElementByID("app")

	r.Render(App(), domTarget)
}