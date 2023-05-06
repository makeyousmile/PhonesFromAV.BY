package main

var db DB

func main() {
	//run(999)
	//getToken()
	//sendSms("test", []string{})
	db = Newdb()
	run()
	//phones := getPhones(10)
	//log.Print(phones)
	//dbInsert(db, phones)
	//db()
	startHttpServer()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
