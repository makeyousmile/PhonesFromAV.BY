package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"io/ioutil"
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
	onExit := func() {
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
	}

	go systray.Run(onReady, onExit)
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

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Autodvor SMS Bot")
	systray.SetTooltip("+375 29 666 8485")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)
}

func onExit() {
	// clean up here
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
