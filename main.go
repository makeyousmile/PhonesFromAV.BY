package main

import "time"

var db DB

func main() {

	startHttpServer()
	//run(999)
	//getToken()
	//sendSms("test", []string{})
	db = Newdb()
	phones := make(chan Phone)

	go proc(phones)
	//phones := getPhones(10)
	//log.Print(phones)
	//dbInsert(db, phones)
	//db()

}

func proc(phones chan Phone) {
	//1й запуск - сбор всех номор
	go getPhone(200, phones)
	go dbInsertPhone(phones)
	for {
		time.Sleep(time.Minute)
		getPhone(5, phones)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
