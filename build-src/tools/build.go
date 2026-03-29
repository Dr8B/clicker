package tools

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const ENV_BINARY_NAME_KEY = "BINARY_NAME"

func Build() {
	fmt.Printf("Запуск сборки...\n")

	binaryName := os.Getenv(ENV_BINARY_NAME_KEY)
	if binaryName == "" {
		log.Fatalf("%s is not set", ENV_BINARY_NAME_KEY)
	}

	cmd := exec.Command("go", "build", "-o", "./build/"+binaryName+".exe", "./src")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Ошибка запуска сборки: %v", err)
	}
	fmt.Printf("Сборка успешно завершена\n")
}
