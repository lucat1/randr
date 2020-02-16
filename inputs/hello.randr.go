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
		var MRAjWwhTHc string
		for _, name1 := range names {
			MRAjWwhTHc += name1
			MRAjWwhTHc += ` once again `
			MRAjWwhTHc += name
		}
		XVlBzgbaiC += randr.MustRender(Hello, &HelloProps{Name: name, Children: `This is a test ` + name + MRAjWwhTHc})
	}

	return `<div><h1>Hello list</h1>` + XVlBzgbaiC + `</div>`

}

func main() {
	fmt.Println(randr.MustRender(MultipleHellos, nil))
}
