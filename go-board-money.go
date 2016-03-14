// go-board-money
package main

import (
	//	"flag"
	"fmt"
	//	"go-bot-news/pkg"
	//	"go-bot-news/pkg/html"
	//	"io"
	"io/ioutil"
	//	"log"
	"net/http"
	//	"os"
	//	"strings"
	//	"time"

	"golang.org/x/net/html/charset"
)

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

//парсер новостей с сайта Яндекса
//func (this *News) ParserNewsYandex() {

//	if this.url == "" {
//		return
//	}

//	body := gethtmlpage(this.url)
//	shtml := string(body)

//	//<h1 class="story__head">Блаттера и Платини отстранили от футбола на 8 лет</h1>

//	stitle, _ := pick.PickText(&pick.Option{
//		&shtml,
//		"h1",
//		&pick.Attr{
//			"class",
//			"story__head",
//		},
//	})

//	if len(stitle) > 0 {
//		this.title = stitle[0]
//		//	<meta name="og:description" content="«У Турции есть полное право проводить антитеррористические операции в Сирии и других странах, где базируются террористические группировки, так как это часть борьбы против стоящих перед нами угроз», — сказал Эрдоган."/>
//		scont, _ := pick.PickAttr(&pick.Option{&shtml, "meta", &pick.Attr{"name", "og:description"}}, "content")
//		this.content = scont[0]
//	}

//	return
//}

func main() {
	fmt.Println("Hello World!")
}
