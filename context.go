package randr

// BasicProps is a struct used in many components
// which only care about their HTML children represented
// as a rendered string of valid HTML code
type BasicProps struct {
	Children string
}

// ContextKey is the unique key which must be
// implemented by any context providing/consuming library
type ContextKey string

// Context is the rendering context in which
// the component is executed(and rendered)
type Context struct {
	// Props are the arguments, attributes given to
	// each component function as the first argument
	Props interface{}

	// Data is a map containing the rendering
	// context data, set by components
	// so that their children can access it
	Data map[ContextKey]interface{}
}

// Inherit inherits all the available data from the ancestor
// context but *resets the props*. It is mainly used when
// calling a custom component with some props, or when
// you wanna give different context data to the component
func Inherit(ancestor Context, successor Context) Context {
	res := Context{
		Props: nil,
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