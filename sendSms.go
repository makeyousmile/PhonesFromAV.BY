package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SMS struct {
	PhoneNumber    int64    `json:"phone_number"`
	ExtraId        string   `json:"extra_id"`
	CallbackUrl    string   `json:"callback_url"`
	StartTime      string   `json:"start_time"`
	Tag            string   `json:"tag"`
	Channels       []string `json:"channels"`
	ChannelOptions struct {
		Sms struct {
			Text      string `json:"text"`
			AlphaName string `json:"alpha_name"`
			Ttl       int    `json:"ttl"`
		} `json:"sms"`
	} `json:"channel_options"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
func sendSMSMTS(sms chan string) {
	for text := range sms {
		smstext := GetSMS()
		sms := SMS{
			ChannelOptions: struct {
				Sms struct {
					Text      string `json:"text"`
					AlphaName string `json:"alpha_name"`
					Ttl       int    `json:"ttl"`
				} `json:"sms"`
			}{},
		}
		sms.Channels = []string{"sms"}
		sms.ChannelOptions.Sms = struct {
			Text      string `json:"text"`
			AlphaName string `json:"alpha_name"`
			Ttl       int    `json:"ttl"`
		}(struct {
			Text      string
			AlphaName string
			Ttl       int
		}{Text: smstext, AlphaName: "", Ttl: 300})
		num, _ := strconv.ParseInt(text, 10, 64)
		sms.PhoneNumber = num
		sms.StartTime = time.Now().String()
		sms.CallbackUrl = "https://send-dr-here.com"
		sms.StartTime = time.Now().Format("2006-01-02 15:04:05")
		httpposturl := "https://api.communicator.mts.by/686/json2/simple"

		jsonData, _ := json.Marshal(sms)
		req, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
		checkErr(err)
		req.Header.Add("Authorization", "Basic "+basicAuth("autodvor.by_bxd5", "oHoHbo"))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERRO] -", err)
		} else {
			defer resp.Body.Close()
			data, _ := io.ReadAll(resp.Body)
			fmt.Println(string(data))
		}
	}

}

//func sendSms(text string, phones []string) {
//	httpposturl := "https://api.sendpulse.com/sms/send"
//	bearer := "Bearer " + getToken()
//	var jsonData = []byte(`{
//   "sender":"Avtodvor.by",
//   "phones":[
//      "375293304983",
//	"375296668485"
//   ],
//   "body":"Выкупим ваше авто",
//   "transliterate":1,
//   "route":{
//      "BY":"international"
//   },
//   "emulate":0,
//   "use_dynamic_list":true,
//   "date":"2023-05-03 21:08:00"
//	}`)
//	req, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
//	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
//	req.Header.Set("Authorization", bearer)
//
//	client := &http.Client{}
//
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Println("Error on response.\n[ERRO] -", err)
//	} else {
//		defer resp.Body.Close()
//		data, _ := io.ReadAll(resp.Body)
//		fmt.Println(string(data))
//	}
//}
//func getToken() string {
//	type Token struct {
//		AccessToken string `json:"access_token"`
//		TokenType   string `json:"token_type"`
//		ExpiresIn   int    `json:"expires_in"`
//	}
//	token := Token{}
//	httpposturl := "https://api.sendpulse.com/oauth/access_token"
//	fmt.Println("HTTP JSON POST URL:", httpposturl)
//
//	var jsonData = []byte(`{
//   "grant_type":"client_credentials",
//   "client_id":"4f1bd0384e49319b489aa60642ebdaeb",
//   "client_secret":"6541189fb587f29384f13a17291c4a9a"
//	}`)
//	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
//	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
//
//	client := &http.Client{}
//	response, error := client.Do(request)
//	if error != nil {
//		panic(error)
//	}
//	defer response.Body.Close()
//
//	fmt.Println("response Status:", response.Status)
//	fmt.Println("response Headers:", response.Header)
//	body, _ := io.ReadAll(response.Body)
//	fmt.Println("response Body:", string(body))
//	json.Unmarshal(body, &token)
//	log.Print(token.AccessToken)
//	return token.AccessToken
//}
