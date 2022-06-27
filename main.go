package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/paij0se/xfiv/messages"
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

type Csml struct {
	RequestID string `json:"request_id"`
	Client    struct {
		BotID     string `json:"bot_id"`
		ChannelID string `json:"channel_id"`
		UserID    string `json:"user_id"`
	} `json:"client"`
	ConversationEnd bool `json:"conversation_end"`
	Messages        []struct {
		ConversationID   string `json:"conversation_id"`
		Direction        string `json:"direction"`
		InteractionOrder int    `json:"interaction_order"`
		Payload          struct {
			Content struct {
				Text    string `json:"text"`
				Url     string `json:"url"`
				Buttons []struct {
					Content struct {
						Accepts     []string `json:"accepts"`
						Payload     string   `json:"payload"`
						TitleButton string   `json:"title"`
					} `json:"content"`
					ContentType string `json:"content_type"`
				} `json:"buttons"`
				Title string `json:"title"`
			} `json:"content"`
			ContentType string `json:"content_type"`
		} `json:"payload"`
	} `json:"messages"`
	ReceivedAt   time.Time `json:"received_at"`
	IsAuthorized bool      `json:"is_authorized"`
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

func PostCsml(name, text, phone string) {
	url := "https://clients.csml.dev/v1/api/chat"
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"client": {` + "" + `"user_id": "1132110816"` + "" + `},` + "" + `"metadata": {` + "" + `"first_name": "` + name + `",` + "" + `"last_name": "Marret"` + "" + `},` + "" + `"request_id": "3be063f9-a803-4dbd-af52-0760257348d4",` + "" + `"payload": {` + "" + `"content": {` + "" + `"text": "` + text + `"` + "" + `},` + "" + `"content_type": "text"` + "" + `}` + "" + `}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("x-api-key", others.GoDotEnvVariable("CSML"))
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		log.Println(error)
	}
	fmt.Println(prettyJSON.String())
	var csml Csml
	json.Unmarshal(body, &csml)
	ContentType := csml.Messages[0].Payload.ContentType
	ButtonsTitle := csml.Messages[0].Payload.Content.Buttons
	fmt.Println("Content Type:", ContentType)
	switch ContentType {
	case "question":
		Title := strings.ReplaceAll(csml.Messages[0].Payload.Content.Title, "Null", name)
		messages.SendButton("917834811114", phone, others.GoDotEnvVariable("API_KEY"), "Testspaij0se", Title, "", ButtonsTitle[0].Content.TitleButton, ButtonsTitle[1].Content.TitleButton, ButtonsTitle[2].Content.TitleButton)
	case "text":
		fmt.Println(csml.Messages[0].Direction, ":", csml.Messages[0].Payload.Content.Text)
		messages.SendMessage("917834811114", phone, others.GoDotEnvVariable("API_KEY"), "Testspaij0se", csml.Messages[0].Payload.Content.Text)
	case "image":
		fmt.Println(csml.Messages[0].Payload.Content.Title, ":", csml.Messages[0].Payload.Content.Url)
	default:
		fmt.Println("No se puede procesar el tipo de contenido")
	}
}

func Post(e echo.Context) error {
	reqBody, err := ioutil.ReadAll(e.Request().Body)
	if err != nil {
		log.Println(err)
	}
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, reqBody, "", "\t")
	if error != nil {
		log.Println(error)
	}
	fmt.Println(prettyJSON.String())
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
		PostCsml(name, text, phone)
		fmt.Println("_________________________________________________")
	}
	return nil

}

func main() {
	e := echo.New()
	e.POST("/", Post)
	e.Logger.Fatal(e.Start(":8000"))
}
