// go-board-money
package main

import (
	//	"flag"
	"fmt"
	"go-board-money/pick"
	//	"go-bot-news/pkg/html"
	//	"io"
	"io/ioutil"
	//	"log"
	"net/http"
	//	"os"
	//	"strconv"
	"strings"
	//	"time"

	"golang.org/x/net/html/charset"
)

type Kurs struct {
	namebank string  // название банка
	valuta   string  // название валюты
	pokupka  float64 // покупка валюты (покупает банк)
	prodaja  float64 // продажа валюты  (продает банк)
}

//удаление пустых или строк в которых только пробелы
func delspace(ss []string) []string {
	res := make([]string, 0)
	for _, s := range ss {
		if strings.TrimSpace(s) != "" {
			res = append(res, strings.TrimSpace(s))
		}
	}
	return res
}

func convstrtofloat(s string) float64 {
	res := 0.0
	return res
}

//получение страницы из урла url
func gethtmlpage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP error:", err)
		panic("HTTP error")
	}
	defer resp.Body.Close()
	// вот здесь и начинается самое интересное
	utf8, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Encoding error:", err)
		panic("Encoding error")
	}
	body, err := ioutil.ReadAll(utf8)
	if err != nil {
		fmt.Println("IO error:", err)
		panic("IO error")
	}
	return body
}

// вывод на печать массива строк
func printarray(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	return
}

//парсер валют со сбербанка
func ParserSbrf(url string) []Kurs {

	kursvaluta := make([]Kurs, 0)

	if url == "" {
		return kursvaluta
	}

	body := gethtmlpage(url)
	shtml := string(body)
	//	fmt.Println(shtml)

	// выделяем данные из таблицы
	stable, _ := pick.PickText(&pick.Option{
		&shtml,
		"table",
		&pick.Attr{
			"class",
			"table3_eggs4",
		},
	})

	stable = delspace(stable)
	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{namebank: "SBRF", valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{namebank: "SBRF", valuta: "EUR"})

	ss := stable[2]
	ss[strings.Index(ss, ",")] = "."
	fmt.Println(ss)
	//	fmt.Println(fpr)
	kursvaluta[0].pokupka = 0 //fpr

	return kursvaluta
}

func main() {
	//	var vkurs Kurs

	fmt.Println("Start parser")

	vkurs := ParserSbrf("http://data.sberbank.ru/tatarstan/ru/quotes/currencies/?base=beta")
	fmt.Println(vkurs)

	fmt.Println("End parser")

}
