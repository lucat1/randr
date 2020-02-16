package parser

import (
	"errors"
	"fmt"
	"strconv"
)

// Validate checks if the ammount of exprs is appropriate
// for the ammount of raw strins
func Validate(raws []string, exprs []string) error {
	if len(raws)-1 != len(exprs) {
		fmt.Printf("%q %q\n", raws, exprs)
		r, e := strconv.Itoa(len(raws)), strconv.Itoa(len(exprs))
		return errors.New("The number of exprs doesn't match the number of raws-1 (" + r + ", " + e + ")")
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