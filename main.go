package main

import (
	"log"
	"time"
)

var db DB

func main() {

	//run(999)
	//getToken()
	//sendSms("test", []string{})
	db = Newdb()
	//GetPhonesCount()
	log.Print(GetSMS())
	phones := make(chan Phone)
	//
	go proc(phones)
	////phones := getPhones(10)
	////log.Print(phones)
	////dbInsert(db, phones)
	////db()
	startHttpServer()

}

func proc(phones chan Phone) {
	//1й запуск - сбор всех номор
	go getPhone(200, phones)
	go dbInsertPhone(phones)
	for {
		time.Sleep(time.Minute)
		log.Print("Get phones fom loop")
		getPhone(5, phones)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
