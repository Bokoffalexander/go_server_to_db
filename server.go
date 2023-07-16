/*
Небольшой веб-сервер.
Получает из строки запроса значение переменной name 
и выводить это значение в окно браузера.
*/
package main

import (
	"fmt" // для вывода в браузер
	"log" // будет выводить информацию в консоль
	"net/http" // предоставляет реализации HTTP-клиента и сервера
)
var Book_id string
// запрос с строке: http://127.0.0.1:8000/?name=Sasha
// ответ в браузере: My name is Sasha.
func greetings(w http.ResponseWriter, r *http.Request) { 
	// w - интерфейс, записывает ответ(вывод) в браузер.
	// r - http запрос, который будет приходить на наш веб-сервер
	//name := r.URL.Query().Get("name") // значение из URL строчки т.е. строки запроса браузера
	printTable(w)

	fmt.Fprintf(w, "Go to /sql?book_id=1\n")
	fmt.Fprintf(w, "book_id = from 1 to 6. \n")
	
}

func sql_query(w http.ResponseWriter, r *http.Request) {
	printTable(w)
	Book_id = r.URL.Query().Get("book_id") // значение из URL строчки т.е. строки запроса браузера
	if Book_id == "" {
		Book_id = "1"
	}
	myDBconnect() //from main.go start db connect to postgresql
	
	fmt.Fprintf(w, "Initial book_id is 1.\n")
	fmt.Fprintf(w, "Current book_id is '%s'.\n", Book_id)
	sql_result := Name // global Name is from main.go
	fmt.Fprintf(w, "My SQL query of name is '%s'.\n", sql_result) // выводит значение в какой-либо интерфейс

}

func printTable(w http.ResponseWriter) {
	lineS := "---------+-----------------------+------------------+--------+--------"
	lineT := "--book_id|---------title---------|------author------|-price--|-amount-"
	line1 := "       1 | Мастер и Маргарита    | Булгаков М.А.    | 670.99 |      3 "
	line2 := "       2 | Белая гвардия         | Булгаков М.А.    | 540.50 |      5 "
	line3 := "       3 | Идиот                 | Достоевский Ф.М. | 460.00 |     10 "
	line4 := "       4 | Братья Карамазовы     | Достоевский Ф.М. | 799.01 |      3 "
	line5 := "       5 | Игрок                 | Достоевский Ф.М. | 480.50 |     10 "
	line6 := "       6 | Стихотворения и поэмы | Есенин С.А.      | 650.00 |     15 "
	lineE := "---------+-----------------------+------------------+--------+--------"
	fmt.Fprintf(w, "%s\n",lineS)
	fmt.Fprintf(w, "%s\n",lineT)
	fmt.Fprintf(w, "%s\n",line1)
	fmt.Fprintf(w, "%s\n",line2)
	fmt.Fprintf(w, "%s\n",line3)
	fmt.Fprintf(w, "%s\n",line4)
	fmt.Fprintf(w, "%s\n",line5)
	fmt.Fprintf(w, "%s\n",line6)
	fmt.Fprintf(w, "%s\n",lineE)
	
}

func server() { // компилятор ищет функцию server
	http.HandleFunc("/", greetings) // обработичик по такому адресу
	http.HandleFunc("/sql", sql_query) // обработичик по такому адресу
	log.Println("http://127.0.0.1:7007") // вывод информации в консоль
	http.ListenAndServe("127.0.0.1:7007", nil) // запускает дефолтный листенер http-сервера
}

