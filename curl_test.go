package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestAjax(t *testing.T) {
	var c Curl
	// var data PostResponse
	df := DataForm{
		Iddb:       "",
		Id:         "",
		Author:     "",     //Автор
		Title:      "",     //Заглавие
		P_Date:     "1998", //Год
		Publishing: "",     //Издательство
		Location:   "",     //Место
		LINK:       "",
		Adress:     "",
	}
	c.response(df)
	c.response_json()
}

func TestAjaxR(t *testing.T) {
	var c CurlR
	var data PostResponseRUSMARC
	df := DataForm{
		Iddb:       "",
		Id:         `RU NOVOSIBIRSK\BIBL\2000101178`,
		Author:     "", //Автор
		Title:      "", //Заглавие
		P_Date:     "", //Год
		Publishing: "", //Издательство
		Location:   "", //Место
		LINK:       "", //RUSMARC
		Adress:     "",
	}
	data = c.response(df)
	// log.Println(data.Result[0].Archive)
	// log.Println(data.Result[0].ControlType)
	// log.Println(data.Result[0].ID)
	// log.Println(data.Result[0].Iddb)
	// log.Println(data.Result[0].Isn)
	// log.Println(data.Result[0].Level)
	// log.Println(data.Result[0].ResourceType)
	// log.Println(data.Result[0].SourceIddb)
	// log.Println(data.Result[0].Status)
	// log.Println(data.Result[0].UNIMARC)
	for _, s := range data.Result[0].UNIMARC {
		// log.Println(s)
		if strings.HasPrefix(s, "852") == true {
			r, _ := regexp.Compile(`\D!\d*`)
			re_gr := r.FindAllString(s, 2)
		}
	}

}
