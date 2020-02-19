package randr

// HTML is a dummy function to enable auto-completion
// when writing randr components, but it is never actually
// called during runtime since `randr` is an entirely pre-compiled library
func HTML(tmpl string) string {
	panic("Cannot call randr.HTML at runtime, please run your templates\ntrough the `rcc` tool before including them into your golang application")
}