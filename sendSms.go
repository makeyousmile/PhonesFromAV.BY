package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type MtsSms struct {
	Messages []struct {
		Content struct {
			ShortText string `json:"short_text"`
		} `json:"content"`
		To []struct {
			Msisdn    string `json:"msisdn"`
			MessageId string `json:"message_id"`
		} `json:"to"`
	} `json:"messages"`
	Options struct {
		Class int `json:"class"`
		From  struct {
			SmsAddress string `json:"sms_address"`
		} `json:"from"`
	} `json:"options"`
}

func sendSms(text string, phones []string) {
	httpposturl := "https://api.sendpulse.com/sms/send"
	bearer := "Bearer " + getToken()
	var jsonData = []byte(`{
   "sender":"Avtodvor.by",
   "phones":[
      "375293304983",
	"375296668485"
   ],
   "body":"Выкупим ваше авто",
   "transliterate":1,
   "route":{
      "BY":"international"
   },
   "emulate":0,
   "use_dynamic_list":true,
   "date":"2023-05-03 21:08:00"
	}`)
	req, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
}
func getToken() string {
	type Token struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}
	token := Token{}
	httpposturl := "https://api.sendpulse.com/oauth/access_token"
	fmt.Println("HTTP JSON POST URL:", httpposturl)

	var jsonData = []byte(`{
   "grant_type":"client_credentials",
   "client_id":"4f1bd0384e49319b489aa60642ebdaeb",
   "client_secret":"6541189fb587f29384f13a17291c4a9a"
	}`)
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	json.Unmarshal(body, &token)
	log.Print(token.AccessToken)
	return token.AccessToken
}
