package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PostRequestRUSMARC struct {
	iddb    string
	ID      string
	AU      string
	TI      string
	PY      string
	PU      string
	PP      string
	RUSMARC string
}

// Response Body via Post for json
type PostResponseRUSMARC struct {
	MaxLastResult string `json:"maxLastResult"`
	Size          string `json:"size"`
	Result        []struct {
		Isn    string `json:"isn"`
		ID     string `json:"id"`
		Level  string `json:"level"`
		Locate struct {
			Room    string `json:"room"`
			Stelach string `json:"stelach"`
		} `json:"locate"`
		Iddb         string `json:"iddb"`
		SourceIddb   string `json:"sourceIddb"`
		Archive      string `json:"archive"`
		ControlType  string `json:"controlType"`
		ResourceType string `json:"resourceType"`
		Status       string `json:"status"`
		UNIMARC      []string
	} `json:"result"`
	Iddb []struct {
		Number string
		Title  string
	} `json:iddb`
}

type CurlR struct {
	date_request  PostRequest
	date_response PostResponse
}

func (CurlR) decode(data []byte) PostResponseRUSMARC {
	var k int
	result := PostResponseRUSMARC{}
	for i := 3; i < len(data); i++ {
		if data[i] == '{' {
			k = i
			break
		}
	}
	data = data[k:]
	if err := json.Unmarshal(data, &result); err != nil {
		log.Println(err)
		return PostResponseRUSMARC{}
	}
	return result
}

func (CurlR) response(pr DataForm) PostResponseRUSMARC {
	c := CurlR{}
	path := "http://poisk.ngonb.ru/opacg.integration.smev/ajax.php"
	q := string("f")
	q = "iddb=" + pr.Iddb + "&ID=" + pr.Id + "&AU=" + pr.Author + "&TI=" + pr.Title + "&PY=" + pr.P_Date + "&PU=" + pr.Publishing + "&PP=" + pr.Location + "&RUSMARC=1"
	buf := bytes.NewBuffer([]byte(q))
	resp, err := http.Post(path, "application/x-www-form-urlencoded; charset=UTF-8", buf)
	if err != nil {
		fmt.Println(err)
		return PostResponseRUSMARC{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	return c.decode(body)
}
