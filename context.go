package randr

// BasicProps is a struct which must be implemented
// by any props implementation, and contains the children
// string which can always be given to any custom component
type BasicProps struct {
	Children string
}

// Context is the rendering context in which
// the component is executed(and rendered)
type Context struct {
	// Props are the arguments, attributes given to
	// each component function as the first argument
	Props *BasicProps

	// Data is a map containing the rendering
	// context data, set by components
	// so that their children can access it
	Data *map[string]interface{}
}

// Inherit inherits all the available data from the ancestor
// context but *resets the props*. It is mainly used when
// calling a custom component with some props, or when
// you wanna give different context data to the component
func Inherit(ancestor Context, successor Context) Context {
	res := Context{
		Props: &BasicProps{},
		Data: ancestor.Data,
	}
	if successor.Props != nil {
		res.Props = successor.Props
	}

	if successor.Data != nil {
		res.Data = successor.Data
	}

	return res
}