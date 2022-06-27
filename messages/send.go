package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Send a message with 3 Buttons
func SendButton(from, to, apiKey, BotName, Title, Caption, m1, m2, m3 string) {
	params := url.Values{}
	params.Add("channel", `whatsapp`)
	params.Add("source", from)
	params.Add("destination", to)
	params.Add("message", `{"type":"quick_reply","content":{"type":"text","text":"`+Title+`","caption":"`+Caption+`"},"options":[{"type":"text","title":"`+m1+`"},{"type":"text","title":"`+m2+`"},{"type":"text","title":"`+m3+`"}]}`)
	params.Add("src.name", BotName)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://api.gupshup.io/sm/api/v1/msg", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, res, "", "\t")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(prettyJSON.String())
}

// Send a Message
func SendMessage(from, to, apiKey, BotName, Text string) {
	params := url.Values{}
	params.Add("channel", `whatsapp`)
	params.Add("source", from)
	params.Add("destination", to)
	params.Add("message", `{"type":"text","text":"`+Text+`"}`)
	params.Add("src.name", BotName)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://api.gupshup.io/sm/api/v1/msg", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, res, "", "\t")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(prettyJSON.String())
}

// Esto es para ver cuales usuarios son opt-in, o sea tienen la autorazación para que el bot les escriba.
func GetOptInUserList(AppName, ApiKey string) string {
	req, err := http.NewRequest("GET", "https://api.gupshup.io/sm/api/v1/users/"+AppName, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Apikey", ApiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	return string(res)
}

// Bueno esta no sé si sea tan necesario
func SendLocation(From, To, ApiKey, BotName, Place, Address, Lon, Lat string) string {
	params := url.Values{}
	params.Add("channel", `whatsapp`)
	params.Add("source", From)
	params.Add("destination", To)
	params.Add("message", `{"type":"location","name":"`+Place+`","address":"`+Address+`","longitude":`+Lon+`,"latitude":`+Lat+`,"caption":"Hello world"}`)
	params.Add("src.name", BotName)
	body := strings.NewReader(params.Encode())
	req, err := http.NewRequest("POST", "https://api.gupshup.io/sm/api/v1/msg", body)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	return string(res)
}

// Bueno esta no es la mejor forma de hacer esto, Además neesitas un json file
func SendList(From, To, ApiKey, BotName string, JsonFile string) string {
	params := url.Values{}
	params.Add("channel", `whatsapp`)
	params.Add("source", From)
	params.Add("destination", To)
	// read from test.json
	read, err := ioutil.ReadFile(JsonFile)
	if err != nil {
		log.Println(err)
	}
	params.Add("message", string(read))
	params.Add("src.name", BotName)
	body := strings.NewReader(params.Encode())
	req, err := http.NewRequest("POST", "https://api.gupshup.io/sm/api/v1/msg", body)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	return string(res)

}
