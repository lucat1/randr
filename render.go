package randr

// Component is a struct representing any component
// which receives the rendering context and returns
// a string with valid HTML code inside
type Component func(context Context) string

// RenderWithData renders a component into a string with the given props and initial context data
func RenderWithData(root Component, ctx Context) (string, Context) {
	return root(ctx), ctx
}

// Render renders a component into a string with the given props
func Render(root Component, props interface{}) (string, Context) {
	ctx := Context{
		Props: props,
		Data: map[ContextKey]interface{}{},
	}

	return RenderWithData(root, ctx)
}

// MustRender renders a component into a string with the given props
// NOTE: Only used internally to render a component into a string
// ignoring the context, in cases when it's useless, such as when
// rendering a custom element from a computed template literal.
//
// Should always use `Render` or `RenderWithData`
func MustRender(c Component, ctx Context, props interface{}) string {
	res, _ := RenderWithData(c, Inherit(ctx, Context{Props: props}))
	return res
}