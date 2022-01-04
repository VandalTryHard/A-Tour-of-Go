// Загружает веб страницу
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Использование каналов в программе для вывода размера веб-страниц page 430
type Page struct { //Для хранения URL с размером страницы
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Geting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://example.com/",
		"https://golang.org/",
		"https://golang.org/doc"}

	for _, url := range urls {
		go responseSize(url, pages) // Вызываем responseSize для каждого URL.
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size) //Получение данных из канала по одному разу для каждой отправки, выполненной responseSize.
	}
}
