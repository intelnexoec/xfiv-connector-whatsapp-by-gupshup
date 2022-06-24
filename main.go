package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
)

/*
	{
	  "app": "paij0setest",
	  "timestamp": 1656028014012,
	  "version": 2,
	  "type": "message",
	  "payload": {
	    "id": "ABEGVzIpEGVAAgo-sJLQUSvAaAOC",
	    "source": "5732xxxxxxxx",
	    "type": "text",
	    "payload": { "text": "ehola" },
	    "sender": {
	      "phone": "57322xxxxxxxx",
	      "name": "Jose",
	      "country_code": "57",
	      "dial_code": "322xxxxxxxx"
	    }
	  }
	}

*/
type Message struct {
	App       string `json:"app"`
	Timestamp int64  `json:"timestamp"`
	Version   int    `json:"version"`
	Type      string `json:"type"`
	Payload   struct {
		ID      string `json:"id"`
		Source  string `json:"source"`
		Type    string `json:"type"`
		Payload struct {
			Text string `json:"text"`
		} `json:"payload"`
		Sender struct {
			Phone       string `json:"phone"`
			Name        string `json:"name"`
			CountryCode string `json:"country_code"`
			DialCode    string `json:"dial_code"`
		} `json:"sender"`
	} `json:"payload"`
}

func Post(e echo.Context) error {
	reqBody, err := ioutil.ReadAll(e.Request().Body)
	if err != nil {
		log.Println(err)
	}
	var message Message
	err = json.Unmarshal(reqBody, &message)
	if err != nil {
		log.Println(err)
	}
	name := message.Payload.Sender.Name
	text := message.Payload.Payload.Text
	phone := message.Payload.Sender.Phone
	fmt.Println("Name:"+name, "Mensaje:"+text, "Teléfono:"+phone)
	fmt.Println(message)
	template := `Hola *` + name + `* `
	return e.String(200, template)
}
func main() {
	e := echo.New()
	e.POST("/", Post)
	e.Logger.Fatal(e.Start(":8000"))
}
