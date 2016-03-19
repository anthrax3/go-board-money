// parsebank
package parsebank

import (
	"fmt"
	"go-board-money/pick"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Kurs struct {
	Namebank string  // название банка
	Valuta   string  // название валюты
	Pokupka  float64 // покупка валюты (покупает банк)
	Prodaja  float64 // продажа валюты  (продает банк)
}

// если продажа или покупка равно или меньше нуля то true
func (k *Kurs) isNullKurs() bool {
	if (k.Prodaja > 0) && (k.Pokupka > 0) {
		return false
	} else {
		return true
	}
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
		//		panic("IO error")
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
		fmt.Printf("|%-7s|%-7s|%8.2f|%8.2f|\n", v.Namebank, v.Valuta, v.Pokupka, v.Prodaja)
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

	kursvaluta = append(kursvaluta, Kurs{Namebank: "SBRF", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "SBRF", Valuta: "EUR"})
	if len(stable) >= 6 {
		// USD
		kursvaluta[0].Pokupka = convstrtofloat(stable[2])
		kursvaluta[0].Prodaja = convstrtofloat(stable[3])
		// EUR
		kursvaluta[1].Pokupka = convstrtofloat(stable[4])
		kursvaluta[1].Prodaja = convstrtofloat(stable[5])
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

	kursvaluta = append(kursvaluta, Kurs{Namebank: "AKBARS", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "AKBARS", Valuta: "EUR"})
	if len(stable) >= 14 {
		//		// USD
		kursvaluta[0].Pokupka = convstrtofloat(stable[3])
		kursvaluta[0].Prodaja = convstrtofloat(stable[6])
		// EUR
		kursvaluta[1].Pokupka = convstrtofloat(stable[10])
		kursvaluta[1].Prodaja = convstrtofloat(stable[13])
	} else {
		fmt.Println("Error parse ParserAkBars ")
		fmt.Println(stable)
	}

	if (kursvaluta[0].isNullKurs()) || (kursvaluta[1].isNullKurs()) {
		fmt.Println("See parse AkBars Bank")
		//		fmt.Println(stable)
		//		// USD
		kursvaluta[0].Pokupka = convstrtofloat(stable[3])
		kursvaluta[0].Prodaja = convstrtofloat(stable[5])
		// EUR
		kursvaluta[1].Pokupka = convstrtofloat(stable[8])
		kursvaluta[1].Prodaja = convstrtofloat(stable[10])
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

	kursvaluta = append(kursvaluta, Kurs{Namebank: "TFB", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "TFB", Valuta: "EUR"})
	if (len(stable) >= 3) && (len(stable2) >= 3) {
		//		// USD
		kursvaluta[0].Pokupka = convstrtofloat(stable[1])
		kursvaluta[0].Prodaja = convstrtofloat(stable[2])
		//		// EUR
		kursvaluta[1].Pokupka = convstrtofloat(stable2[1])
		kursvaluta[1].Prodaja = convstrtofloat(stable2[2])
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

	kursvaluta = append(kursvaluta, Kurs{Namebank: "BINBANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "BINBANK", Valuta: "EUR"})
	//	if (len(stable) >= 3) && (len(stable2) >= 3) {
	//		// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[3])
	kursvaluta[0].Prodaja = convstrtofloat(stable[4])
	//		// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[6])
	kursvaluta[1].Prodaja = convstrtofloat(stable[7])
	//	} else {
	//		fmt.Println("Error parse ParserAkBars ")
	//		fmt.Println("stable = ", stable)
	//		fmt.Println("stable2 = ", stable2)
	//	}

	return kursvaluta
}

//парсер валют с Банка Казани
func ParserValutaBankkazan(url string) []Kurs {

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
		"div",
		&pick.Attr{
			"class",
			"b-aside-currency__line background_green",
		},
	})

	stable2, _ := pick.PickText(&pick.Option{
		&shtml,
		"div",
		&pick.Attr{
			"class",
			"b-aside-currency__line background_blue",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)
	stable2 = delspace(stable2)
	//	fmt.Println(stable2)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "BANKKAZAN", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "BANKKAZAN", Valuta: "EUR"})
	//	if (len(stable) >= 3) && (len(stable2) >= 3) {
	//		// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[1])
	kursvaluta[0].Prodaja = convstrtofloat(stable[2])
	//		// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable2[1])
	kursvaluta[1].Prodaja = convstrtofloat(stable2[2])
	//	} else {
	//		fmt.Println("Error parse ParserAkBars ")
	//		fmt.Println("stable = ", stable)
	//		fmt.Println("stable2 = ", stable2)
	//	}

	return kursvaluta
}

//парсер валют с РосИнтерБанк
func ParserValutaRosinterbank(url string) []Kurs {

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
		"div",
		&pick.Attr{
			"class",
			"hold",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "ROSINTERBANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "ROSINTERBANK", Valuta: "EUR"})
	//	if (len(stable) >= 3) && (len(stable2) >= 3) {
	//		// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[5])
	kursvaluta[0].Prodaja = convstrtofloat(stable[6])
	//		// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[9])
	kursvaluta[1].Prodaja = convstrtofloat(stable[10])
	//	} else {
	//		fmt.Println("Error parse ParserAkBars ")
	//		fmt.Println("stable = ", stable)
	//		fmt.Println("stable2 = ", stable2)
	//	}

	return kursvaluta
}

//парсер валют с Интехбанк
func ParserValutaIntechbank(url string) []Kurs {

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
			"course",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "INTECHBANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "INTECHBANK", Valuta: "EUR"})
	// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[3])
	kursvaluta[0].Prodaja = convstrtofloat(stable[4])
	// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[6])
	kursvaluta[1].Prodaja = convstrtofloat(stable[7])

	return kursvaluta
}

//парсер валют с Втб24  - РАЗОБРАТЬСЯ и СДЕЛАТЬ!!!!
func ParserValutaVtb24(url string) []Kurs {

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
			"course",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "INTECHBANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "INTECHBANK", Valuta: "EUR"})
	// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[3])
	kursvaluta[0].Prodaja = convstrtofloat(stable[4])
	// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[6])
	kursvaluta[1].Prodaja = convstrtofloat(stable[7])

	return kursvaluta
}

//парсер валют с HomeCredit
func ParserValutaHomecredit(url string) []Kurs {

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
			"din",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(st`able)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "HOMECREDIT", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "HOMECREDIT", Valuta: "EUR"})
	// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[3])
	kursvaluta[0].Prodaja = convstrtofloat(stable[4])
	// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[6])
	kursvaluta[1].Prodaja = convstrtofloat(stable[7])

	return kursvaluta
}

//парсер валют с Alfabank  - СДЕЛАТЬ -РАЗОБРАТЬСЯ - Курсы для операционных касс, регионы
func ParserValutaAlfabank(url string) []Kurs {

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
			"table",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "ALFABANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "ALFABANK", Valuta: "EUR"})
	// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[3])
	kursvaluta[0].Prodaja = convstrtofloat(stable[4])
	// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[6])
	kursvaluta[1].Prodaja = convstrtofloat(stable[7])

	return kursvaluta
}

//парсер валют с Akibank
func ParserValutaAkibank(url string) []Kurs {

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
		"div",
		&pick.Attr{
			"id",
			"kursy",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "AKIBANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "AKIBANK", Valuta: "EUR"})
	// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[6])
	kursvaluta[0].Prodaja = convstrtofloat(stable[7])
	// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[9])
	kursvaluta[1].Prodaja = convstrtofloat(stable[10])

	return kursvaluta
}

//парсер валют с СпуртБанк
func ParserValutaSpurtbank(url string) []Kurs {

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
			"money-info",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "SPURTBANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "SPURTBANK", Valuta: "EUR"})
	// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[3])
	kursvaluta[0].Prodaja = convstrtofloat(stable[4])
	// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[6])
	kursvaluta[1].Prodaja = convstrtofloat(stable[7])

	return kursvaluta
}

//парсер валют с Русский стандарт
func ParserValutaRusstandartbank(url string) []Kurs {

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
			"id",
			"table_cash",
		},
	})

	stable = delspace(stable)
	//	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "RUSSTANDARTBANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "RUSSTANDARTBANK", Valuta: "EUR"})
	// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[3])
	kursvaluta[0].Prodaja = convstrtofloat(stable[4])
	// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[6])
	kursvaluta[1].Prodaja = convstrtofloat(stable[7])

	return kursvaluta
}

//парсер валют с Росбанк
func ParserValutaRosbank(url string) []Kurs {

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
		"div",
		&pick.Attr{
			"class",
			"rates",
		},
	})

	stable = delspace(stable)
	fmt.Println(stable)

	kursvaluta = append(kursvaluta, Kurs{Namebank: "ROSBANK", Valuta: "USD"})
	kursvaluta = append(kursvaluta, Kurs{Namebank: "ROSBANK", Valuta: "EUR"})
	// USD
	kursvaluta[0].Pokupka = convstrtofloat(stable[3])
	kursvaluta[0].Prodaja = convstrtofloat(stable[4])
	// EUR
	kursvaluta[1].Pokupka = convstrtofloat(stable[6])
	kursvaluta[1].Prodaja = convstrtofloat(stable[7])

	return kursvaluta
}
