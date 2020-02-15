package parser

import "errors"

// Validate checks if the ammount of exprs is appropriate
// for the ammount of raw strins
func Validate(raws []string, exprs []string) error {
	if len(raws)-1 != len(exprs) {
		return errors.New("Error, invalid parsing. The number of exprs doesn't match the number of raws-1")
	}

	// loop trough all experssions and count the depth to see
	// if some experssions aren't properly closed
	indent := 0
	for _, expr := range exprs {
		if expr[:2] == "{#" {
			indent++
		}

		if expr[:2] == "{/" {
			indent--
		}
	}

	if indent != 0 {
		return errors.New("Unclosed tag in template")
	}

	return nil
}