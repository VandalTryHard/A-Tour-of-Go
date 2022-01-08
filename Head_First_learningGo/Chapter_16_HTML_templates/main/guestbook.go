// отображение основной страницы гостевой книги
package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Guestbook - структура, используемая при отображении view.html.Методу Render значения Template может переда-ваться только одно значение, поэтому эта струк-тура будет содержать все необходимые данные.
type Guestbook struct {
	SignatureCount int      //Для хранения общего количества записей.
	Signatures     []string //Для хранения самих записей.
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)          // Запросы на просмотр списка записей будут обрабатываться функцией viewHander.
	http.HandleFunc("/guestbook/new", newHandler)       //Запросы на получение формы HTML будут обрабатываться функцией newHandler.
	http.HandleFunc("/guestbook/create", createHandler) //Запросы на отправку формы будут обрабатываться функ-цией createHandler.
	err := http.ListenAndServe("localhost:8080", nil)   // сервер запускается для прослушивания с портом 8080.
	log.Fatal(err)                                      // Ошибка никогда не равна nil, поэтому мы не вызываем для нее «check»
}

func check(err error) {
	if err != nil { //Будет вызываться, когда потребуется проверить зна-чение ошибки, возвращаемое функцией или методом.
		log.Fatal(err)
	}
}

// viewHandler читает записи гостевой книги и выводит их
// вместе со счетчиком записей.
func viewHandler(writer http.ResponseWriter, request *http.Request) { //Как обычно, функциям-обра-ботчикам передается значе-ние ResponseWriter, а также указатель на значение Request.
	signatures := getString("signatures.txt")
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(writer, guestbook)
	check(err)
}

// getStrings возвращает сегмент строк, прочитанный из fileName,
// по одной строке на каждую строку файла.
func getString(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

// newHandler отображает форму для ввода записи.
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

// createHandler получает запрос POST с добавляемой записью
// и присоединяет ее к файлу signatures.
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}
