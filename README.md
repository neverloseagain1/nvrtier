# NVRTIER - Terminal Tier-List 🔥

[Читать на русском языке (Read in Russian)](README_RU.md)

An interactive, full-screen Terminal User Interface (TUI) utility for creating tier lists directly inside the Kitty terminal. It renders source images in 100% original pixel quality using the native Kitty Graphics Protocol.

# !ATTENTION!

This project was partially created by AI, so there may be bugs. If you encounter any bugs, please open an issue.

## Features

* **Original Pixel Quality**: Native GPU-accelerated image rendering via the Kitty Protocol.
* **Interactive TUI**: Independent fullscreen buffer mode driven by the `termbox-go` engine. No `Enter` key required for navigation.
* **Localization**: Seamless toggle between English and Russian interfaces on the fly.
* **Persistent Settings**: Saves your language preference automatically to a local configuration file.

## Prerequisites

1. **Terminal**: This utility strictly requires the **Kitty terminal** or any terminal emulator that fully implements the Kitty Graphics Protocol.
2. **Go Environment**: Go compiler version 1.16 or higher must be installed on your system.

## Installation

Download repository and run the interactive installation script from the project root directory:

```bash
git clone https://github.com/neverloseagain1/nvrtier

cd nvrtier

chmod +x install.sh

./install.sh
```

During installation, select `2` for English. The script will automatically download dependencies, compile the source code, create the image folder, and move the binary to `/usr/local/bin/`.

## Usage

1. Copy your target images (`.png`, `.jpg`, or `.jpeg`) into the following directory inside your home catalog:
   ```bash
   ~/nvrtier_pictures
   ```
2. Launch the application from any directory in your shell:
   ```bash
   nvrtier
   ```

## Controls

* `Left Arrow` / `Right Arrow`: Navigate through the current image stock repository.
* `1` - `5`: Assign the selected image to the corresponding Tier row.
* `L`: Toggle user interface language (English ↔ Russian).
* `Q` / `Ctrl+C`: Clear terminal graphic layers and exit the utility.