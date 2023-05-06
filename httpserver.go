package main

import (
	"html/template"
	"net/http"
)

type TemplateVars struct {
	AllPhonesCount  string
	AllSmsCount     string
	TodayPhoneCount string
	TodaySMScount   string
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	tmp := TemplateVars{}
	tmp.AllPhonesCount = "3000"
	tmp.AllSmsCount = "3000"
	tmp.TodayPhoneCount = "10"
	tmp.TodaySMScount = "10"
	templ, err := template.ParseFiles("html/index.html")
	checkErr(err)
	templ.Execute(w, tmp)

}
func startHttpServer() {
	http.HandleFunc("/", HttpHandler)
	http.ListenAndServe(":8080", nil)
}
