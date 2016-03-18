// genhtml
package pick

import (
	//	"fmt"
	"os"
	"strconv"
	"time"
)

func Savestrtofile(namef string, str string) int {
	file, err := os.Create(namef)
	if err != nil {
		// handle the error here
		return -1
	}
	defer file.Close()

	file.WriteString(str)
	return 0
}

// ----------------- функции генерации html page
//-- генерация ячейки таблицы в html
func tablecell(str string) string {
	return "<TD>" + str + "</TD>" + "\n"
}

//-- генерация ссылки в html
func Link(str string, url string) string {
	return "<a href=\"" + url + "\" >" + str + "</a> <br>" + "\n"
}

//-- генерация строки таблицы в html
func gentablestroka(str []string) string {
	res0 := ""
	for i := 0; i < len(str); i++ {
		res0 += tablecell(str[i])
	}
	return "<TR>" + "\n" + res0 + "</TR>" + "\n"
}

func Htmlpage(surl []string) string {
	zagol := "НОВОСТИ"
	begstr := "<html>\n <head>\n <meta charset='utf-8'>\n <title>" + zagol + "</title>\n </head>\n <body>\n"
	bodystr := ""
	for i := 0; i < len(surl); i++ {
		ch := strconv.Itoa(i)
		bodystr += Link("новость "+ch, surl[i])
	}
	//	bodystr := genhtmltable0(datas, zagol, keys)
	endstr := "</body>\n" + "</html>"
	return begstr + bodystr + endstr
}

//--------------------

// генерация html главной страницы Начало
func HtmlpageBegins() string {
	zagol := "ДОСКА ВАЛЮТ и МЕТАЛЛОВ"
	stime := "<br>" + "Выгружено: " + time.Now().String() + "<br>"
	begstr := "<html>\n <head>\n <meta charset='utf-8'>\n <title>" + zagol + "</title>\n </head>\n <body>\n" + "<h1 align=\"center\"><a name=\"MainPage\"> ДОСКА ВАЛЮТ и МЕТАЛЛОВ </a></h1>" + stime
	return begstr + "<br>"
}

// генерация html главной страницы Конец
func HtmlpageEnds() string {
	endstr := "</body>\n" + "</html>"
	return endstr
}

// шаблон оформления таблицы
func HtmlTableValuta(ss string) string {
	bodystr := "<h3 align=\"center\"></h3><br>" + "<TABLE align=\"center\" border=\"1\"><TR><TD>БАНК</TD><TD>ВАЛЮТА</TD><TD>ПОКУПКА</TD><TD>ПРОДАЖА</TD></TR>"
	//	for i := 0; i < len(sn); i++ {
	//		bodystr += "<TR> <TD width=\"350\"> <b>" + genhtml.Link(sn[i].title, sn[i].url) + "</b></TD>" + "<TD width=\"550\"><br>" + sn[i].content + "" + "<br> <a href=\"#MainPage\"> В начало </a>" + " <a href=\"#" + titlenews + "\"> К " + titlenews + " </a> " + "</TD> </TR>"
	//	}
	bodystr += ss + "</TABLE>"
	return bodystr
}

// шаблон оформления таблицы
func HtmlTableValutaOld(ss string) string {
	bodystr := "<h3 align=\"center\"></h3><br>" + "<TABLE align=\"center\" border=\"1\"><TR><TD>БАНК</TD><TD>ВАЛЮТА</TD><TD>ПОКУПКА</TD><TD>ПРОДАЖА</TD></TR>"
	//	for i := 0; i < len(sn); i++ {
	//		bodystr += "<TR> <TD width=\"350\"> <b>" + genhtml.Link(sn[i].title, sn[i].url) + "</b></TD>" + "<TD width=\"550\"><br>" + sn[i].content + "" + "<br> <a href=\"#MainPage\"> В начало </a>" + " <a href=\"#" + titlenews + "\"> К " + titlenews + " </a> " + "</TD> </TR>"
	//	}
	bodystr += ss + "</TABLE>"
	return bodystr
}
