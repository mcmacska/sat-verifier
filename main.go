package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

var NOT_ALLOWED_CHARACTERS = map[string]bool{"a": true, "e": true, "f": true, "l": true, "r": true, "s": true, "t": true, "u": true}
var NOT_ALLOWED_STRING = "aeflrstu"

const CORRECT_USAGE = "Correct usage: ./satverifier.exe [expression] [solution]"
const VARIABLE_NAMES_NOT_ALLOWED = "Variable names cannot contain the following letters: a, e, f, l, r, s, t, u"

// Returns the boolean expression with the assigned values.
// Variable names cannot contain the following letters: a, e, f, l, r, s, t, u
// Runtime: O(N^2)
func assignVariables(boolExpression string, solution map[string]string) (string, error) {
	// iterate over the variables
	for k := range solution {

		// check if character is allowed
		if strings.Contains(NOT_ALLOWED_STRING, k) {
			return "", errors.New("empty name")
		}

		// replace the current variables name with its value
		boolExpression = strings.ReplaceAll(boolExpression, k, string(solution[k]))
	}

	return boolExpression, nil
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println(CORRECT_USAGE)
		return
	}

	// first argument (after the program itself) is the boolean expression as a string
	// "x1 && (x2 && !x3 || x4)"
	expression := os.Args[1]

	// second argument is the proof as a string
	// `{"x1": "true", "x2": "false", "x3": "true", "x4": "true"}`
	str := os.Args[2]
	solution := map[string]string{}
	json.Unmarshal([]byte(str), &solution)

	assignedExpression, err := assignVariables(expression, solution)
	if err != nil {
		fmt.Println(VARIABLE_NAMES_NOT_ALLOWED)
		return
	}

	verify(assignedExpression)
}
