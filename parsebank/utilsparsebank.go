// parsebank
package parsebank

//	"fmt"
//	"go-board-money/pick"
//	"io/ioutil"
//	"net/http"
//	"strconv"
//	"strings"

//	"golang.org/x/net/html/charset"

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
func MaxPokupkaValuta(s []Kurs) (string, float64) {
	bank := ""
	pokupka := 0.0
	if len(s) <= 0 {
		return bank, pokupka
	}

	bank = s[0].Namebank
	pokupka = s[0].Pokupka
	for i := 1; i < len(s); i++ {
		if pokupka < s[i].Pokupka {
			pokupka = s[i].Pokupka
			bank = s[i].Namebank
		}
	}

	return bank, pokupka
}

// минимальная цена продажи банком валюты и название банка
func MinProdajaValuta(s []Kurs) (string, float64) {
	bank := ""
	prodaja := 0.0
	if len(s) <= 0 {
		return bank, prodaja
	}

	bank = s[0].Namebank
	prodaja = s[0].Prodaja
	for i := 1; i < len(s); i++ {
		if prodaja > s[i].Prodaja {
			prodaja = s[i].Prodaja
			bank = s[i].Namebank
		}
	}

	return bank, prodaja
}
