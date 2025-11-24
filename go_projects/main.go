package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var url string
	fmt.Print("Enter website URL: ")
	fmt.Scanln(&url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	os.Mkdir("images", 0755)


	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")

		if src == "" {
			return
		}

		
		if strings.HasPrefix(src, "//") {
			src = "https:" + src
		} else if strings.HasPrefix(src, "/") {
			src = url + src
		}

		fmt.Println("Downloading:", src)
		downloadImage(src, i)
	})

	fmt.Println("✔ Done! Check the images folder.")
}

func downloadImage(url string, index int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed:", url)
		return
	}
	defer resp.Body.Close()

	fileName := fmt.Sprintf("images/img_%d%s", index, path.Ext(url))

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("File error:", err)
		return
	}
	defer file.Close()

	io.Copy(file, resp.Body)
}
