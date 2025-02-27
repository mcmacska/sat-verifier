package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

var variableRegex = regexp.MustCompile(`\b[a-zA-Z0-9_]+\b`)

const CORRECT_USAGE = "Correct usage: ./satverifier.exe [expression] [solution]"

// Returns the boolean expression with the assigned values.
// Runtime: O(N^2)
func assignVariables(boolExpression string, solution map[string]string) (string, error) {

	modifiedExpression := variableRegex.ReplaceAllStringFunc(boolExpression, func(match string) string {
		// replaces the variable name with its value, if it exists in the solution
		value, exists := solution[match]
		if exists {
			return value
		}

		// leaves unchanged, if variable doesn't exist in the solution
		return match
	})

	return modifiedExpression, nil
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
		fmt.Println(err)
		return
	}

	verify(assignedExpression)
}
