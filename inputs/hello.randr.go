package main

import (
	"fmt"

	"github.com/lucat1/randr"
)

type HelloProps struct {
	Children	string
	Name		string
}

// Hello renders a single h1 tag with an hello message
func Hello(ctx randr.Context) string {
	props := ctx.Props.(*HelloProps)

	return `<h1 style="color: red">Hello ` + props.Name + `!; children: ` + props.Children + `</h1>`

}

// MultipleHellos renders various hellos
func MultipleHellos(ctx randr.Context) string {
	names := []string{
		"Anna",
		"Tess",
	}
	var XVlBzgbaiC string
	for _, name := range names {
		XVlBzgbaiC += randr.MustRender(Hello, &HelloProps{Name: `name`, Children: name})
	}

	return `<div><h1>Hello list</h1>` + XVlBzgbaiC + `</div>`

}

func main() {
	fmt.Println(randr.MustRender(MultipleHellos, nil))
}