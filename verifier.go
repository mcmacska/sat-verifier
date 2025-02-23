package main

import (
	"fmt"
	"os"
	"os/exec"
)

const TARGET_FOLDER = "boolexpression/"

const SAT_SOURCE_FILE = "sat.go"
const SAT_EXE_FILE = "sat.exe"

func readFile(filename string) string {
	var dat, err = os.ReadFile(filename)

	if err != nil {
		fmt.Println("Could not read file: " + filename)
	}

	return string(dat)
}

func writeFile(path string, content string) bool {
	contentByteArray := []byte(content)

	err := os.WriteFile(path, contentByteArray, 0644)

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func compileFile(source string, exeFile string) bool {
	// run 'go build' command
	cmd := exec.Command("go", "build", "-o", exeFile, source)

	// set the output to display any compilation errors
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run the command and check for errors
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to compile %s: %v\n", source, err)
		return false
	}

	// fmt.Printf("Successfully compiled %s into %s\n", source, exeFile)
	return true
}

func runFile(exeFile string) {
	cmd := exec.Command(exeFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running program: %v", err)
	}
	fmt.Printf("Output: %s\n", output)
}

func verify(expression string) {

	// create a sat.go inside directory
	if !writeFile(TARGET_FOLDER+SAT_SOURCE_FILE, "package main\n\nimport (\n\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println("+expression+")\n}\n") {
		return
	}

	//compile the file
	if !compileFile(TARGET_FOLDER+SAT_SOURCE_FILE, TARGET_FOLDER+SAT_EXE_FILE) {
		return
	}

	// run the file
	runFile(TARGET_FOLDER + SAT_EXE_FILE)
}
