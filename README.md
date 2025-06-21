<h1 align="center">
  <img src="https://raw.githubusercontent.com/ryu-ryuk/porterm/main/docs/logo.png" width="120" alt="Porterm Logo" /><br>
  <span style="color:#cdd6f4; vertical-align: middle; font-size: 2.5rem; margin-left: 16px;">porterm</span>
</h1>



<h6 align="center" style="color:#bac2de;">
  A Catppuccin-themed, interactive terminal portfolio and resume viewer.
</h6>

<p align="center">
  <a href="https://github.com/ryu-ryuk/porterm/stargazers"><img src="https://img.shields.io/github/stars/ryu-ryuk/porterm?colorA=1e1e2e&colorB=cba6f7&style=for-the-badge&logo=github&logoColor=cdd6f4"></a>
  <a href="https://github.com/ryu-ryuk/porterm/issues"><img src="https://img.shields.io/github/issues/ryu-ryuk/porterm?colorA=1e1e2e&colorB=f38ba8&style=for-the-badge&logo=github&logoColor=cdd6f4"></a>
  <a href="https://github.com/ryu-ryuk/porterm/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-89b4fa?style=for-the-badge&logo=gnu&logoColor=1e1e2e&colorA=1e1e2e"></a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-89b4fa?style=for-the-badge&logo=go&logoColor=white&colorA=1e1e2e" />
  <img src="https://img.shields.io/badge/Catppuccin-Mocha-cba6f7?style=for-the-badge&logo=catppuccin&logoColor=white&colorA=1e1e2e" />
  <img src="https://img.shields.io/badge/Terminal-BubbleTea-a6e3a1?style=for-the-badge&logo=gnu-bash&logoColor=white&colorA=1e1e2e" />
  <img src="https://img.shields.io/badge/Markdown-Glamour-fab387?style=for-the-badge&logo=markdown&logoColor=white&colorA=1e1e2e" />
  <img src="https://img.shields.io/badge/Style-Lipgloss-94e2d5?style=for-the-badge&logo=stylelint&logoColor=white&colorA=1e1e2e" />
  <img src="https://img.shields.io/badge/Maintained-Yes-89b4fa?style=for-the-badge&logo=github&logoColor=white&colorA=1e1e2e" />
</p>

<p align="center">
  <a href="https://scripts.alokranjan.me/porterm.sh">
    <img src="https://img.shields.io/badge/Install%20with%20curl-%23a6e3a1?style=for-the-badge&logo=gnubash&logoColor=1e1e2e&colorA=1e1e2e&colorB=a6e3a1" alt="Install with curl (Catppuccin Green)"/>
  </a>
</p>

<p align="center" style="color:#a6adc8; font-size: 14.5px; line-height: 1.6; max-width: 700px; margin: auto;">
  <strong style="color:#cdd6f4;">porterm</strong> is a Catppuccin-themed, interactive terminal portfolio and resume viewer.<br/>
  Built with <span style="color:#89b4fa;">Go</span>, styled with <span style="color:#cba6f7;">Catppuccin Mocha</span>, and powered by <span style="color:#a6e3a1;">Bubble Tea</span>, <span style="color:#fab387;">Glamour</span>, and <span style="color:#94e2d5;">Lipgloss</span>.
</p>

---

## Quick Install

```sh
curl -sL https://scripts.alokranjan.me/porterm.sh | bash
```

---

## Preview

<p align="center">
  <img src="docs/preview.gif" alt="Preview of porterm" />
</p>


## Overview

**porterm** is a modern, Catppuccin-themed terminal portfolio and resume viewer.  
It features animated banners, a dynamic About section, clickable project links, markdown resume rendering, and a clean, centered UI—built for developers who care about style and clarity in the terminal.

---

## Features

- **Catppuccin Mocha Theme**: All UI and markdown rendering use the Catppuccin Mocha palette for a cohesive, modern look.
- **Animated ASCII Banners**: Custom, non-distracting banner animations and glitch effects.
- **Centered Responsive Layout**: UI is always centered, adapting to any terminal size.
- **Dynamic About Section**: Rotating fun facts, skills, and passions.
- **Project Showcase**: Clickable project links with markdown descriptions.
- **Resume Viewer**: Scrollable, zoomable, and beautifully rendered markdown resume.
- **Badges & Webrings**: Show off your achievements and affiliations.
- **Keyboard Navigation**: Navigate sections with numeric keys.
- **Modern Stack**: Built with Bubble Tea, Lipgloss, and Glamour.

---

## Getting Started

### Prerequisites

- Go 1.22 or later
- A terminal emulator with OSC-8 hyperlink support recommended

### Installation

```sh
git clone https://github.com/ryu-ryuk/porterm.git
cd porterm
go build
```

### Run

```sh
go run .
```

---

## Usage

- `1` — About Me (dynamic fun facts, skills, passions)
- `2` — My Works (clickable project links)
- `3` — Resume (scrollable, zoomable, markdown)
- `4` — Webrings & Badges
- `q` — Quit
- `Esc` - Back
- Arrow keys / PgUp / PgDn — Scroll in resume

---

## Theming

All colors are from the [Catppuccin Mocha palette](https://catppuccin.com/).  
You can adapt the theme by editing `assets/glamour-catppuccin.json` and the color constants in `styles/theme.go`.
> PS. I have hardcoded some colors here & there ;>
---

## Customization

- **Resume**: Place your markdown resume in `content/resume.md`.
- **Projects**: Edit `assets/paste.txt` for your project list.
- **Badges & Webrings**: Update `views/badges.go` as needed.
- **Banner and vector art**: Customize in `model.go`.

---

## Architecture

- **Bubble Tea** — State management, event loop, and UI structure
- **Lipgloss** — Layout, color, and style primitives
- **Glamour** — Markdown rendering with custom styles


## Credits

- [Charmbracelet](https://github.com/charmbracelet) for Bubble Tea, Lipgloss, and Glamour
- [Catppuccin](https://catppuccin.com/) for the color palette

---

*This project is designed for developers who only care about terminal aesthetics.*
