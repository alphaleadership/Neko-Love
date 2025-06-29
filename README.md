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
  "url": "http://localhost:3030/api/images/neko/04.webp"
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

## 🎨 Filters

The API also includes an **image filtering endpoint**:

```
GET /api/v4/filters/:filter?image=<url>
```

### ✅ Supported formats

- JPEG
- PNG
- WEBP
- **GIF** (animated) – frame-by-frame filtering is supported.

### 🧪 Available Filters

| Filter        | Description                                 |
|---------------|---------------------------------------------|
| `blurple`     | Applies a Discord blurple color overlay     |
| `fuchsia`     | Fuchsia / pinkish recoloring                |
| `glitch`      | RGB split + distortion effect               |
| `neon`        | Glowing edge highlights                     |
| `deepfry`     | Chaotic contrast & saturation (meme style)  |
| `posterize`   | Reduces color depth to retro/flat look      |
| `pixelate`    | Pixel art style                             |
| `vaporwave`   | Retro 90s purple-cyan aesthetic             |
| `anime_outline` | Anime-style black outlines               |

---

### 📷 Example Renders

#### PNG

Original | With `blurple`
:--:|:--:
![Original PNG](examples/original.png) | ![Blurple PNG](examples/blurple.png)

#### GIF

Original | With `blurple`
:--:|:--:
![Original GIF](examples/original.gif) | ![Blurple GIF](examples/blurple.gif)

The following filters are currently implemented and showcased in the [`Examples`](examples/) folder of the project:

You can test a filter with a public image URL like this:

```
http://localhost:3030/api/v4/filters/deepfry?image=https://example.com/image.png
```

You can also pass animated GIFs, and the API will apply the filter to each frame before returning a new animated result.

---

## 🤝 Contributing

Want to add a new image category?

1. Create a folder inside `assets/<name>`
2. Add your image files there

No extra routes need to be added to the code since they are managed automatically.

Thanks to everyone who used the original Neko-Love, and to all those who want to bring it back with a new twist ✨

---

## 🔍 Example API Call

To get a random image (returns `{ "url": "http://localhost:3030/api/images/neko/04.webp" }`):

```
GET http://localhost:3030/api/v4/neko
```

To access the image directly (after receiving the URL from the JSON response):

---

## 🔄 Another version of the project:

- Rust (Framework Axum) by [reinacchi](https://github.com/reinacchi/Neko-Love/tree/rust)
