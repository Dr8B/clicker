package main

import (
	"log"
	"os"

	"github.com/Dr8B/clicker/build-src/tools"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Укажите команду запуска")
	}
	switch os.Args[1] {
	case "clean":
		tools.Clean()
	case "build":
		tools.Build()
	}
}
