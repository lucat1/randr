package main

import "github.com/lucat1/randr"

// ExternalLibrary renders a component fron another library
func ExternalLibrary(_ randr.Context) string {
  return randr.HTML(`
		<test.Test>
			External library
		</test.Test>
  `)
}