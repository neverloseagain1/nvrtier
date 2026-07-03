package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/nsf/termbox-go"
)

type model struct {
	tiers         map[string][]string
	pool          []string
	tierOrder     []string
	selectedIndex int
	isFinished    bool
	isEnglish     bool 
}

func tbPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func tbLine(startX, endX, y int, fg, bg termbox.Attribute, ch rune) {
	for x := startX; x <= endX; x++ {
		termbox.SetCell(x, y, ch, fg, bg)
	}
}

func renderUI(m *model) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	termWidth, _ := termbox.Size()
	if termWidth < 40 {
		termWidth = 80
	}

	rightEdge := termWidth - 2

	tierColors := map[string]termbox.Attribute{
		"1": termbox.ColorRed,
		"2": termbox.ColorYellow, 
		"3": termbox.ColorWhite,  
		"4": termbox.ColorGreen,
		"5": termbox.ColorBlue,
	}

	// 1. РИСУЕМ КРЫШУ ТАБЛИЦЫ
	termbox.SetCell(1, 1, '╔', termbox.ColorWhite, termbox.ColorDefault)
	tbLine(2, 16, 1, termbox.ColorWhite, termbox.ColorDefault, '═')
	termbox.SetCell(17, 1, '╦', termbox.ColorWhite, termbox.ColorDefault)
	tbLine(18, rightEdge-1, 1, termbox.ColorWhite, termbox.ColorDefault, '═')
	termbox.SetCell(rightEdge, 1, '╗', termbox.ColorWhite, termbox.ColorDefault)

	// Строим сетку тиров
	for tIdx, tier := range m.tierOrder {
		startLine := 2 + (tIdx * 4) 
		bgColor := tierColors[tier]

		// 2. РИСУЕМ ЦВЕТНУЮ ПЛАШКУ С ЗАКРЫТЫМИ БОКОВИНАМИ
		for y := startLine; y < startLine+3; y++ {
			termbox.SetCell(1, y, '║', termbox.ColorWhite, termbox.ColorDefault)
			tbLine(2, 16, y, termbox.ColorBlack, bgColor, ' ')
			termbox.SetCell(17, y, '║', termbox.ColorWhite, termbox.ColorDefault)
		}

		// Выравнивание названий плашек по центру
		tierText := fmt.Sprintf("ТИР %s", tier)
		textX := 5
		if m.isEnglish {
			tierText = fmt.Sprintf("TIER %s", tier)
			textX = 4 
		}
		tbPrint(textX, startLine+1, termbox.ColorBlack, bgColor, tierText)

		// 3. ЗАКРЫВАЕМ ВНЕШНЮЮ ПРАВУЮ СТЕНКУ ТИРА
		termbox.SetCell(rightEdge, startLine, '║', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(rightEdge, startLine+1, '║', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(rightEdge, startLine+2, '║', termbox.ColorWhite, termbox.ColorDefault)

		// 4. РИСУЕМ ГОРИЗОНТАЛЬНЫЕ РАЗДЕЛИТЕЛИ С ПРАВИЛЬНЫМИ СТЫКАМИ
		if tIdx < 4 { 
			termbox.SetCell(1, startLine+3, '╠', termbox.ColorWhite, termbox.ColorDefault)
			tbLine(2, 16, startLine+3, termbox.ColorWhite, termbox.ColorDefault, '═')
			termbox.SetCell(17, startLine+3, '╬', termbox.ColorWhite, termbox.ColorDefault)
			tbLine(18, rightEdge-1, startLine+3, termbox.ColorWhite, termbox.ColorDefault, '═')
			termbox.SetCell(rightEdge, startLine+3, '╣', termbox.ColorWhite, termbox.ColorDefault)
		}

		if len(m.tiers[tier]) == 0 {
			emptyText := ""
			if m.isEnglish {
				emptyText = ""
			}
			tbPrint(20, startLine+1, termbox.ColorWhite, termbox.ColorDefault, emptyText)
		}
	}

	// 5. РИСУЕМ ДНО ТАБЛИЦЫ
	lastLine := 2 + (5 * 4) - 1
	termbox.SetCell(1, lastLine, '╚', termbox.ColorWhite, termbox.ColorDefault)
	tbLine(2, 16, lastLine, termbox.ColorWhite, termbox.ColorDefault, '═')
	termbox.SetCell(17, lastLine, '╩', termbox.ColorWhite, termbox.ColorDefault)
	tbLine(18, rightEdge-1, lastLine, termbox.ColorWhite, termbox.ColorDefault, '═')
	termbox.SetCell(rightEdge, lastLine, '╝', termbox.ColorWhite, termbox.ColorDefault)

	// 6. УПРАВЛЕНИЕ И ИНФО-ПАНЕЛЬ ВНИЗУ ПО ЦЕНТРУ
	if m.isFinished || len(m.pool) == 0 {
		finishText := "🎉 Отличная работа! Все элементы распределены. Нажмите [Q] для выхода."
		if m.isEnglish {
			finishText = "🎉 Great job! All items are sorted. Press [Q] to exit."
		}
		centerX := (termWidth - len([]rune(finishText))) / 2
		if centerX < 1 {
			centerX = 1
		}
		tbPrint(centerX, 23, termbox.ColorGreen | termbox.AttrBold, termbox.ColorDefault, finishText)
	} else {
		currentImg := m.pool[m.selectedIndex]
		
		selectLabel := "Выбор: "
		itemLabel := fmt.Sprintf(" (Элемент %d из %d)", m.selectedIndex+1, len(m.pool))
		if m.isEnglish {
			selectLabel = "Selected: "
			itemLabel = fmt.Sprintf(" (Item %d of %d)", m.selectedIndex+1, len(m.pool))
		}
		fullInfoText := selectLabel + filepath.Base(currentImg) + itemLabel
		
		infoX := (termWidth - len([]rune(fullInfoText))) / 2
		if infoX < 1 {
			infoX = 1
		}
		tbPrint(infoX, 23, termbox.ColorCyan | termbox.AttrBold, termbox.ColorDefault, fullInfoText)

		controlsText := "Управление: [← / →] Выбор | [1-5] Отправить в тир | [L] Язык | [Q] Выход"
		if m.isEnglish {
			controlsText = "Controls: [← / →] Select | [1-5] Send to Tier | [L] Language | [Q] Exit"
		}
		
		controlsX := (termWidth - len([]rune(controlsText))) / 2
		if controlsX < 1 {
			controlsX = 1
		}
		tbPrint(controlsX, 24, termbox.ColorYellow, termbox.ColorDefault, controlsText)
	}

	termbox.Flush()

	// Микропауза 5мс для фиксации текстовых линий
	time.Sleep(5 * time.Millisecond)

	// 7. Накатываем оригинальные пиксели поверх собранной разметки
	for tIdx, tier := range m.tierOrder {
		startLine := 2 + (tIdx * 4)
		if len(m.tiers[tier]) > 0 {
			for cIdx, tPath := range m.tiers[tier] {
				posX := 19 + (cIdx * 9)
				RenderKittyImage(tPath, posX, startLine+1, 8, 3) 
			}
		}
	}

	// Превью склада
	if !m.isFinished && len(m.pool) > 0 {
		currentImg := m.pool[m.selectedIndex]
		RenderKittyImage(currentImg, 3, 26, 16, 6)
	}
}

func main() {
	dir, err := GetStorageDir()
	if err != nil {
		fmt.Printf("❌ Ошибка определения директории: %v\n", err)
		os.Exit(1)
	}

	// Железно создаем папку nvrtier_pictures, если её вдруг стерли или её нет в системе
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755)
	}

	images, err := GetImagesList(dir)
	if err != nil || len(images) == 0 {
		fmt.Printf("\n❌ ОШИБКА: Папка для картинок пуста: %s\n", dir)
		fmt.Printf("Пожалуйста, закиньте ваши логотипы (.png/.jpg) в директорию ~/nvrtier_pictures и запустите программу заново!\n")
		os.Exit(1)
	}

	err = termbox.Init()
	if err != nil {
		fmt.Printf("Ошибка инициализации TUI: %v\n", err)
		os.Exit(1)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	// АВТОЗАГРУЗКА НАСТРОЕК: считываем флаг языка из сохраненного конфига
	savedLanguage := LoadConfig()

	m := model{
		tiers: map[string][]string{
			"1": {}, "2": {}, "3": {}, "4": {}, "5": {},
		},
		pool:          images,
		tierOrder:     []string{"1", "2", "3", "4", "5"},
		selectedIndex: 0,
		isFinished:    false,
		isEnglish:     savedLanguage, // Применяем сохраненный язык
	}

	fmt.Print("\x1b[2J\x1b[H")
	renderUI(&m)

	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			if ev.Ch == 'q' || ev.Ch == 'Q' || ev.Key == termbox.KeyCtrlC {
				break
			}

			// Смена языка теперь автоматически ПЕРЕЗАПИСЫВАЕТ файл конфигурации config.json на лету!
			if ev.Ch == 'l' || ev.Ch == 'L' {
				m.isEnglish = !m.isEnglish
				SaveConfig(m.isEnglish) // Железно сохраняем выбор на диск
				renderUI(&m) 
				continue
			}

			if m.isFinished {
				continue
			}

			if ev.Key == termbox.KeyArrowLeft {
				if m.selectedIndex > 0 {
					m.selectedIndex--
					renderUI(&m)
				}
			}
			if ev.Key == termbox.KeyArrowRight {
				if m.selectedIndex < len(m.pool)-1 {
					m.selectedIndex++
					renderUI(&m)
				}
			}

			if ev.Ch >= '1' && ev.Ch <= '5' {
				if len(m.pool) == 0 {
					continue
				}

				tierKey := string(ev.Ch)
				img := m.pool[m.selectedIndex]

				m.tiers[tierKey] = append(m.tiers[tierKey], img)
				m.pool = append(m.pool[:m.selectedIndex], m.pool[m.selectedIndex+1:]...)

				if len(m.pool) == 0 {
					m.isFinished = true
				} else if m.selectedIndex >= len(m.pool) {
					m.selectedIndex = len(m.pool) - 1
				}

				ClearKittyImages()
				renderUI(&m)
			}
		}
	}
	ClearKittyImages()
	fmt.Print("\x1b[2J\x1b[H")
}
