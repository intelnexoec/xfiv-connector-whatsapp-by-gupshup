package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
)

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
			Text    string `json:"text"`
			Url     string `json:"url"`
			Caption string `json:"caption"`
		} `json:"payload"`
		Sender struct {
			Phone       string `json:"phone"`
			Name        string `json:"name"`
			CountryCode string `json:"country_code"`
			DialCode    string `json:"dial_code"`
		} `json:"sender"`
	} `json:"payload"`
}

func ValidateMessage(e echo.Context, message string) error {
	events := []string{"delivered", "seen", "enqueued"}
	for _, event := range events {
		if strings.Contains(message, event) {
			log.Println(color.Blue("[INFO]"), "Message event:", event)
		}
	}
	switch message {
	case "text":
		log.Println(color.Green("Texto"))
	case "image":
		log.Println(color.Green("Imagen"))
	case "audio":
		log.Println(color.Green("Audio"))
	case "video":
		log.Println(color.Green("Video"))
	case "contact":
		log.Println(color.Green("Contacto"))
	case "sticker":
		log.Println(color.Green("Sticker"))
	case "location":
		log.Println(color.Green("Ubicación"))
	}
	return nil
}

func Post(e echo.Context) error {
	reqBody, err := ioutil.ReadAll(e.Request().Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(reqBody))
	var message Message
	err = json.Unmarshal(reqBody, &message)
	if err != nil {
		log.Println(err)
	}
	name := message.Payload.Sender.Name
	text := message.Payload.Payload.Text
	caption := message.Payload.Payload.Caption
	phone := message.Payload.Sender.Phone
	fileUrl := message.Payload.Payload.Url
	ValidateMessage(e, message.Payload.Type)
	if name == "" && text == "" && phone == "" && fileUrl == "" && caption == "" {
		// Esto es porque recibo "basura"
	} else {
		fmt.Println(color.Green("Nombre:"), name)
		fmt.Println(color.Green("Texto:"), text)
		fmt.Println(color.Green("Caption:"), caption)
		fmt.Println(color.Green("Teléfono:"), phone)
		fmt.Println(color.Green("Url:"), fileUrl)
	}
	// Por si alguno mensaje comienza en hola.
	// No es necesario, solo es de prueba.
	if strings.HasPrefix(text, "hola") || strings.HasPrefix(text, "Hola") {
		return e.String(200, `¡Hola! *`+name+`*, ¿En qué puedo ayudarte?`)
	}
	fmt.Println("----------------------")
	e.String(200, "👋🏿")
	return nil

}

func main() {
	e := echo.New()
	e.POST("/", Post)
	e.Logger.Fatal(e.Start(":8000"))
}
