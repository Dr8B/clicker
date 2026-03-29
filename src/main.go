package main

import (
	"log"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/goccy/go-yaml"
	"golang.org/x/sys/windows"
)

var (
	user32               = windows.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

const (
	VK_LCONTROL = 0xA2 // Код левого Ctrl
	VK_RCONTROL = 0xA3 // Код правого Ctrl
)

const (
	CONFIG_FILE = "config.yaml"
)

type Config struct {
	Profile  string `yaml:"profile"`
	Profiles map[string]struct {
		Clicks []string `yaml:"clicks"`
		Keys   []string `yaml:"keys"`
	} `yaml:"profiles"`
}

func isCtrlPressed() bool {
	// Вызываем GetAsyncKeyState для левого Ctrl
	// ret1, _, _ := procGetAsyncKeyState.Call(uintptr(VK_LCONTROL))
	// left := int16(ret1) < 0

	// Вызываем GetAsyncKeyState для правого Ctrl
	ret2, _, _ := procGetAsyncKeyState.Call(uintptr(VK_RCONTROL))
	right := int16(ret2) < 0

	return right
}

func main() {
	data, err := os.ReadFile("./" + CONFIG_FILE)
	if err != nil {
		log.Fatalf("Ошибка чтения файла %s: %v", CONFIG_FILE, err)
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Ошибка парсинга файла %s: %v", CONFIG_FILE, err)
	}

	if config.Profile == "" {
		log.Fatalf("Не указан profile в файле %s", CONFIG_FILE)
	}

	profile, exists := config.Profiles[config.Profile]
	if !exists {
		log.Fatalf("Профиль %s не существует в %s", config.Profile, CONFIG_FILE)
	}
	if len(profile.Clicks) == 0 && len(profile.Keys) == 0 {
		log.Fatalf("Нечем клиакать, укажите в профиле %s clicks или/или keys в файле %s", config.Profile, CONFIG_FILE)
	}

	log.Println("Автокликер запущен. Зажмите правый Ctrl для активации. Ctrl+C — выход.")

	for {
		if isCtrlPressed() {
			log.Println("Активирован автокликер (Ctrl зажат)")

			for isCtrlPressed() {
				for _, click := range profile.Clicks {
					robotgo.Click(click)
				}
				for _, key := range profile.Keys {
					// robotgo.Space
					robotgo.KeyTap(key)
				}
				// time.Sleep(10 * time.Millisecond)
			}
			log.Println("Автокликер остановлен (Ctrl отпущен)")
		}

		// Пауза между проверками — снижает нагрузку на CPU
		time.Sleep(100 * time.Millisecond)
	}
}
