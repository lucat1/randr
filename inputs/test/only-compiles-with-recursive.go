package main

// +build ignore

import "github.com/lucat1/randr"

// Static renders a static component
func Static(_ randr.Props) string {
	return randr.HTML(`
    <h1>Only compiles with recursive</h1>
  `)
}
