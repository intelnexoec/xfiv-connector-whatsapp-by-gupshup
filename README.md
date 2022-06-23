<h1>A little golang library to interact with the gupshup.io WhatsApp API</h1>

<h1>Getting started</h1>
- For now, you have 6 functions

```go
func SendList(From string, To string, ApiKey string, BotName string, JsonFile string) string

func SendLocation(From string, To string, ApiKey string, BotName string, Place string, Address string, Lon string, Lat string) string

func GetOptInUserList(AppName string, ApiKey string) string

func SendMessage(from string, to string, apiKey string, BotName string, Text string) string

func SendButton(from string, to string, apiKey string, BotName string, Title string, Caption string, m1 string, m2 string, m3 string) string

func GoDotEnvVariable(key string) string
```

<h1>Testing</h1>

-  Create a .env file with your API_KEY. The recipt number and the sender Number

```sh
API_KEY=
NUMBER=
API_NUMBER=
```

- Then run:

```go
$ go test
```

<h1>Demo</h1>



https://user-images.githubusercontent.com/69026987/175425493-7345397e-71b6-4860-9a11-078ca9320dc9.mp4

