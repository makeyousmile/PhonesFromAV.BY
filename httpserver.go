package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type TemplateVars struct {
	AllPhonesCount  string
	AllSmsCount     string
	TodayPhoneCount string
	TodaySMScount   string
	SMS             string
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	tmp := TemplateVars{}
	tmp.AllPhonesCount = GetPhonesCount()
	tmp.AllSmsCount = "-"
	tmp.TodayPhoneCount = GetTodayPhonesCount()
	tmp.TodaySMScount = "-"
	tmp.SMS = GetSMS()
	templ, err := template.ParseFiles("html/index.html")
	checkErr(err)
	templ.Execute(w, tmp)

}
func Ajax(w http.ResponseWriter, r *http.Request) {
	log.Print("post")
	if r.Method == "POST" {
		SMS := r.FormValue("SMS")
		fmt.Println("My request is: ", SMS)
		SetSMS(SMS)
	}
}
func AjaxCount(w http.ResponseWriter, r *http.Request) {

	count := GetPhonesCount()
	if r.Method == "POST" {
		log.Print("post")
		//w.Write(byte(count))
		log.Print(count)

	}
	io.WriteString(w, count)
}

func AjaxCountToday(w http.ResponseWriter, r *http.Request) {

	count := GetTodayPhonesCount()
	if r.Method == "POST" {
		log.Print("post")
		//w.Write(byte(count))
		log.Print(count)

	}
	io.WriteString(w, count)
}

func startHttpServer() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/post", Ajax)
	http.HandleFunc("/count", AjaxCount)
	http.HandleFunc("/counttoday", AjaxCountToday)
	http.ListenAndServe(":80", nil)
}
