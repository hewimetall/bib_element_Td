package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DataForm struct {
	Iddb       string
	Id         string
	Author     string `json:"author"` //Автор
	Title      string //Заглавие
	P_Date     string //Год
	Publishing string //Издательство
	Location   string `json:"location"` //Место
	LINK       string //RUSMARC
	Adress     string
}

type PostRequest struct {
	iddb string
	ID   string
	AU   string //Автор
	TI   string //Заглавие
	PY   string //Год
	PU   string //Издательство
	PP   string //Место
}

// Response Body via Post for json
type PostResponse struct {
	MaxLastResult string     `json:"maxLastResult"`
	Size          string     `json:"size"`
	Result        []DataForm `json:"result"`
	Iddb          []struct {
		Number string `json:"number"`
		Title  string `json:"title"`
	} `json:"iddb"`
}

type Curl struct {
	date_request  PostRequest
	date_response PostResponse
}

func (Curl) decode(data []byte) PostResponse {
	var k int
	result := PostResponse{}
	for i := 3; i < len(data); i++ {
		if data[i] == '{' {
			k = i
			break
		}
	}
	data = data[k:]
	data_t := string(data)
	fmt.Println(data_t)
	if err := json.Unmarshal(data, &result); err != nil {
		log.Println(err)
		return PostResponse{}
	}
	return result
}

func (Curl) response(pr DataForm) PostResponse {
	c := Curl{}
	path := "http://poisk.ngonb.ru/opacg.integration.smev/ajax.php"
	q := string("f")
	q = "iddb=" + pr.Iddb + "&ID=" + pr.Id + "&AU=" + pr.Author + "&TI=" + pr.Title + "&PY=" + pr.P_Date + "&PU=" + pr.Publishing + "&PP=" + pr.Location
	buf := bytes.NewBuffer([]byte(q))
	resp, err := http.Post(path, "application/x-www-form-urlencoded; charset=UTF-8", buf)
	if err != nil {
		fmt.Println(err)
		return PostResponse{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	return c.decode(body)
}

// Not work
func (Curl) response_json() {
	path := "http://poisk.ngonb.ru/opacg.integration.smev/STORAGE/opacfindd/FindView/2.3.0"
	buf := bytes.NewBuffer([]byte(`"iddb":"","ID":"","AU":"","TI":"","PY":"1998","PU":"","PP":""`))
	resp, err := http.Post(path, "application/json", buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Println(string(body))

	println(resp)
}
