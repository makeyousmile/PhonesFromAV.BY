package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateVars struct {
	AllPhonesCount  string
	AllSmsCount     string
	TodayPhoneCount string
	TodaySMScount   string
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	tmp := TemplateVars{}
	tmp.AllPhonesCount = "3000"
	tmp.AllSmsCount = "3000"
	tmp.TodayPhoneCount = "10"
	tmp.TodaySMScount = "10"
	templ, err := template.ParseFiles("html/index.html")
	checkErr(err)
	templ.Execute(w, tmp)

}
func Ajax(w http.ResponseWriter, r *http.Request) {
	log.Print("post")
	if r.Method == "POST" {
		sendedData := r.FormValue("sendedData")
		fmt.Println("My request is: ", sendedData)
	}
}
func startHttpServer() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/post", Ajax)
	http.ListenAndServe(":80", nil)
}
