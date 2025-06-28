# Neko-Love API V4 (Go / Fiber Rewrite)

> A community-powered rewrite of the original Neko-Love API, rebuilt in Go with the Fiber web framework.

---

## 🌟 What is this?

This project is a **modern reimplementation of the original Neko-Love API**, which was once hosted at `neko-love.xyz`. It served random anime-style images like "neko", "hug", "kiss", and more — often used in Discord bots, anime projects, and other community tools.

The original API was written in Node.js using Koa. This version is a **fresh and solid base built with Go and the Fiber framework**, designed for others to easily clone, customize, and host themselves.

---

## 🔧 Goals

- Provide a **lightweight and fast** REST API to serve random images.
- Offer a **clean modular structure** for adding new routes, categories, and assets.
- **SFW only**: This version contains strictly Safe For Work content.

---

## ⚠️ Please Note

This API **is not hosted by the original author**.

> It is provided as an open-source base only. You are free to clone, modify, self-host, and integrate it into your own projects.

---

## 🚀 Quick Overview

- Each route returns a random image from a local folder (e.g. `/assets/neko/`)
- Example JSON response:

```json
{
  "url": "/images/neko/neko_04.png"
}
```

- Images are served at `/images/<category>/<image>`

---

## 🚩 Why a new version?

> After the original API was shut down, several community members asked if they could bring it back. Unfortunately, the original source code was lost. This rewrite aims to provide a fresh, modern foundation.

- ✅ Easier to maintain
- ✅ Fast and lightweight (Go + Fiber)
- ✅ Clean structure for contributions

---

## 💻 Run Locally

### Requirements

Make sure you have **Go installed** (version 1.18+ recommended):  
→ [Download Go](https://golang.org/dl/)

---

### Installation

1. Clone the repository:

```bash
git clone https://github.com/Otaku17/neko-love.git
cd neko-love-go
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the API:

```bash
go run main.go
```

4. Add your images in the corresponding folders inside the `assets/` directory (e.g. `assets/neko/`, `assets/hug/`, etc.)

---

## 🤝 Contributing

Want to add a new image category?

1. Create a folder inside `assets/<name>`
2. Add your image files there
3. Add a new route in `routes/image_routes.go`

Example:

```go
v1.Get("/pat", handlers.GetRandomImage("assets/pat"))
```

---

Thanks to everyone who used the original Neko-Love, and to all those who want to bring it back with a new twist ✨

## 🔍 Example API Call

To get a random image (returns `{ "url": "/images/neko/04.webp" }`):

```
GET http://localhost:3030/api/v4/neko
```

To access the image directly (after receiving the URL from the JSON response):

```
http://localhost:3030/images/neko/04.webp
```
