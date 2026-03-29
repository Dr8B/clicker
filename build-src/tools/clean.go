package tools

import (
	"fmt"
	"os"
)

func Clean() {
	fmt.Printf("Remove build directory... ")
	os.RemoveAll("./build")
	fmt.Printf("Done.\n")
}
