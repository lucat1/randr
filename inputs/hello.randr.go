package main

import (
	"fmt"
	"strconv"

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
		for i :=
			1; i < 5; i++ {
			var tcuAxhxKQF string
			if name ==
				"Anna" {
				tcuAxhxKQF += `<h1>what a special name</h1>`
			} else {
				tcuAxhxKQF += `<h1>boring name</h1>`
			}
			MRAjWwhTHc += name + ` -- #` + strconv.Itoa(i) + ` ` + tcuAxhxKQF
		}
		XVlBzgbaiC += randr.MustRender(Hello, ctx, &HelloProps{Name: name, Children: `This is a test ` + name + MRAjWwhTHc})
	}

	return `<div><h1>Hello list</h1>` + XVlBzgbaiC + `</div>`

}

func main() {
	res, _ := randr.Render(MultipleHellos, nil)
	fmt.Println(res)
}