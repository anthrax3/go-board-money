// parsebank
package parsebank

import (
	//	"fmt"
	"go-board-money/pick"
	//	"io/ioutil"
	//	"net/http"
	//	"strconv"
	//	"strings"

	//	"golang.org/x/net/html/charset"
)

//основная функция которая озвращает html код резльтата и лучшие цены по валюте
func RunGoBoardValutaHtml() (string, []Kurs) {
	linksbanks := Initlinksbank()
	board_Valuta := GetBoardValuta(linksbanks)
	res := GetBestPriceValuta(board_Valuta)
	str := GetHtmlBoardValuta(board_Valuta, linksbanks, res)
	return str, res
}

//инциализация ссылок на банки
func Initlinksbank() map[string]string {
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
	return linksbanks
}

// получить доску валют из известных банков
func GetBoardValuta(linksbanks map[string]string) []Kurs {
	board_Valuta := make([]Kurs, 0)
	for i, vk := range linksbanks {
		var vkurs []Kurs
		switch i {
		case "SBRF":
			vkurs = ParserValutaSbrf(vk)
		case "TFB":
			vkurs = ParserValutaTfb(vk)
		case "AKBARS":
			vkurs = ParserValutaAkBars(vk)
		case "BINBANK":
			vkurs = ParserValutaBibbank(vk)
		case "BANKKAZAN":
			vkurs = ParserValutaBankkazan(vk)
		case "ROSINTERBANK":
			vkurs = ParserValutaRosinterbank(vk)
		case "INTECHBANK":
			vkurs = ParserValutaIntechbank(vk)
		case "VTB24":
			//			vkurs = ParserValutaVtb24(vk)
			vkurs = make([]Kurs, 2)
		case "HOMECREDIT":
			vkurs = ParserValutaHomecredit(vk)
		case "ALFABANK":
			//			vkurs = ParserValutaAlfabank(vk)
			vkurs = make([]Kurs, 2)
		case "AKIBANK":
			vkurs = ParserValutaAkibank(vk)
		case "SPURTBANK":
			vkurs = ParserValutaSpurtbank(vk)
		case "RUSSTANDARTBANK":
			vkurs = ParserValutaRusstandartbank(vk)
		case "ROSBANK":
			vkurs = ParserValutaRosbank(vk)
		default:
			vkurs = make([]Kurs, 2)
		}

		board_Valuta = append(board_Valuta, vkurs[0])
		board_Valuta = append(board_Valuta, vkurs[1])
	}

	return board_Valuta

}

//опредление лучших цен на валюты
func GetBestPriceValuta(board_Valuta []Kurs) []Kurs {

	res := make([]Kurs, 0)

	//---опредление лучших цен на валюты
	board_usd := FilterValuta(board_Valuta, "USD")
	usdkurspokupka := MaxPokupkaValuta(board_usd)
	//	fmt.Print("Лучшая покупка USD: ")
	//	fmt.Println(usdkurspokupka)

	usdkursprodaja := MinProdajaValuta(board_usd)
	//	fmt.Print("Лучшая продажа USD: ")
	//	fmt.Println(usdkursprodaja)

	board_eur := FilterValuta(board_Valuta, "EUR")
	eurkurspokupka := MaxPokupkaValuta(board_eur)
	//	fmt.Print("Лучшая покупка EUR: ")
	//	fmt.Println(eurkurspokupka)

	eurkursprodaja := MinProdajaValuta(board_eur)
	//	fmt.Print("Лучшая продажа EUR: ")
	//	fmt.Println(eurkursprodaja)
	//---END опредление лучших цен на валюты
	res = append(res, usdkurspokupka)
	res = append(res, usdkursprodaja)
	res = append(res, eurkurspokupka)
	res = append(res, eurkursprodaja)
	return res
}

//получение html кода (страницы) доски валют
func GetHtmlBoardValuta(board_Valuta []Kurs, linksbanks map[string]string, res []Kurs) string {

	usdkurspokupka := res[0]
	usdkursprodaja := res[1]
	eurkurspokupka := res[2]
	eurkursprodaja := res[3]

	ss := GenTableKursValuta(board_Valuta, linksbanks, usdkurspokupka, usdkursprodaja, eurkurspokupka, eurkursprodaja)
	str := pick.HtmlpageBegins() + pick.HtmlTableValuta(ss) + pick.HtmlpageEnds()
	return str
}
