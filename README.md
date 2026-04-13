# 🎨 ascii-art-web-stylize

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)
![HTML](https://img.shields.io/badge/HTML-15%25-E34F26?style=flat&logo=html5&logoColor=white)
![CSS](https://img.shields.io/badge/CSS-10%25-1572B6?style=flat&logo=css3&logoColor=white)
![Repo Size](https://img.shields.io/badge/Repo%20Size-9.20%20kB-brightgreen?style=flat)
![Standard Library](https://img.shields.io/badge/Dependencies-Standard%20Library%20Only-00ADD8?style=flat&logo=go&logoColor=white)

> A beautifully styled, interactive web application that transforms plain text into expressive ASCII art — built with Go, designed for humans.

---

## 📋 Table of Contents

- [About the Project](#-about-the-project)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [How to Run](#how-to-run)
- [Usage](#-usage)
- [Design Principles](#-design-principles)
- [Allowed Packages](#-allowed-packages)
- [Learning Outcomes](#-learning-outcomes)
- [Authors](#-authors)

---

## 📖 About the Project

**ascii-art-web-stylize** is an extension of the [ascii-art-web](https://acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize) project, rebuilt with a strong focus on user experience, visual design, and interactivity.

The original project provided the core functionality of rendering ASCII art in a browser. This iteration takes it further — transforming a functional tool into a polished, user-friendly web application. The goal was to address a common shortcoming in utility-first tools: they work, but they don't *feel* good to use.

This project solves that by applying fundamental principles of human-computer interface design, ensuring that every user interaction — from typing text to selecting a banner style — feels intuitive, responsive, and satisfying.

**Why it was built:**

- To practise applying CSS styling principles to a real Go web server
- To explore what makes a web interface genuinely usable, not just functional
- To understand the relationship between back-end logic and front-end presentation
- To deliver meaningful user feedback rather than silent success or failure

---

## ✨ Features

- **Responsive Layout** — the interface adapts gracefully to different screen sizes, from desktop monitors to mobile devices, without breaking the design or usability.

- **Interactive UI** — form elements, buttons, and controls respond visually to user actions (hover states, focus rings, active effects), making the application feel alive and reactive.

- **Real-time User Feedback** — the application provides clear visual cues for loading states, successful outputs, and errors, so the user always knows what is happening.

- **Multiple Banner Styles** — users can choose from a selection of ASCII art banner templates (e.g., `standard`, `shadow`, `thinkertoy`) to customise the output style.

- **Readable Text at All Times** — regardless of background colours, themes, or styling configurations applied, all text content maintains sufficient contrast and remains fully legible.

- **Consistent Visual Language** — typography, spacing, colour use, and component behaviour follow a unified design system across every page and state.

- **Clean Error Handling** — invalid inputs or server-side issues are communicated to the user through styled, friendly error messages rather than raw HTTP responses.

- **Accessible Colour Choices** — colour combinations throughout the UI are selected with readability and accessibility in mind, avoiding clashes or low-contrast pairings.

---

## 🛠 Tech Stack

| Technology | Role | Details |
|---|---|---|
| **Go** | Backend | HTTP server, request routing, ASCII art rendering logic. Uses only the Go standard library — no third-party frameworks or packages. |
| **HTML** | Structure | Semantic markup defining the layout and content of the web pages. Properly linked to CSS for styling and served by the Go template engine. |
| **CSS** | Styling & UX | Responsible for visual design, layout, responsiveness, animations, and interactive states. Linked externally to HTML templates for clean separation of concerns. |

---

## 📁 Project Structure

```
ascii-art-web-stylize/
├── cmd/
│   └── main.go                 # Entry point — starts the HTTP server
├── fonts/
│   ├── shadow.txt              # Shadow ASCII banner font
│   ├── standard.txt            # Standard ASCII banner font
│   └── thinker_toy.txt         # Thinkertoy ASCII banner font
├── internal/
│   ├── config/
│   │   └── app.go              # Application configuration logic
│   ├── handler/
│   │   └── handler.go          # HTTP handler functions (GET, POST routes)
│   ├── service/
│   │   ├── service.go          # Business logic orchestration
│   │   └── transform_ascii.go  # High-level ASCII transformation logic
│   └── Transformer/
│       ├── covert_to_ascii.go  # Logic for character mapping
│       ├── format_ascii.go     # Formatting and alignment logic
│       ├── read_font.go        # Font file parsing
│       ├── render_ascii_output.go # Output generation logic
│       ├── split_text.go       # Input string processing
│       └── transformer.go      # Transformer interface or core engine
├── pkg/
│   └── models/
│       └── models.go           # Shared data structures (e.g., ASCIIRequest)
├── web/
│   ├── static/
│   │   ├── ascii_output.html   # Static version of output page
│   │   ├── index.html          # Static landing page
│   │   └── style.css           # Global stylesheets
│   └── templates/
│       ├── index.html          # Dynamic HTML template
│       ├── lajose_image.jpeg   # Template-specific image asset
│       └── style.css           # Template-specific stylesheet
├── go.mod                      # Go module file
└── README.md                   # Project documentation
```

> **Note:** The exact structure may vary slightly depending on refactoring decisions made during development. The tree above represents the recommended organisation.

---

## 🚀 Getting Started

### Prerequisites

Before running this project, make sure you have the following installed on your machine:

- **Go 1.18 or higher** — [Download Go](https://go.dev/dl/)

You can verify your Go installation by running:

```bash
go version
```

No other dependencies are required. This project uses only the Go standard library.

---

### Installation

**1. Clone the repository:**

```bash
git clone https://acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize.git
```

**2. Navigate into the project directory:**

```bash
cd ascii-art-web-stylize
```

**3. Verify the module is initialised (it should already be):**

```bash
cat go.mod
```

---

### How to Run

Start the web server using the Go toolchain:

```bash
go run ./cmd
```



Once the server is running, open your browser and visit:

```
http://localhost:8081
```

> The default port is `8081`. If your application uses a different port, update the URL accordingly.

To stop the server, press `Ctrl + C` in your terminal.

---

## 💻 Usage

1. **Open the app** in your browser at `http://localhost:8081`.
2. **Type your text** into the input field on the main page.
3. **Select a banner style** from the available options (Standard, Shadow, Thinkertoy).
4. **Click the "Generate" button** to render your ASCII art.
5. **View the output** displayed directly on the page in a styled, readable format.

---



---

## 🎨 Design Principles

The styling and UX decisions in this project were guided by the following principles:

**Colour Contrast and Readability**
All text elements are placed against backgrounds with sufficient contrast to ensure legibility for all users, including those with visual impairments. No decorative colour choice was allowed to compromise the readability of content.

**Responsiveness**
The layout was built with a mobile-first mindset. Using CSS Flexbox and/or Grid, the interface reflows gracefully across a range of viewport sizes — from small phones to wide desktop screens — without horizontal scrolling or broken layouts.

**Interactivity and Feedback**
Every interactive element — buttons, inputs, selectors — provides visual feedback on hover, focus, and active states. Loading states and result transitions are handled smoothly to prevent the interface from feeling static or unresponsive.

**Consistency**
A unified set of CSS variables governs colours, font sizes, spacing, and border radii across the entire application. This ensures that the interface looks and feels cohesive on every page and in every state.

**Simplicity**
The interface is clean and purposeful. Visual noise was deliberately avoided — each element on the page earns its place by serving the user's goal of generating and reading ASCII art.

---

## 📦 Allowed Packages

This project was built using **Go's standard library only**. No third-party packages or external dependencies are imported.

Standard library packages used may include (but are not limited to):

```
net/http       — HTTP server and request handling
html/template  — Server-side HTML template rendering
fmt            — Formatted I/O
os             — File system access (reading banner files)
strings        — String manipulation
log            — Logging server events and errors
```

This constraint was intentional — it reinforces a deep understanding of Go's built-in capabilities and avoids unnecessary dependency complexity.

---

## 🎓 Learning Outcomes

Working through this project develops practical understanding in the following areas:

- **Human-Computer Interface Design** — understanding what makes an interface intuitive, accessible, and satisfying to use, beyond just functional correctness.

- **CSS Fundamentals** — applying selectors, the box model, layout systems (Flexbox, Grid), pseudo-classes, and custom properties to produce a polished visual result.

- **Linking CSS and HTML** — correctly referencing external stylesheets from HTML templates, understanding how the browser resolves static asset paths, and how Go serves static files.

- **Principles of Good Website Design** — applying concepts like visual hierarchy, whitespace, colour theory, consistency, and responsiveness to a real project.

- **Go Web Server Fundamentals** — handling HTTP requests and responses, serving static files, rendering HTML templates, and organising a Go web application cleanly.

- **Separation of Concerns** — keeping logic (Go handlers), structure (HTML templates), and presentation (CSS) cleanly separated and properly connected.

---



## 👥 Authors

This project was built and maintained by:

**Divine Ositadinma**
> 📎 [Profile Link — https://acad.learn2earn.ng/git/dositadi]

**Ajose Lawal**
> 📎 [Profile Link — https://acad.learn2earn.ng/git/lajose]

---

