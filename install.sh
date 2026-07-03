#!/bin/bash

# Останавливаем скрипт при любой ошибке
set -e

# Выбор языка установки
echo "================================================="
echo " Выберите язык установки / Select Install Language"
echo "================================================="
echo " 1) Русский"
echo " 2) English"
echo "================================================="
read -p " Выберите опцию / Select option (1-2): " LANG_CHOICE

if [ "$LANG_CHOICE" == "2" ]; then
    # ТЕКСТЫ НА АНГЛИЙСКОМ
    TXT_START="🚀 Starting NVRIER TERMINAL TIER-LIST installation..."
    TXT_NO_GO="❌ Error: Go is not installed. Please install go (e.g., sudo pacman -S go)."
    TXT_CHECK_DIR="📂 Checking directory for images: $HOME/nvrtier_pictures"
    TXT_CREATED_DIR="✅ Created directory: $HOME/nvrtier_pictures"
    TXT_DIR_EXISTS="ℹ️  Directory already exists, skipping creation."
    TXT_TIDY="📦 Synchronizing Go dependencies..."
    TXT_BUILD="🏗️  Compiling the program..."
    TXT_BUILD_OK="✅ Build completed successfully!"
    TXT_MOVE="🔒 Moving the utility to /usr/local/bin/..."
    TXT_SUDO="ℹ️  Please enter your sudo password to complete the installation:"
    TXT_DONE_TITLE="🎉 INSTALLATION COMPLETED SUCCESSFULLY! 🎉"
    TXT_DONE_COPY="🖼️  Copy your images (.png/.jpg) to folder: $HOME/nvrtier_pictures"
    TXT_DONE_RUN="🔥 Run the utility from anywhere using command: nvrtier"
else
    # ТЕКСТЫ НА РУССКОМ
    TXT_START="🚀 Начинаем установку NVRIER ТЕРМИНАЛЬНЫЙ ТИР-ЛИСТ..."
    TXT_NO_GO="❌ Ошибка: Go не установлен в системе. Пожалуйста, установите go (например, sudo pacman -S go)."
    TXT_CHECK_DIR="📂 Проверяем директорию для изображений: $HOME/nvrtier_pictures"
    TXT_CREATED_DIR="✅ Создана папка: $HOME/nvrtier_pictures"
    TXT_DIR_EXISTS="ℹ️  Папка уже существует, пропускаем создание."
    TXT_TIDY="📦 Синхронизируем зависимости Go..."
    TXT_BUILD="🏗️  Компилируем программу..."
    TXT_BUILD_OK="✅ Сборка успешно завершена!"
    TXT_MOVE="🔒 Перемещаем утилиту в системную папку /usr/local/bin/..."
    TXT_SUDO="ℹ️  Пожалуйста, введите ваш пароль sudo для завершения установки:"
    TXT_DONE_TITLE="🎉 УСТАНОВКА УСПЕШНО ЗАВЕРШЕНА! 🎉"
    TXT_DONE_COPY="🖼️  Скопируйте ваши картинки (.png/.jpg) в папку: $HOME/nvrtier_pictures"
    TXT_DONE_RUN="🔥 Запускайте утилиту из любого места терминала командой: nvrtier"
fi

# Выполнение установки с выбранной локализацией
echo ""
echo "$TXT_START"
echo ""

# 1. Проверяем Go
if ! command -v go &> /dev/null; then
    echo "$TXT_NO_GO"
    exit 1
fi

# 2. Создаем папку в домашней директории пользователя
TARGET_PICTURES_DIR="$HOME/nvrtier_pictures"
echo "$TXT_CHECK_DIR"
if [ ! -d "$TARGET_PICTURES_DIR" ]; then
    mkdir -p "$TARGET_PICTURES_DIR"
    echo "$TXT_CREATED_DIR"
else
    echo "$TXT_DIR_EXISTS"
fi

# 3. Синхронизируем зависимости
echo "$TXT_TIDY"
go mod tidy

# 4. Компилируем бинарник
echo "$TXT_BUILD"
go build -o nvrtier main.go storage.go image.go
echo "$TXT_BUILD_OK"

# 5. Перемещаем в систему
echo "$TXT_MOVE"
echo "$TXT_SUDO"
sudo mv nvrtier /usr/local/bin/nvrtier
sudo chmod +x /usr/local/bin/nvrtier

# Финальные инструкции
echo ""
echo "=================================================================="
echo " $TXT_DONE_TITLE"
echo "=================================================================="
echo " $TXT_DONE_COPY"
echo " $TXT_DONE_RUN"
echo "=================================================================="
