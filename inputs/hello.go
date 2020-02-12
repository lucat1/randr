package main


import "github.com/lucat1/randr"

type Props struct {
  name string
}

// Hello renders a single h1 tag with an hello message
func Hello(props Props) string {
  name := props.name + "!"

  return randr.HTML(`
    <h1 style="color: red">Hello {name}{name}</h1>
  `)
}

// MultipleHellos renders various hellos
func MultipleHellos(props Props) string {
  names := []string{
    "Anna",
    "Tess",
  }
  
  return randr.HTML(`
		<div>
			<h1>Hello list</h1>

      {#for k := range names}
        <Hello name={k}></Hello>
			{/for}
    </div>
  `)
}