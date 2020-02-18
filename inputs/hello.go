package main

// +build ignore

import (
	"fmt"
	"strconv"

	"github.com/lucat1/randr"
)

type HelloProps struct {
	Children string
	Name     string
}

// Hello renders a single h1 tag with an hello message
func Hello(ctx randr.Context) string {
	props := ctx.Props.(*HelloProps)

	return randr.HTML(`
		<h1 style="color: red">Hello {props.Name}!; children: {props.Children}</h1>
		<div>{ctx.Data["parentData"].(string)}</div>
  `)
}

// MultipleHellos renders various hellos
func MultipleHellos(ctx randr.Context) string {
	names := []string{
		"Anna",
		"Tess",
	}

	ctx.Data["parentData"] = "Parent data should be displayed"

	return randr.HTML(`
		<div>
			<h1>Hello list</h1>

			{#for _, name := range names}
				<Hello name="{name}">
					This is a test {name}

					{#for i := 1; i < 5; i++}
						{name} -- #{strconv.Itoa(i)} 
					{/for}
				</Hello>

				{#if name == "Anna"}
					<h1>what a special name</h1>
				{else}
					<h1>boring name</h1>
				{/if}
			{/for}
    </div>
  `)
}

func main() {
	res, _ := randr.Render(MultipleHellos, nil)
	fmt.Println(res)
}
