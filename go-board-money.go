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
	"strconv"
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
	res, _ := strconv.ParseFloat(strings.Replace(s, ",", ".", 1), 64)
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

func printarraykurs(s []Kurs) {
	fmt.Println("|BANK   |", "VALUTA|", "POKUPKA|", "PRODAJA|")
	fmt.Println("------------------------------------------------------------")
	for _, v := range s {
		fmt.Printf("|%-7s|%-7s|%8.2f|%8.2f|\n", v.namebank, v.valuta, v.pokupka, v.prodaja)
	}
	return
}

//парсер валют со сбербанка
func ParserValutaSbrf(url string) []Kurs {

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
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{namebank: "SBRF", valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{namebank: "SBRF", valuta: "EUR"})
	if len(stable) >= 6 {
		// USD
		kursvaluta[0].pokupka = convstrtofloat(stable[2])
		kursvaluta[0].prodaja = convstrtofloat(stable[3])
		// EUR
		kursvaluta[1].pokupka = convstrtofloat(stable[4])
		kursvaluta[1].prodaja = convstrtofloat(stable[5])
	} else {
		fmt.Println("Error parse ParserSbrf ")
		fmt.Println(stable)
	}

	return kursvaluta
}

//парсер валют с ак барс банка
func ParserValutaAkBars(url string) []Kurs {

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
			"tableDesc",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{namebank: "AKBARS", valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{namebank: "AKBARS", valuta: "EUR"})
	if len(stable) >= 14 {
		//		// USD
		kursvaluta[0].pokupka = convstrtofloat(stable[3])
		kursvaluta[0].prodaja = convstrtofloat(stable[5])
		// EUR
		kursvaluta[1].pokupka = convstrtofloat(stable[9])
		kursvaluta[1].prodaja = convstrtofloat(stable[11])
	} else {
		fmt.Println("Error parse ParserAkBars ")
		fmt.Println(stable)
	}

	return kursvaluta
}

//парсер валют с Татфондбанка банка
func ParserValutaTfb(url string) []Kurs {

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
		"tr",
		&pick.Attr{
			"class",
			"usd",
		},
	})

	stable2, _ := pick.PickText(&pick.Option{
		&shtml,
		"tr",
		&pick.Attr{
			"class",
			"euro",
		},
	})

	stable = delspace(stable)
	stable2 = delspace(stable2)
	//	fmt.Println(stable2)

	kursvaluta = append(kursvaluta, Kurs{namebank: "TFB", valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{namebank: "TFB", valuta: "EUR"})
	if (len(stable) >= 3) && (len(stable2) >= 3) {
		//		// USD
		kursvaluta[0].pokupka = convstrtofloat(stable[1])
		kursvaluta[0].prodaja = convstrtofloat(stable[2])
		//		// EUR
		kursvaluta[1].pokupka = convstrtofloat(stable2[1])
		kursvaluta[1].prodaja = convstrtofloat(stable2[2])
	} else {
		fmt.Println("Error parse ParserAkBars ")
		fmt.Println("stable = ", stable)
		fmt.Println("stable2 = ", stable2)
	}

	return kursvaluta
}

//парсер валют с Бинбанка
func ParserValutaBibbank(url string) []Kurs {

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
			"step4_cours",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{namebank: "BINBANK", valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{namebank: "BINBANK", valuta: "EUR"})
	//	if (len(stable) >= 3) && (len(stable2) >= 3) {
	//		// USD
	kursvaluta[0].pokupka = convstrtofloat(stable[3])
	kursvaluta[0].prodaja = convstrtofloat(stable[4])
	//		// EUR
	kursvaluta[1].pokupka = convstrtofloat(stable[6])
	kursvaluta[1].prodaja = convstrtofloat(stable[7])
	//	} else {
	//		fmt.Println("Error parse ParserAkBars ")
	//		fmt.Println("stable = ", stable)
	//		fmt.Println("stable2 = ", stable2)
	//	}

	return kursvaluta
}

func main() {
	//	var vkurs Kurs
	linksbanks := make(map[string]string, 0)
	linksbanks["SBRF"] = "http://data.sberbank.ru/tatarstan/ru/quotes/currencies/?base=beta"
	linksbanks["TFB"] = "http://tfb.ru/"
	linksbanks["AKBARS"] = "https://www.akbars.ru/"
	linksbanks["BINBANK"] = "http://www.binbank.ru/"

	fmt.Println(linksbanks)

	board_valuta := make([]Kurs, 0)

	fmt.Println("Start parser")

	vkurs := ParserValutaSbrf(linksbanks["SBRF"])
	board_valuta = append(board_valuta, vkurs[0])
	board_valuta = append(board_valuta, vkurs[1])

	vkurs = ParserValutaAkBars(linksbanks["AKBARS"])
	board_valuta = append(board_valuta, vkurs[0])
	board_valuta = append(board_valuta, vkurs[1])

	vkurs = ParserValutaTfb(linksbanks["TFB"])
	board_valuta = append(board_valuta, vkurs[0])
	board_valuta = append(board_valuta, vkurs[1])

	vkurs = ParserValutaBibbank(linksbanks["BINBANK"])
	board_valuta = append(board_valuta, vkurs[0])
	board_valuta = append(board_valuta, vkurs[1])

	//	printarraykurs(board_valuta)

	ss := ""
	for _, v := range board_valuta {
		ss += "<TR>" + "<TD> " + v.namebank + "</TD>" + "<TD> " + v.valuta + "</TD>" + "<TD> " + strconv.FormatFloat(v.pokupka, 'f', 2, 32) + "</TD>" + "<TD> " + strconv.FormatFloat(v.prodaja, 'f', 2, 32) + "</TD>" + "</TR>"
	}

	str := pick.HtmlpageBegins() + pick.HtmlTableValuta(ss) + pick.HtmlpageEnds()
	pick.Savestrtofile("board-money.html", str)

	fmt.Println("End parser")

}
