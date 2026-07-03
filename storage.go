package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// AppConfig хранит настройки приложения для сохранения на диск
type AppConfig struct {
	IsEnglish bool `json:"is_english"`
}

// GetStorageDir возвращает путь к папке ~/nvrtier_pictures
func GetStorageDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("не удалось найти домашнюю директорию: %w", err)
	}
	return filepath.Join(home, "nvrtier_pictures"), nil
}

// GetImagesList собирает все картинки из папки
func GetImagesList(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var images []string
	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
				images = append(images, filepath.Join(dir, file.Name()))
			}
		}
	}
	return images, nil
}

// SaveConfig записывает выбор языка в ~/.config/nvrtier/config.json
func SaveConfig(isEnglish bool) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	configDir := filepath.Join(home, ".config", "nvrtier")
	_ = os.MkdirAll(configDir, 0755)

	filePath := filepath.Join(configDir, "config.json")
	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	cfg := AppConfig{IsEnglish: isEnglish}
	encoder := json.NewEncoder(file)
	_ = encoder.Encode(cfg)
}

// LoadConfig считывает сохраненный язык при старте утилиты
func LoadConfig() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	filePath := filepath.Join(home, ".config", "nvrtier", "config.json")
	file, err := os.Open(filePath)
	if err != nil {
		return false // Если файла нет, запускаем по дефолту (русский)
	}
	defer file.Close()

	var cfg AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return false
	}
	return cfg.IsEnglish
}
