// go-board-money
package main

import (
	"fmt"
	"go-board-money/parsebank"
	"go-board-money/pick"
	//	"io/ioutil"
	//	"net/http"
	"strconv"
	//	"strings"

	//	"golang.org/x/net/html/charset"
)

func main() {
	//	var vkurs parsebank.parsebank.Kurs
	linksbanks := make(map[string]string, 0)
	linksbanks["SBRF"] = "http://data.sberbank.ru/tatarstan/ru/quotes/currencies/?base=beta"
	linksbanks["TFB"] = "http://tfb.ru/"
	linksbanks["AKBARS"] = "https://www.akbars.ru/"
	linksbanks["BINBANK"] = "http://www.binbank.ru/"
	linksbanks["BANKKAZAN"] = "http://www.bankofkazan.ru/"
	linksbanks["ROSINTERBANK"] = "http://kazan.rosinterbank.ru/"
	linksbanks["INTECHBANK"] = "http://www.intechbank.ru/"
	linksbanks["VTB24"] = "http://www.vtb24.ru/"
	linksbanks["HOMECREDIT"] = "http://kazan.homecredit.ru/?my_reg_id=46"
	linksbanks["ALFABANK"] = "https://alfabank.ru/kazan/currency/"
	linksbanks["AKIBANK"] = "http://www.akibank.ru/"
	linksbanks["SPURTBANK"] = "http://www.spurtbank.ru/"
	linksbanks["RUSSTANDARTBANK"] = "http://www.rsb.ru/courses/"
	linksbanks["ROSBANK"] = "http://www.rosbank.ru/ru/"

	//	fmt.Println(linksbanks)

	board_Valuta := make([]parsebank.Kurs, 0)

	fmt.Println("Start parser")

	vkurs := parsebank.ParserValutaSbrf(linksbanks["SBRF"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaAkBars(linksbanks["AKBARS"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaTfb(linksbanks["TFB"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaBibbank(linksbanks["BINBANK"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaBankkazan(linksbanks["BANKKAZAN"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaRosinterbank(linksbanks["ROSINTERBANK"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaIntechbank(linksbanks["INTECHBANK"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaHomecredit(linksbanks["HOMECREDIT"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaAkibank(linksbanks["AKIBANK"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaSpurtbank(linksbanks["SPURTBANK"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaRusstandartbank(linksbanks["RUSSTANDARTBANK"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	vkurs = parsebank.ParserValutaRosbank(linksbanks["ROSBANK"])
	board_Valuta = append(board_Valuta, vkurs[0])
	board_Valuta = append(board_Valuta, vkurs[1])

	//	vkurs = parsebank.ParserValutaAlfabank(linksbanks["ALFABANK"])
	//	board_Valuta = append(board_Valuta, vkurs[0])
	//	board_Valuta = append(board_Valuta, vkurs[1])

	//	vkurs = ParserValutaVtb24(linksbanks["VTB24"])
	//	board_Valuta = append(board_Valuta, vkurs[0])
	//	board_Valuta = append(board_Valuta, vkurs[1])

	//	printarraykurs(board_Valuta)

	board_usd := parsebank.FilterValuta(board_Valuta, "USD")
	usdbank, usdpokupka := parsebank.MaxPokupkaValuta(board_usd)
	fmt.Print("Лучшая покупка USD: ")
	fmt.Println(usdbank, usdpokupka)

	usdbank2, usdprodaja := parsebank.MinProdajaValuta(board_usd)
	fmt.Print("Лучшая продажа USD: ")
	fmt.Println(usdbank2, usdprodaja)

	board_eur := parsebank.FilterValuta(board_Valuta, "EUR")
	eurbank, eurpokupka := parsebank.MaxPokupkaValuta(board_eur)
	fmt.Print("Лучшая покупка EUR: ")
	fmt.Println(eurbank, eurpokupka)

	eurbank2, eurprodaja := parsebank.MinProdajaValuta(board_eur)
	fmt.Print("Лучшая продажа EUR: ")
	fmt.Println(eurbank2, eurprodaja)

	ss := "<TR><TD colspan='4' align='center'>USD</TD></TR>"
	//USD
	for _, v := range board_Valuta {
		if v.Valuta == "USD" {
			if usdpokupka == v.Pokupka {
				ss += "<TR>" + "<TD> <A href ='" + linksbanks[v.Namebank] + "'>" + v.Namebank + "</a></TD>" + "<TD> " + v.Valuta + "</TD>" + "<TD bgcolor='#008000'> " + strconv.FormatFloat(v.Pokupka, 'f', 2, 32) + "</TD>"
			} else {
				ss += "<TR>" + "<TD> <A href ='" + linksbanks[v.Namebank] + "'>" + v.Namebank + "</A></TD>" + "<TD> " + v.Valuta + "</TD>" + "<TD> " + strconv.FormatFloat(v.Pokupka, 'f', 2, 32) + "</TD>"
			}
			if usdprodaja == v.Prodaja {
				ss += "<TD bgcolor='#008000'> " + strconv.FormatFloat(v.Prodaja, 'f', 2, 32) + "</TD>" + "</TR>"
			} else {
				ss += "<TD> " + strconv.FormatFloat(v.Prodaja, 'f', 2, 32) + "</TD>" + "</TR>"
			}

		}
	}
	//EUR
	ss += "<TR><TD colspan='4' align='center'>EUR</TD></TR>"
	//USD
	for _, v := range board_Valuta {
		if v.Valuta == "EUR" {
			if eurpokupka == v.Pokupka {
				ss += "<TR>" + "<TD> <A href ='" + linksbanks[v.Namebank] + "'>" + v.Namebank + "</A></TD>" + "<TD> " + v.Valuta + "</TD>" + "<TD bgcolor='#008000'> " + strconv.FormatFloat(v.Pokupka, 'f', 2, 32) + "</TD>"
			} else {
				ss += "<TR>" + "<TD> <A href ='" + linksbanks[v.Namebank] + "'>" + v.Namebank + "</A></TD>" + "<TD> " + v.Valuta + "</TD>" + "<TD> " + strconv.FormatFloat(v.Pokupka, 'f', 2, 32) + "</TD>"
			}
			if eurprodaja == v.Prodaja {
				ss += "<TD bgcolor='#008000'> " + strconv.FormatFloat(v.Prodaja, 'f', 2, 32) + "</TD>" + "</TR>"
			} else {
				ss += "<TD> " + strconv.FormatFloat(v.Prodaja, 'f', 2, 32) + "</TD>" + "</TR>"
			}
		}
	}

	str := pick.HtmlpageBegins() + pick.HtmlTableValuta(ss) + pick.HtmlpageEnds()
	pick.Savestrtofile("board-money.html", str)

	fmt.Println("End parser")

}
