package main

import "github.com/lucat1/randr"

// Static renders a static component
func Static(_ randr.Props) string {
	return `<h1>Only compiles with recursive</h1>`

}
