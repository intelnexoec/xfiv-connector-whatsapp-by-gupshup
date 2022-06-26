package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/paij0se/xfiv/others"
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
			Title   string `json:"title"`
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
	events := []string{"delivered", "seen", "enqueued", "sent"}
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
	// Por si algun mensaje comienza en hola.
	// No es necesario, solo es de prueba.
	if strings.HasPrefix(text, "hola") || strings.HasPrefix(text, "Hola") {
		params := url.Values{}
		params.Add("channel", `whatsapp`)
		params.Add("source", "917834811114")
		params.Add("destination", phone)
		params.Add("message", `{"type":"quick_reply","content":{"type":"text","text":"¡Hola *`+name+`*!, ¿Qué servicio quieres contratar?","caption":"Servicios⬇"},"options":[{"type":"text","title":"IA🤖"},{"type":"text","title":"Cripto💱"},{"type":"text","title":"Cloud☁"}]}`)
		params.Add("src.name", "MrTelephoneMen")
		body := strings.NewReader(params.Encode())

		req, err := http.NewRequest("POST", "https://api.gupshup.io/sm/api/v1/msg", body)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Apikey", others.GoDotEnvVariable("API_KEY"))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		res, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(res))
	}
	fmt.Println("----------------------")
	reply := message.Payload.Payload.Title
	switch reply {
	case "IA🤖":
		e.String(200, `*`+name+`*, La IA es el futuro...`)
	case "Cripto💱":
		e.String(200, `*`+name+`*, Es el banco de criptomonedas...`)
	case "Cloud☁":
		e.String(200, `¡Hola *`+name+`*!, Los servicios Cloud están en desarrollo...`)
	}

	e.String(200, "👋🏿")
	return nil

}

func main() {
	e := echo.New()
	e.POST("/", Post)
	e.Logger.Fatal(e.Start(":8000"))
}
