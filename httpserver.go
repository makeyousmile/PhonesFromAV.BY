package main

import (
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

func httpHandler(w http.ResponseWriter, _ *http.Request) {
	tmp := TemplateVars{}
	tmp.AllPhonesCount = GetPhonesCount()
	tmp.AllSmsCount = "-"
	tmp.TodayPhoneCount = GetTodayPhonesCount()
	tmp.TodaySMScount = "-"
	tmp.SMS = GetSMS()
	templ, err := template.ParseFiles("html/index.html")
	checkErr(err)
	err = templ.Execute(w, tmp)
	checkErr(err)

}
func Ajax(_ http.ResponseWriter, r *http.Request) {
	log.Print("post")
	if r.Method == "POST" {
		SMS := r.FormValue("SMS")
		if SMS != "" {
			SetSMS(SMS)
		}
		baned := r.FormValue("baned")
		if baned != "" {
			SetBaned(baned)
		}
	}
}
func AjaxCount(w http.ResponseWriter, r *http.Request) {

	count := GetPhonesCount()
	if r.Method == "POST" {
		log.Print("post")
		//w.Write(byte(count))
		log.Print(count)

	}
	_, err := io.WriteString(w, count)
	checkErr(err)
}

func AjaxCountToday(w http.ResponseWriter, r *http.Request) {

	count := GetTodayPhonesCount()
	if r.Method == "POST" {
		log.Print("post")
		//w.Write(byte(count))
		log.Print(count)

	}
	_, err := io.WriteString(w, count)
	checkErr(err)
}

func AjaxBaned(w http.ResponseWriter, r *http.Request) {
	resp := ""
	baned := GetBaned()
	if r.Method == "POST" {
		log.Print("post")
		//w.Write(byte(count))

	}
	for _, phone := range baned {
		resp += "<p>" + phone + "</p>"
	}
	_, err := io.WriteString(w, resp)
	checkErr(err)
}

func AjaxCleardb(_ http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		cleardb := r.FormValue("cleardb")
		if cleardb == "1" {
			Cleardb()
		}
	}

}
func AjaxRegionSet(_ http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		region := r.FormValue("reg")
		city := r.FormValue("city")
		log.Print(city)
		SetRegions(region, city)
	}

}
func AjaxRegionGet(w http.ResponseWriter, r *http.Request) {
	filter := cfg.region

	_, err := io.WriteString(w, filter)
	checkErr(err)
}

func AjaxCitiesGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cities := GetCities()
		_, err := io.WriteString(w, cities)
		checkErr(err)
	}
	if r.Method == "POST" {
		data := r.FormValue("cfg")
		log.Print(data)
		fulljson := GetCitiesWithRegion(data)
		_, err := io.WriteString(w, fulljson)
		log.Print(fulljson)
		checkErr(err)
	}
}

func startHttpServer() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/post", Ajax)
	http.HandleFunc("/count", AjaxCount)
	http.HandleFunc("/counttoday", AjaxCountToday)
	http.HandleFunc("/baned", AjaxBaned)
	http.HandleFunc("/cleardb", AjaxCleardb)
	http.HandleFunc("/setregion", AjaxRegionSet)
	http.HandleFunc("/getregion", AjaxRegionGet)
	http.HandleFunc("/getcities", AjaxCitiesGet)

	err := http.ListenAndServe(":80", nil)
	checkErr(err)
}
