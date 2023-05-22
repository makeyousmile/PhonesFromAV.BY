package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"log"
	"os"
	"time"
)

var db DB

func init() {
	db = Newdb()
}
func main() {

	logToFile()

	sms := make(chan string)

	go sendSMSMTS(sms)
	sms <- "+375293304983"
	//go proc()
	//addMessageId("296668485", "test")
	go systray.Run(onReady, onExit)
	startHttpServer()

}

func proc() {
	phones := make(chan Phone)
	sms := make(chan string)

	go sendSMSMTS(sms)

	//1й запуск - сбор всех номор
	go getPhone(200, phones)
	go dbInsertPhone(phones, sms)
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
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
		os.Exit(0)

	}()
}

func onExit() {
	// clean up here
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
