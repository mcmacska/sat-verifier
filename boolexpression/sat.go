package main

import (
"fmt"
)

func main() {
	fmt.Println(true && ((true) && !true || (true)))
}
