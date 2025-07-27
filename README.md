# 🔗 Go Runtime URL Shortener

This is a simple **URL shortener API** built in Go that stores data **in-memory at runtime** (like a lightweight version of MongoDB). It was created as my **first Go project** to get hands-on with basic web server handling, hashing, and JSON request/response handling in Golang.

---

## ✨ Features

- 🔐 Generates short URLs using `MD5` hash
- 💾 Stores original and short URLs in a runtime map (in-memory DB)
- 🧭 Redirects short URLs to the original URLs
- ⚡ Built with pure `net/http` package, no external libraries
- ⏳ Timestamp stored with each URL

---

## 📦 Tech Stack

- **Language**: Go (Golang)
- **Hashing**: `crypto/md5`
- **Data Storage**: In-memory map (Go map)
- **JSON Handling**: `encoding/json`
- **Web Server**: `net/http`

---

## 🚀 How It Works

### 1. Start the server

go run main.go
2. Create a Short URL
Endpoint: POST /shorten

Request Body (JSON):

json

Edit
{
  "url": "https://example.com/some/long/link"
}
3. Redirect to Original URL
Endpoint: GET /redirect/{short_url}

Example:

http

Edit
GET http://localhost:8080/redirect/a1b2c3d4
