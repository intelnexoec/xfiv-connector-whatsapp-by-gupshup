package messages

import (
	"log"
	"testing"

	"github.com/paij0se/xfiv/others"
)

func TestButton(T *testing.T) {
	/*
		Numbers := []string{"573008181818", "573008181818", "573008181818", "573221234567"}
		for _, Number := range Numbers {
			log.Println(SendButton(GoDotEnvVariable("ApiNumber"), Number, GoDotEnvVariable("ApiKey"), "paij0seXfiv", "Selecciona el metodo de pago", "⬇⬇⬇", "Efectivo💲💰", "Crédito/Débito💳", "Btc/Eth🤖"))
		}
	*/
	log.Println(SendButton(others.GoDotEnvVariable("API_NUMBER"), others.GoDotEnvVariable("NUMBER"), others.GoDotEnvVariable("API_KEY"), "paij0setest", "Selecciona el metodo de pago", "⬇⬇⬇", "Efectivo💲💰", "Crédito/Débito💳", "Btc/Eth🤖"))

}

func TestMessage(T *testing.T) {
	log.Println(SendMessage(others.GoDotEnvVariable("API_NUMBER"), others.GoDotEnvVariable("NUMBER"), others.GoDotEnvVariable("API_KEY"), "paij0setest", "Bienvenido a Xfiv"))

}

func TestOptInUserList(T *testing.T) {
	log.Println(GetOptInUserList("paij0setest", others.GoDotEnvVariable("API_KEY")))
}

func TestLocation(T *testing.T) {
	log.Println(SendLocation(others.GoDotEnvVariable("API_NUMBER"), others.GoDotEnvVariable("NUMBER"), others.GoDotEnvVariable("API_KEY"), "paij0setest", "La Loma", "La Loma, El Paso, Cesar", "9.6191", "73.603325")) // Bueno, aquí seria poner las coordeanas bien.
}

func TestList(t *testing.T) {
	log.Println(SendList(others.GoDotEnvVariable("API_NUMBER"), others.GoDotEnvVariable("NUMBER"), others.GoDotEnvVariable("API_KEY"), "paij0setest", "templates/list.json"))
}
