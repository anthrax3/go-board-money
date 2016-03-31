// parsebank
package parsebank

import (
	//	"fmt"
	//	"go-board-money/pick"
	//	"io/ioutil"
	//	"net/http"
	"strconv"
	//	"strings"

	//	"golang.org/x/net/html/charset"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

// возвращает массив курсов удовлетворяющих фильтру по валюте svaluta
func FilterValuta(s []Kurs, svaluta string) []Kurs {
	res := make([]Kurs, 0)
	for _, v := range s {
		if v.Valuta == svaluta {
			res = append(res, v)
		}
	}

	return res
}

// максимальная цена покупки банком валюты и название банка
func MaxPokupkaValuta(s []Kurs) Kurs {
	var res Kurs
	bank := ""
	pokupka := 0.0
	if len(s) <= 0 {
		return res
	}

	bank = s[0].Namebank
	pokupka = s[0].Pokupka
	for i := 1; i < len(s); i++ {
		if pokupka < s[i].Pokupka {
			pokupka = s[i].Pokupka
			bank = s[i].Namebank
		}
	}
	res.Namebank = bank
	res.Pokupka = pokupka
	return res
}

// минимальная цена продажи банком валюты и название банка
func MinProdajaValuta(s []Kurs) Kurs {
	var res Kurs
	bank := ""
	prodaja := 0.0
	if len(s) <= 0 {
		return res
	}

	bank = s[0].Namebank
	prodaja = s[0].Prodaja
	for i := 1; i < len(s); i++ {
		if prodaja > s[i].Prodaja {
			prodaja = s[i].Prodaja
			bank = s[i].Namebank
		}
	}
	res.Namebank = bank
	res.Prodaja = prodaja
	return res
}

func GenTableKursValuta(board_Valuta []Kurs, linksbanks map[string]string, usdkurspokupka, usdkursprodaja, eurkurspokupka, eurkursprodaja Kurs) string {
	usdpokupka := usdkurspokupka.Pokupka
	usdprodaja := usdkursprodaja.Prodaja
	eurpokupka := eurkurspokupka.Pokupka
	eurprodaja := eurkursprodaja.Prodaja
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
	return ss
}
