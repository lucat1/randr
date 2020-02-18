package main

import "github.com/lucat1/randr"

// Static renders a static component
func Static(_ randr.Context) string {
  return randr.HTML(`
    <h1>Hello world!</h1>
  `)
}