# Go Image Scraper

A simple Go project that scrapes all images from any website and downloads them into an `images/` folder.  
This project is perfect for beginners learning Go, HTTP requests, file handling, and HTML parsing.

---

## 🚀 Features

- Fetches any webpage using `http.Get`
- Parses HTML using **goquery**
- Extracts all `<img>` tags
- Automatically fixes image URLs (relative, protocol missing, etc.)
- Downloads all images locally with unique filenames
- Simple, clean, beginner-friendly Go code

---

## 📦 Requirements

Before running the project, install dependencies:

```bash
go mod tidy
This installs goquery and other required packages.

▶️ How to Run
go run main.go


Then enter any website URL, for example:

https://example.com


All images will be downloaded into the images/ folder automatically.

📁 Project Structure
go-image-scraper/
│── main.go
│── go.mod
│── go.sum
│── README.md
└── images/          # Downloaded images (ignored by Git)

🧠 How It Works (Quick Explanation)

The user types a website URL.

The program sends an HTTP GET request.

goquery parses the HTML.

All <img> tags are found.

Each image URL is cleaned and normalized.

The file is downloaded with a unique filename such as img_0.jpg, img_1.png, etc.

All files are saved inside /images.

⚙️ Technologies Used

Go

goquery (github.com/PuerkitoBio/goquery)

Standard packages: fmt, net/http, os, io, strings

📝 License

MIT License — feel free to use, modify, and improve the project.

🤝 Contributing

Pull requests are welcome!
If you want to add new features (concurrency, better URL parsing, duplicate checking), feel free to contribute.

⭐ Show Support

If you like the project, consider giving the repository a star ⭐ on GitHub!
