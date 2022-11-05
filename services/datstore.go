package services

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
	"url-shortener/models"

	"github.com/speps/go-hashids"
)

type urlStore struct{}

const (
	filePath = "./urlmapping.txt"
)

func New() URL {
	return &urlStore{}
}

func (u *urlStore) GetShortURL(url models.URL) (string, error) {

	_, e := os.Stat(filePath)
	if e != nil {
		// Checking if the given file exists or not
		if os.IsNotExist(e) {
			log.Println("File not Found !!")

		}
		ff, errr := os.Create("urlmapping.txt")
		if errr != nil {
			log.Fatal(errr)
		}
		defer ff.Close()

	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, eachLine := range text {
		strArr := strings.Split(eachLine, "|")
		if url.LongURL == strArr[0] {
			return strArr[1], nil
		}
	}

	// generate short URL
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	id, _ := h.Encode([]int{int(time.Now().Unix())})
	shortURL := "http://localhost:8880/" + id

	writeToFile(shortURL, url)

	return shortURL, nil
}

func writeToFile(shortURL string, url models.URL) {
	data := (url.LongURL + "|" + shortURL + "\n")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	textWriter := bufio.NewWriter(file)
	_, err = textWriter.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
	textWriter.Flush()

}
