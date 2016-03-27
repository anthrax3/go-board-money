// go-board-money
package main

import (
	"flag"
	"fmt"
	"go-board-money/parsebank"
	"go-board-money/pick"
)

var todir string

// функция парсинга аргументов программы
func parse_args() bool {
	flag.StringVar(&todir, "todir", "", "Конечная папка для выгрузки результируюих файлов.")
	flag.Parse()
	if todir == "" {
		todir = ""
	}
	return true
}

func main() {

	parse_args()

	fmt.Println("Start parser")

	str, res := parsebank.RunGoBoardValutaHtml()
	pick.Savestrtofile(todir+"board-money.html", str)

	usdkurspokupka := res[0]
	usdkursprodaja := res[1]
	eurkurspokupka := res[2]
	eurkursprodaja := res[3]

	fmt.Print("Лучшая покупка USD: ")
	fmt.Println(usdkurspokupka)

	fmt.Print("Лучшая продажа USD: ")
	fmt.Println(usdkursprodaja)

	fmt.Print("Лучшая покупка EUR: ")
	fmt.Println(eurkurspokupka)

	fmt.Print("Лучшая продажа EUR: ")
	fmt.Println(eurkursprodaja)

	fmt.Println("End parser")

}
