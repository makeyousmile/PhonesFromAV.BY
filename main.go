package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"log"
	"os"
	"time"
)

type Cfg struct {
	region  string
	city    string
	timeout time.Duration
	depth   int
}

var (
	db  DB
	cfg = &Cfg{}
)

func init() {
	db = Newdb()
	cfg.region = GetRegions()
	cfg.timeout = time.Minute
	cfg.depth = 10
}
func main() {

	//logToFile()
	log.Print(GetCities())

	//sms := make(chan string)
	//
	//go sendSMSMTS(sms)
	//sms <- "+375296668485"
	//go proc()

	go systray.Run(onReady, onExit)
	startHttpServer()

}

func proc() {
	phones := make(chan Phone)
	sms := make(chan string)

	go sendSMSMTS(sms)

	//1й запуск - сбор всех номор
	go getPhone(cfg.depth, phones)
	go dbInsertPhone(phones, sms)
	for {
		time.Sleep(cfg.timeout)
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
