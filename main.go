package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

type ViewData struct {
	Title     string
	Message   string
	StaticDir string
}

func get_ajax(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		c := Curl{}
		form := DataForm{}
		q := req.URL.Query()
		form.Author = q.Get("AU")
		form.Id = q.Get("ID")
		form.Location = q.Get("PP")
		form.Publishing = q.Get("PU")
		form.P_Date = q.Get("PY")
		form.Title = q.Get("TI")
		form.Iddb = q.Get("iddb")
		data := c.response(form)
		bolB, _ := json.Marshal(data)
		fmt.Fprint(w, string(bolB))

	case "POST":
		fmt.Fprintln(w, "POST")
	default:
		fmt.Fprintln(w, "sorri")
	}
}
func user_r(w http.ResponseWriter, req *http.Request) {
	data := ViewData{
		StaticDir: "/static/",
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, data)
}
func admin_r(w http.ResponseWriter, r *http.Request) {
	data := ViewData{
		StaticDir: "/static/",
	}
	tmpl, _ := template.ParseFiles("templates/admin.html")
	tmpl.Execute(w, data)
}

func get_one(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var c CurlR
		var data PostResponseRUSMARC
		df := DataForm{}
		q := r.URL.Query()
		df.Title = q.Get("TI")
		data = c.response(df)
		for _, s := range data.Result[0].UNIMARC {
			if strings.HasPrefix(s, "852") == true {
				r, _ := regexp.Compile(`\D!\d*`)
				re_gr := r.FindAllString(s, 2)
				data.Result[0].Locate.Room = re_gr[0]
				data.Result[0].Locate.Stelach = re_gr[1]
			}
		}
		a, err := json.Marshal(data.Result[0])
		if err != nil {
			fmt.Fprint(w, err)
		}
		fmt.Fprint(w, string(a))
	case "POST":
		fmt.Fprintln(w, "not POST")
	default:
		fmt.Fprintln(w, "sorri")
	}
}

func main() {
	var listenaddr = "0.0.0.0:8000"
	http.HandleFunc("/", user_r)
	http.HandleFunc("/admin/", admin_r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/api/searchs", get_ajax)
	http.HandleFunc("/api/search", get_one)
	http.ListenAndServe(listenaddr, nil)
}
