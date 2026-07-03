package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg" // Гарантирует декодирование JPG
	"image/png"    // Гарантирует работу с PNG
	"os"
)

// RenderKittyImage берет любой JPG/PNG, делает из него чистый PNG в памяти и шлет в Kitty
func RenderKittyImage(filePath string, x, y, widthChars, heightRows int) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// 1. Декодируем исходную картинку (Go сам поймет, JPG это или PNG)
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}

	// 2. Кодируем её на лету в формат PNG прямо в буфер оперативной памяти
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return
	}

	// 3. Кодируем полученный PNG-поток в base64
	b64Data := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Перемещаем курсор текстового терминала в нужные координаты
	fmt.Printf("\x1b[%d;%dH", y, x)

	// Шлем в Kitty команду вывода 100% пикселей
	fmt.Printf("\x1b_Ga=T,f=100,c=%d,r=%d;%s\x1b\\", widthChars, heightRows, b64Data)
}

// ClearKittyImages полностью стирает наложенные пиксели перед обновлением экрана
func ClearKittyImages() {
	fmt.Print("\x1b_Ga=d;\x1b\\")
}
