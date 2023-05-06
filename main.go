package main

var db DB

func main() {
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

	startHttpServer()
}

func proc(phones chan Phone) {
	go getPhone(10, phones)
	go dbInsertPhone(phones)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
