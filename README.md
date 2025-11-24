# Go Image Scraper

  A simple Go project that scrapes all images from any website and downloads them into an images/ folder.
  Perfect for beginners learning Go, HTTP requests, file handling, and HTML parsing.

## 🚀 Features

Fetches any webpage using http.Get

Parses HTML using goquery

Extracts all <img> tags

Automatically fixes relative or missing-protocol image URLs

Downloads images with unique filenames

Clean, beginner-friendly Go code

## 📦 Requirements

Install the dependencies:

go mod tidy

This installs goquery and other required packages.

## ▶️ How to Run

go run main.go

Enter any website URL, for example:

https://example.com

All images will be downloaded into the images/ folder automatically.

## 📁 Project Structure

go-image-scraper/
│── main.go
│── go.mod
│── go.sum
│── README.md
└── images/ # Downloaded images (ignored by Git)

## 🧠 How It Works

##### You enter a website URL.

The program sends an HTTP GET request.

goquery parses the page’s HTML.

All <img> tags are extracted.

URLs are normalized (absolute, fixed protocol, etc.).

Images are downloaded as:
img_0.jpg, img_1.png, img_2.jpeg, ...

Every file is saved inside images/.

## ⚙️ Technologies Used

##### Go

##### goquery (github.com/PuerkitoBio/goquery)

##### Standard libraries: fmt, net/http, os, io, strings

## 📝 License

MIT License — free to use, modify, and improve.

## 🤝 Contributing

Pull requests are welcome!
If you want to add features like concurrency, duplicate checking, or better URL handling, feel free to contribute.

## ⭐ Show Support

If you like this project, give it a ⭐ on GitHub!
