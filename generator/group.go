package generator

// group groups bot the raws and the exprs
// arrays into a single array of ast nodes
func group(raws []string, exprs []string) []string {
	res := []string{}

	for i, raw := range raws {
		if i < len(exprs) {
			res = append(res, raw, exprs[i])
		} else {
			res = append(res, raw)
		}
	}

	return res
}
